在实现代码的过程中，我们是在解决一个个的问题。

将镜像中的分层内容解析到一个文件夹下。

docker镜像的基础知识

manifest： 基本上对应了一个镜像，里面包含了一个镜像的所有layers digest，客户端拉取镜像的时候一般都是先获取manifest 文件，在根据 manifest 文件里面的内容拉取镜像各个层(tar+gzip)



layer：镜像层，镜像层不包含任何的运行时信息，只包含文件系统的信息。镜像是通过最底层的rootfs加上各层的changeset(对上一层的add、update、delete操作)组合而成的。



# 一、管理镜像Image

## 3、1 下载镜像

### 1、下载镜像并存盘至目录

image 镜像临时存储路径:/var/lib/gocker/tmp/c059bfaa849c/package.tar



TODO:编写demo验证功能

crane.Pull和crane.SaveLegacy函数



### 2、解压tar格式镜像

解压tar格式的镜像

TODO:编写demo验证功能（提供tar格式镜像前提下，解压）



### 3、解析manifest信息，计算出哈希值

这个好弄。



### 4、删除临时存储目录

使用os.RemoveAll函数



## 3、2 枚举镜像信息



## 3、3 删除本地镜像



镜像的操作库采用go-containerregistry，是google开源的一个项目，我们这里只用到crane相关的api接口即可。



/var/lib/gocker/images/c059bfaa849c/47d7af55c64c/fs

其中c059bfaa849c和47d7af55c64c这两个值分别代表什么？



layer.tar文件是个简单的文件系统的集合，我们需要把解压后的文件夹放到什么地方呢？



image hash value is manifest [:12] bits
imageShaHex = manifest.Config.Digest.Hex[:12]

different layers file dst path
/var/lib/gocker/images/c059bfaa849c/47d7af55c64c/fs 

/var/lib/gocker/images/images.json

create gocker container path
/var/run/gocker/containers/e910dffe0e1a





# 二、容器隔离 Namespace

- File system (via `chroot`)
- PID
- IPC
- UTS (hostname)
- Mount
- Network



# 三、Cgroup使用

Cgroup限制cpu，内存，pids（带领大家实验下）



命名空间的使用

lsns命令的输出结果

![image-20220118155945325](picture/image-20220118155945325.png)







资料收集

# Building container images in Go

https://ahmet.im/blog/building-container-images-in-go/



使用到的开源库

https://github.com/google/go-containerregistry



https://www.51cto.com/article/697935.html

