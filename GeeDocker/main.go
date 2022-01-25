package main

import (
	"7_day_golang_implement_from_zero/GeeDocker/image"
	"strings"
)

func main() {
	//测试下载文件是否成功
	image.DownloadImageIfNessary(strings.Join([]string{"alpine", "latest"}, ":"))
	//image.DownloadImageIfNessary(strings.Join([]string{"ubuntu", "latest"}, ":"))
}
