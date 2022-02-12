package image

import (
	"day1-download-image-and-untar/common"
	"github.com/google/go-containerregistry/pkg/crane"
	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/pkg/errors"
	"log"
	"os"
)

type ManiFest struct {
	Config   string   `json:"config"`
	RepoTags []string `json:"RepoTags"`
	Layers   []string `json:"Layers"`
}

func DownloadImageIfNessary(imageFullName string) error {
	//TODO:判断镜像在本地是否存在，不存在则下载，存在则直接返回镜像的哈希值
	//such as "alpine:latest" parse to "alpine" and "latest"
	var (
		image v1.Image
		err   error
	)

	//0.校验参数
	if imageFullName == "" {
		return errors.New("download image error,src can't empty!")
	}

	//1.从远程仓库拉取镜像
	image, err = crane.Pull(imageFullName)
	if err != nil {
		return errors.Errorf("crane.Pull error: %s", err)
	}

	//2.获取镜像的哈希值(manifest hex值的前12位)
	m, err := image.Manifest()
	imageFullHash := m.Config.Digest.Hex
	imageHexHash := imageFullHash[:12]

	err = downloadImage(image, imageFullName, imageHexHash)
	if err != nil {
		log.Println("downloadImage error:", err)
		return err
	}

	//3.decompress tar archive file
	untarFile(imageHexHash)

	return err
}

//downloadImage 下载镜像,src is like "alpine:latest"
func downloadImage(image v1.Image, src, imageHash string) error {
	var (
		err error
	)

	//1.构造镜像存储路径，并确保路径存在，默认存储路径为"/var/lib/gocker/tmp/{imageHash}"
	imageStorageDir := common.GockerTempPath + imageHash
	err = os.MkdirAll(imageStorageDir, 0755)
	if err != nil {
		return errors.Errorf("os.MkdirAll dir %s error", imageStorageDir)
	}
	imagePath := imageStorageDir + "/package.tar"

	//2.保存镜像到本地路径,SaveLegacy保存的镜像格式为tarball
	err = crane.SaveLegacy(image, src, imagePath)
	if err != nil {
		return errors.Errorf("crane.SaveLegacy error: %s", err)
	}

	return err
}

//untarFile decompress the tar archive file
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
