package collections

import (
	file_storage "medicine/internal/appcore/dependencies/file-storage"
	"medicine/internal/layers/storage/file-storage/gorm/file"
)

type FileStorageGateways struct {
	file *file.GORMGateway
}

func NewFileStorageGateways(
	fs *file_storage.FileStorage,
	mappers *FileStorageMappers,
) *FileStorageGateways {
	var c FileStorageGateways

	c.file = file.NewGORMGateway(fs.GormFileStorage, mappers.file)

	return &c
}
