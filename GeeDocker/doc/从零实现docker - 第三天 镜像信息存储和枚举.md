# 三、镜像信息持久化

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