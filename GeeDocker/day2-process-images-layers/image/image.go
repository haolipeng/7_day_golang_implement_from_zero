package image

import (
	"day2-process-images-layers/common"
	"encoding/json"
	"fmt"
	"github.com/google/go-containerregistry/pkg/crane"
	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
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

	//2.获取镜像的摘要信息，如sha值
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

	//4.process tarball layers
	err = ProcessLayers(imageHexHash, imageFullHash)
	if err != nil {
		fmt.Println("ProcessLayers failed", err)
	}

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

//ProcessLayers process multiple layers of container images
func ProcessLayers(imageHexHash, imageFullHash string) error {
	var (
		err          error
		content      []byte
		manifestPath string
		mf           []ManiFest
	)

	//1.locate manifest.json path and parser json content to object
	manifestPath = filepath.Join(common.GetGockerTempPath(), imageHexHash, "manifest.json")
	content, err = ioutil.ReadFile(manifestPath)
	if err != nil || len(content) == 0 {
		return err
	}

	//2.layers info in manifest.json file
	err = json.Unmarshal(content, &mf)
	if err != nil {
		return err
	}

	//3.judge manifest info,such as multiple manifest or layer is empty
	if len(mf) == 0 {
		return errors.New("can't handle empty manifest!")
	}

	//3.untar the layer.tar archive file
	for i, layer := range mf[0].Layers {
		fmt.Printf("Layer %d:%s\n", i, layer)

		layerTarballPath := filepath.Join(common.GetGockerTempPath(), imageHexHash, layer)
		dstPath := filepath.Join(common.GetGockerImagePath(), imageHexHash, layer[:12], "fs")
		err = os.MkdirAll(dstPath, 0644)
		if err != nil {
			fmt.Printf("os.MkdirAll %s failed!", dstPath)
		}

		//real untar
		err = common.Untar(layerTarballPath, dstPath)
		if err != nil {
			break
		}
	}

	//4.copy manifest.json和{fullImageHex}.json file to /var/lib/gocker/images/{image-hash}/ directory
	//copyFile
	manifestDstPath := filepath.Join(common.GetGockerImagePath(), imageHexHash, "manifest.json")

	configName := fmt.Sprintf("%s.json", imageFullHash)
	configPath := filepath.Join(common.GetGockerTempPath(), imageHexHash, configName)
	configDstPath := filepath.Join(common.GetGockerImagePath(), imageHexHash, configName)

	common.CopyFile(manifestPath, manifestDstPath)
	common.CopyFile(configPath, configDstPath)
	return err
}
