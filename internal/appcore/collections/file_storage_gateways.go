package collections

import (
	file_storage "medicine/internal/appcore/dependencies/file-storage"
)

type FileStorageGateways struct {
}

func NewFileStorageGateways(fs *file_storage.FileStorage) *FileStorageGateways {
	var c FileStorageGateways

	return &c
}
