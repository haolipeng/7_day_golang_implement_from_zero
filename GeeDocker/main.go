package main

import "7_day_golang_implement_from_zero/GeeDocker/image"

func main() {
	//测试下载文件是否成功
	image.DownloadImageIfNessary("alpine:latest")
	tar := archiver.Tar{}
}
