package common

import (
	"archive/tar"
	"bufio"
	"github.com/pkg/errors"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

//Untar 将tar格式的镜像压缩包，解压到指定的目录下
func Untar(tarball string, dstPath string) error {
	//1.打开文件
	file, err := os.Open(tarball)
	if err != nil {
		return errors.New("os.Open failed")
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Println(err)
		}
	}(file)

	//TODO:need bufio.NewReader can make performance better?
	bufReader := bufio.NewReader(file)

	//2.读取文件中的每一行内容
	reader := tar.NewReader(bufReader)
	for {
		header, err := reader.Next()
		if err == io.EOF {
			log.Println("tar file reach end of file EOF!")
			break
		} else if err != nil {
			return err
		}

		fileInfo := header.FileInfo()
		dstFilePath := dstPath + header.Name //dstFilePath can be file path or directory path

		//pax_global_header is git file
		if header.Name == "pax_global_header" {
			continue
		}

		//判断镜像tar包中内容类型，文件和文件夹
		switch header.Typeflag {
		case tar.TypeDir: //目录
			//以什么权限来创建目录
			err = os.MkdirAll(dstFilePath, fileInfo.Mode())
			if err != nil {
				log.Println("os.MkdirAll error", err)
			}

			log.Println("tar.TypeDir")
		case tar.TypeReg: //常规文件
			untarFile(reader, header, dstFilePath)

			log.Println("tar.TypeReg")
		}
		//output header.Name to see see
		log.Printf("name:%s\n", header.Name)
	}

	//4.关闭打开的文件
	return nil
}
func untarFile(reader io.Reader, hdr *tar.Header, dstFilePath string) {
	var (
		err  error
		file *os.File
	)

	//1.mkdir file parent path
	dstPath, _ := filepath.Split(dstFilePath)
	err = os.MkdirAll(dstPath, os.FileMode(hdr.Mode))
	if err != nil {
		log.Println("os.MkdirAll failed", err)
		return
	}

	//2.create dst file
	file, err = os.Create(dstFilePath)
	if err != nil {
		log.Println("os.Create failed", err)
		return
	}

	//3.copy src file content to dst file
	_, err = io.Copy(file, reader)
	if err != nil {
		log.Println("io.Copy failed", err)
		return
	}

	//4.chmod and change time value of file
	err = file.Chmod(os.FileMode(hdr.Mode))
	if err != nil {
		log.Println("Chmod failed", err)
		return
	}
	err = os.Chtimes(dstFilePath, time.Now(), hdr.ModTime)
	if err != nil {
		log.Println("os.Chtimes failed ", err)
	}
}
