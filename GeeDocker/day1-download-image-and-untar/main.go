package main

import (
	"day1-download-image-and-untar/image"
	"strings"
)

func main() {
	//测试下载文件是否成功
	err := image.DownloadImageIfNessary(strings.Join([]string{"alpine", "latest"}, ":"))
	if err != nil {

	}
}
