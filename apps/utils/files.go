package utils

import (
	"errors"
	"os"
)

func FilesAddContent(file *os.File, content string) {
	file.WriteString(content)
}

func FilesCreate(path, fileName string) (*os.File, error) {
	create := path + "/" + fileName
	file, err := os.OpenFile(create, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
	if err != nil {
		if os.IsExist(err) {
			return file, errors.New("file already exists")
		} else {
			return file, errors.New(err.Error())
		}
	}
	return file, nil
}
