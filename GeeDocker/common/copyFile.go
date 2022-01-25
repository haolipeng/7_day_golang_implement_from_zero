package common

import (
	"io"
	"os"
)

func CopyFile(src, dst string) error {
	//1.open src file
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	//2.create dst file
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	//3.copy file content
	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}

	return nil
}
