package service

type FileService interface {
	FindFileDepletion(filename string) []byte
}
