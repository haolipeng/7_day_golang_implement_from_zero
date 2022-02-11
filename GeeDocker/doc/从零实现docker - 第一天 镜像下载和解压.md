**本文是7天从零实现docker的第一篇。**

- 介绍如何用go-containerregistry三方库进行镜像下载，代码量约60行
- 了解镜像tarball包内容及格式
- 解压镜像tarball包的代码



镜像操作的流程图如下：

```json
                           是
镜像 下载 --> 检查是否已缓存 -----> 返回镜像哈希值 ⑴
                |  否             ok           
                |--> 拉取镜像到本地 --> 解压tar格式镜像 --> 处理镜像多层数据 --> 镜像元数据持久化  ⑵
```

本篇文章实现的是流程（2）中的“拉取镜像到本地”和“解压tar格式镜像”这两个步骤。

# 一、镜像下载

虽然docker 官网提供了访问dockerHub仓库的API， https://docs.docker.com/registry/spec/api/，但是用开源库来进行镜像包的下载更方便一些，帮忙处理好很多兼容性问题。

gdocker采用的镜像交互库是go-containerregistry，项目地址：https://github.com/google/go-containerregistry



## 1、1 拉取镜像

这里我最关心的是下载镜像功能，也就是go-containerregistry库中crane pull命令以及其对应的api。

```text
func Pull(src string, opt ...Option) (v1.Image, error)
```

Pull 函数功能：返回 src 标识的远程镜像的 v1.Image。

镜像全称 = 镜像名称 + tag名称 ，如alpine:latest。



## 1、2 计算镜像哈希值

```
// Image defines the interface for interacting with an OCI v1 image.
//go:generate counterfeiter -o fake/image.go . Image
type Image interface {

	// Manifest returns this image's Manifest object.
	Manifest() (*Manifest, error)
}
```

Image接口的Manifest()返回Manifest类型的指针。



```
// Manifest represents the OCI image manifest in a structured way.
type Manifest struct {
	SchemaVersion int64             `json:"schemaVersion,omitempty"`
	MediaType     types.MediaType   `json:"mediaType"`
	Config        Descriptor        `json:"config"`
	Layers        []Descriptor      `json:"layers"`
	Annotations   map[string]string `json:"annotations,omitempty"`
}
```

c059bfaa849c是标识镜像唯一性的hash值，是如何计算出来的呢？

image的manifest的哈希值，取前12位。

标识镜像，必然要有一个唯一的标识，所以我们这样来做。



## 1、3 保存镜像到本地

保存镜像到本地的函数是SaveLegacy

```text
func SaveLegacy(img v1.Image, src, path string) error
```

SaveLegacy 将 v1.Image类型的img写为旧版 tarball压缩包。

path：镜像的存储路径，代码中其默认存储路径为"/var/lib/gocker/tmp/{imageHash}/packaget.tar"



下载的镜像存储到临时目录/var/lib/gocker/tmp/，举例说明：

哈希值为c059bfaa849c的image 镜像的临时存储路径是/var/lib/gocker/tmp/c059bfaa849c/package.tar





## 1、4 完整代码

其完成代码如下所示：

```go
func DownloadImageIfNessary(imageFullName string) error {
   //TODO:判断镜像在本地是否存在，不存在则下载
   //such as "alpine:latest" parse to "alpine" and "latest"
   var (
      image v1.Image
      err   error
   )

   //0.校验参数
   if imageFullName == "" {
      return errors.New("download image error,src can't empty!")
   }

   //1.从远程仓库拉取镜像
   image, err = crane.Pull(imageFullName)
   if err != nil {
      return errors.Errorf("crane.Pull error: %s", err)
   }

   //2.获取镜像的摘要信息，如sha值
   m, err := image.Manifest()
   imageFullHash := m.Config.Digest.Hex
   imageHexHash := imageFullHash[:12]

   err = downloadImage(image, imageFullName, imageHexHash)
   if err != nil {
      log.Println("downloadImage error:", err)
      return err
   }
}

//downloadImage 下载镜像,src is like "alpine:latest"
func downloadImage(image v1.Image, src, imageHash string) error {
	var (
		err error
	)

	//1.构造镜像存储路径，并确保路径存在，默认存储路径为"/var/lib/gocker/tmp/{imageHash}"
	imageStorageDir := common.GockerTempPath + imageHash
	err = os.MkdirAll(imageStorageDir, 0755)
	if err != nil {
		return errors.Errorf("os.MkdirAll dir %s error", imageStorageDir)
	}
	imagePath := imageStorageDir + "/package.tar"

	//2.保存镜像到本地路径,SaveLegacy保存的镜像格式为tarball
	err = crane.SaveLegacy(image, src, imagePath)
	if err != nil {
		return errors.Errorf("crane.SaveLegacy error: %s", err)
	}

	return err
}
```



# 二、解压tar格式镜像

解压tar压缩包到指定目录下

image 镜像存储路径

/var/lib/gocker/tmp/c059bfaa849c/package.tar

将其在当前路径进行解压

即解压package.tar文件中的文件和文件夹到当前目录



解压tar格式的镜像

TODO:编写demo验证功能（提供tar格式镜像前提下，解压）



三、



///////////////////////////////////////////////以下内容留着下次弄////////////////////////////////////////////////

# 三、处理镜像多层Layer数据

- 从指定镜像的manifest.json中解析出镜像的layer分层，manifest.json文件路径为/var/lib/gocker/tmp/c059bfaa849c/manifest.json
- 解压layer层文件，由于一个镜像可能存在多个layer文件，所以存储目录路径不仅要有镜像的hash值，也要有layer的hash值，解压路径定为：/var/lib/gocker/images/{image-hash}/{layer-hash}/fs，{layer-hash}取其layer哈希完整值的前12位
- 将manifest.json和{fullImageHex}.json都拷贝到/var/lib/gocker/images/{image-hash}/下面，供以后使用，{fullImageHex}加上大括号表明是镜像的完整hex值。
- 

# 四、镜像元数据持久化

镜像下载后会存储在临时目录，其默认值为/var/lib/gocker/tmp/{imageHash}

- 下载镜像并存盘至tmp临时目录
- 解压tar格式镜像
- 处理镜像的layer分层
- 解析manifest信息，计算出哈希值
- 
- 删除临时存储目录

### 4、解析manifest信息，计算出哈希值

从image镜像中解析出manifest信息。

manifest.json中的数据都是json格式，定义好json数据后，进行json反序列化即可。



### 5、images镜像信息的维护和更新

images.json文件是我们自己维护？还是docker的镜像中本来就有这部分的信息。

所有images镜像的信息，存储在images.json文件中

/var/lib/gocker/images/images.json

```json
{
	"ubuntu" : {
					"18.04": "[image-hash]",
					"18.10": "[image-hash]",
					"19.04": "[image-hash]",
					"19.10": "[image-hash]"
				},
	"centos" : {
					"6.0": "[image-hash]",
					"6.1": "[image-hash]",
					"6.2": "[image-hash]",
					"7.0": "[image-hash]"
				}
}
```

存储当前系统上的镜像信息的文件格式如上所示。



采用什么数据结构来存储不同镜像的不同标签tag版本呢？

map[string] map[string] string

确定采用双层map的方式来存储镜像的信息，然后将map数据json序列化后，写入images.json文件中持久化保存。



### 6、删除临时存储目录

使用os.RemoveAll函数



2、探究镜像包格式

我保存的文件名称是package.tar，注意此tarball是采用

镜像格式规范

镜像OCI的版本有哪些？

镜像是一个压缩包，其中包括哪些文件，

每个文件中的字段分别代表什么意思



3、镜像包解压









