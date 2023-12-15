package file

import (
	"fmt"
	"io"
	"os"
)

func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func CreateDir(dirs ...string) (err error) {
	for _, v := range dirs {
		exist, err := Exists(v)
		if err != nil {
			return err
		}
		if !exist {
			if err := os.MkdirAll(v, os.ModePerm); err != nil {
				return err
			}
		}
	}
	return err
}

func SaveFile(path, data string, append bool) error {
	var f *os.File
	if exists, err := Exists(path); err == nil && exists { //如果文件存在
		f, _ = os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0666) //打开文件
	} else {
		f, _ = os.Create(path) //创建文件
	}
	defer f.Close()
	_, err := io.WriteString(f, data) //写入文件(字符串)
	return err
}

func CopyFile(src, dest string) error {
	if src == "" || dest == "" {
		return fmt.Errorf("invalid src %s or destion %s", src, dest)
	}

	srcFileStat, err := os.Stat(src)
	if err != nil {
		return err
	}

	if !srcFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destination.Close()
	// copy file
	_, err = io.Copy(destination, source)
	return err
}

func IsDir(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return fi.Mode().IsDir(), nil
}

func IsFile(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return fi.Mode().IsRegular(), nil
}
