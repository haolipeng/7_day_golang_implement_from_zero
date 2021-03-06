package common

//vars for linux only, not support windows and mac
const (
	GockerImagePath     = "/var/lib/gocker/image/"      //镜像文件存储路径
	GockerTempPath      = "/var/lib/gocker/tmp/"        //下载镜像临时存储目录
	GockerContainerPath = "/var/run/gocker/containers/" //container容器运行目录
)

func GetGockerTempPath() string {
	return GockerTempPath
}

func GetGockerImagePath() string {
	return GockerImagePath
}

func GetGockerContainerPath() string {
	return GockerContainerPath
}
