package service

import (
	"io"
	"os"

	"cek/helper"
)

type FileServiceImpl struct{}

func NewFileService() FileService {
	return &FileServiceImpl{}
}

func (service *FileServiceImpl) FindFileDepletion(fileName string) []byte {
	fileOpen, err := os.Open(fileName)
	helper.PanicIfError(err)

	fileBytes, err := io.ReadAll(fileOpen)
	helper.PanicIfError(err)

	return fileBytes
}
