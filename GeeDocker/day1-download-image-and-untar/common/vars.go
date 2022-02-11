package common

//vars for linux only, not support windows and mac
const (
	GockerImagePath     = "/var/lib/gocker/image/"
	GockerTempPath      = "/var/lib/gocker/tmp/"
	GockerContainerPath = "/var/run/gocker/containers/"
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
