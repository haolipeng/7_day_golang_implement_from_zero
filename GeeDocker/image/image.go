package image

import (
	"7_day_golang_implement_from_zero/GeeDocker/common"
	"github.com/google/go-containerregistry/pkg/crane"
	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/pkg/errors"
	"log"
	"os"
)

func DownloadImageIfNessary(imageFullName string) error {
	//TODO:判断镜像在本地是否存在，不存在则下载
	//such as "alpine:latest" parse to "alpine" and "latest"
	var (
		image v1.Image
		err   error
	)

	if imageFullName == "" {
		return errors.New("download image error,src can't empty!")
	}

	//1.从远程仓库拉取镜像
	image, err = crane.Pull(imageFullName)
	if err != nil {
		return errors.Errorf("crane.Pull error: %s", err)
	}

	//2.获取镜像的摘要信息，如sha值
	hash, err := image.Digest()
	if err != nil {
		return errors.Errorf("image.Digest error: %s", err)
	}
	imageHexHash := hash.Hex[:12]

	err = downloadImage(imageFullName, imageHexHash)
	if err != nil {
		log.Println("downloadImage error:", err)
		return err
	}

	//3.decompress tar archive file
	untarFile(imageHexHash)

	return err
}

//downloadImage 下载镜像,src is like "alpine:latest"
func downloadImage(src, imageHash string) error {
	var (
		image v1.Image
		err   error
	)

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

func untarFile(imageHexHash string) {
	var (
		err error
	)

	dstPath := common.GetGockerTempPath() + imageHexHash
	tarballPath := dstPath + "/package.tar"
	err = common.Untar(tarballPath, dstPath)
	if err != nil {
		log.Println(err)
		return
	}
}
