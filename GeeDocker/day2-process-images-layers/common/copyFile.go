package common

import (
	"io"
	"log"
	"os"
)

func CopyFile(src, dst string) error {
	//1.open src file
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer func(in *os.File) {
		err := in.Close()
		if err != nil {
			log.Println("in file close error")
		}
	}(in)

	//2.create dst file
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer func(out *os.File) {
		err := out.Close()
		if err != nil {
			log.Println("out file close error")
		}
	}(out)

	//3.copy file content
	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}

	return nil
}
