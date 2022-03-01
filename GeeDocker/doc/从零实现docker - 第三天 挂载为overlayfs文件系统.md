**本文是 7天从零实现docker的第三篇 挂载为overlayfs文件系统**

- 了解go-containerregistry的tarball镜像格式
- 处理镜像多层Layer数据，**代码约60行**
- 文件内容拷贝函数封装，**代码约30行**

# 一、OverlayFS介绍

## 1、1 OverlayFS是什么

OverlayFS是一个面向Linux的文件系统服务，顾名思义是一种堆叠文件系统，可以将多个目录的内容 叠加到 另一个目录上。



OverlayFS是一种联合文件系统(UnionFS)，可让你使用2个目录挂载文件系统：

- lowerdir:“下层”目录
- upperdir:“上层”目录

基本上：

◈ 文件系统的**下层**目录是只读的

◈ 文件系统的**上层**目录可以读写

上层目录和下层目录分工明确。



## 1、2 overlayfs简单实践





先约定几个目录

| 目录路径         | 意义                    |
| ---------------- | ----------------------- |
| 镜像临时存储路径 |                         |
| 镜像文件存储路径 | /var/lib/gocker/images/ |
|                  |                         |



overlay文件系统的使用



/var/lib/gocker/images/c059bfaa849c/47d7af55c64c/fs

其中c059bfaa849c和47d7af55c64c这两个值分别代表什么？



下载的镜像是alpine:latest

containerID的计算方法是什么原理？

containerID是不是需要基于内容来实现？否则会出现很大的问题。



这几个目录都是干什么的，需要明白下面的几个目录：

1、镜像存储目录

2、



image hash value is manifest [:12] bits
imageShaHex = manifest.Config.Digest.Hex[:12]



image base path目录

/var/lib/gocker/images/c059bfaa849c



容器fs目录

/var/run/gocker/containers/1de5319fea58/fs



文件拷贝到/var/lib/gocker/images/后如图所示，上节课需要展示出来。



/var/lib/gocker/images/c059bfaa849c/47d7af55c64c/fs 

/var/lib/gocker/images/images.json

create gocker container path
/var/run/gocker/containers/e910dffe0e1a



先讲一下overlayfs的用法和效果。让大家有个感性的认识。

mount挂载点内容

lowerdir=/var/lib/gocker/images/c059bfaa849c/47d7af55c64c/fs,

upperdir=/var/run/gocker/containers/1de5319fea58/fs/upperdir,

workdir=/var/run/gocker/containers/1de5319fea58/fs/workdir

分别设置了lowerdir，upperdir，workdir参数，这三个参数会在之前讲解。





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