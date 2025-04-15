package collections

import "medicine/internal/layers/adapters/storage/file-storage/gorm/file"

type FileStorageMappers struct {
	file *file.GORMMapper
}

func NewFileStorageMappers() *FileStorageMappers {
	var c FileStorageMappers

	c.file = file.NewGORMMapper()

	return &c
}
