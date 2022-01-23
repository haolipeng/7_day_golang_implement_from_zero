package image

import (
	"7_day_golang_implement_from_zero/GeeDocker/common"
	"github.com/google/go-containerregistry/pkg/crane"
	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/pkg/errors"
	"os"
)

func DownloadImageIfNessary(imageName string) {
	//TODO:判断镜像在本地是否存在，不存在则下载

	downloadImage(imageName)
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
