package image

import (
	"7_day_golang_implement_from_zero/GeeDocker/common"
	"archive/tar"
	"fmt"
	"github.com/google/go-containerregistry/pkg/crane"
	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/pkg/errors"
	"io"
	"os"
)

//untar 将tar格式的镜像压缩包，解压到指定的目录下
func untar(tarball string, dstPath string) error {
	//1.打开文件
	file, err := os.Open(tarball)
	if err != nil {
		return errors.New("os.Open failed")
	}
	defer file.Close()

	//2.读取文件中的每一行内容
	reader := tar.NewReader(file)
	for {
		header, err := reader.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		fileInfo := header.FileInfo()
		path := dstPath + header.Name

		//镜像tar包中仅有两种类型，文件和文件夹
		switch header.Typeflag {
		case tar.TypeDir:
			//以什么权限来创建目录
			err = os.MkdirAll(path, fileInfo.Mode())
			if err != nil {
				fmt.Println("os.MkdirAll error", err)
			}

			fmt.Println("tar.TypeDir")
		case tar.TypeReg:
			fmt.Println("tar.TypeReg")
		}
		//header.Name的值是什么样的
		fmt.Printf("name:%s\n", header.Name)
	}
	//3.判断内容的类型，进行相应的处理
	//3.1 目录
	//3.2 常规文件
	//4.关闭打开的文件
	return nil
}

func DownloadImageIfNessary(src string) {
	//TODO:判断镜像在本地是否存在，不存在则下载
	downloadImage(src)
}

//downloadImage 下载镜像,src is like "alpine:latest"
func downloadImage(src string) error {
	var (
		image v1.Image
		err   error
	)

	if src == "" {
		return errors.New("download image error,src can't empty!")
	}

	//1.从远程仓库拉取镜像
	image, err = crane.Pull(src)
	if err != nil {
		return errors.Errorf("crane.Pull error: %s", err)
	}

	//2.获取镜像的摘要信息，如sha值
	hash, err := image.Digest()
	if err != nil {
		return errors.Errorf("image.Digest error: %s", err)
	}
	imageHash := hash.Hex[:12]

	//3.构造存储到本地的路径
	imageStorageDir := common.GockerTempPath + imageHash
	err = os.MkdirAll(imageStorageDir, 0755)
	if err != nil {
		return errors.Errorf("os.MkdirAll dir %s error", imageStorageDir)
	}
	imagePath := imageStorageDir + "/package.tar"

	//3.保存镜像到本地路径，先保存到tmp目录
	err = crane.SaveLegacy(image, src, imagePath)
	if err != nil {
		return errors.Errorf("crane.SaveLegacy error: %s", err)
	}

	return err
}
