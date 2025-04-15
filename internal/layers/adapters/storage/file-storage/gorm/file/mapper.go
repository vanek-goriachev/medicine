package file

import (
	"github.com/google/uuid"
	
	fileModels "medicine/internal/layers/business-logic/models/file"
	gormModels "medicine/internal/layers/storage/file-storage/gorm/models"
	entityID "medicine/pkg/entity-id"
)

type GORMMapper struct{}

func NewGORMMapper() *GORMMapper {
	return &GORMMapper{}
}

func (*GORMMapper) FromGORM(dbFile gormModels.File) fileModels.File {
	return fileModels.File{
		ID:   entityID.EntityID(dbFile.ID),
		Data: fileModels.FileDataType(dbFile.Data),
	}
}

func (t *GORMMapper) MultipleFromGORM(dbFiles []gormModels.File) []fileModels.File {
	files := make([]fileModels.File, len(dbFiles))
	for i, dbFile := range dbFiles {
		files[i] = t.FromGORM(dbFile)
	}

	return files
}

func (*GORMMapper) ToGORM(file fileModels.File) gormModels.File {
	return gormModels.File{
		ID:   uuid.UUID(file.ID),
		Data: gormModels.FileDataType(file.Data),
	}
}

func (t *GORMMapper) MultipleToGORM(files []fileModels.File) []gormModels.File {
	dbFiles := make([]gormModels.File, len(files))
	for i, file := range files {
		dbFiles[i] = t.ToGORM(file)
	}

	return dbFiles
}
