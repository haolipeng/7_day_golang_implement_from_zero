package common

import (
	"archive/tar"
	"fmt"
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
