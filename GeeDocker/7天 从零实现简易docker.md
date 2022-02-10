# 一、docker介绍

## 1、1 docker是什么

docker是一个开源的应用容器引擎。docker可以让开发者打包他们的应用以及依赖包到一个轻量级、可移植的容器中，然后发布到任何流量的linux服务器。



## 1、2 docker的三个重要概念

**镜像(image)**: 镜像是一个静态的概念，一个特殊的文件系统，包含容器运行所需要的一切；

**容器(container)**：容器是动态的概念，容器是从镜像创建的运行实例，它可以被启动、开始、停止、删除，每个容器是相互隔离的。

**镜像仓库(Repository)**：是集中存储镜像文件的仓库，类似Github，只不过镜像仓库上包含的是其他人公开的、打包好的镜像。



## 1、3 docker的优势

docker越来越受欢迎，主要是以下原因：

- **灵活**：即使是最复杂的应用也可以集装箱化。

- **轻量级**：容器利用并共享主机内核。

- **可互换**：您可以即时部署更新和升级。

- **便携式**：您可以在本地构建，部署到云，并在任何地方运行。

- **可扩展**：您可以增加并自动分发容器副本。

- **可堆叠**：您可以垂直和即时堆叠服务。

  

## 1、4 docker背后的核心技术

Cgroup

Namespace

veth

OCI标准

ORI标准

。。。。。。



小伙伴想进一步学习和实践docker的，我推荐一本开源书籍《Docker——从入门到实践》

https://yeasy.gitbook.io/docker_practice/



# 二、关于 GeeDocker

## 2、1 为什么要写GeeDocker

我本职工作是做容器安全的研发，需要深入研究docker的底层原理，看了一些docker源代码剖析的文章，但是感觉理解的不到位，所以萌生了仿写一个docker的项目。通过仿写项目，在实践理解核心知识点，在实践中解决问题，从而触发思考和总结。



docker项目的源代码庞大而复杂，不适合直接上手进行仿写，所以我仿写的是gocker这个开源项目，项目地址如下：

https://github.com/shuveb/containers-the-hard-way

gocker是Go语言从头实现的mini版的Docker，包括Docker的核心功能，很满足我的心意。



## 2、2 实现的功能特性

### 1）镜像管理

- 镜像下载
- 镜像枚举
- 镜像删除

其具体命令如下所示：

gocker images：枚举本地可用镜像

gocker rmi  <image-id>:删除一个本地镜像（指定image-id）



### 2） 将镜像挂载为Overlayfs文件系统



### 3） 容器和宿主机通信，以访问互联网



### 4）在新容器中运行进程

```
gocker run <--cpus=cpus-max> <--mem=mem-max> <--pids=pids-max> <image[:tag]> </path/to/command>
```

参数解释如下：

--cpus=cpus-max 限制cpu数

--mem=mem-max 限制内存量

--pids=pids-max 限制进程的pids数

image[:tag] 标识镜像名称及其tag标签

</path/to/command> 用户要执行的命令



### 5） 在已运行容器中执行进程

```
gocker exec <container-id> </path/to/command>
```

<container-id> 容器的标识id

</path/to/command> 用户要执行的命令



### 6）容器隔离和限制

​	容器隔离采用namespace命名空间技术（File system、PID、IPC、UTS、Mount、Network）

​	容器资源限制（CPU核心数、内存、PID 数量（用于限制进程））

容器隔离和容器限制部分，会附带小demo来演示相应效果，加深大家的印象。



# 三、目录

- **第一天**：镜像的下载和解压
- **第二天**：创建容器目录并将镜像挂载为OverlayFS
- **第三天**：使用OverlayFS文件系统创建容器进程并运行
- **第四天**：容器和宿主机通信，以及访问互联网
- **第五天**：对容器进程进行资源隔离和限制
- **第六天**：镜像枚举命令
- **第七天**：镜像删除命令



