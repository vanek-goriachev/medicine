package file

import (
	fileModels "medicine/internal/layers/business-logic/models/file"
	gormModels "medicine/internal/layers/storage/file-storage/gorm/models"
)

type fileGORMMapper interface {
	FromGORM(dbFile gormModels.File) fileModels.File
	MultipleFromGORM(dbFiles []gormModels.File) []fileModels.File
	ToGORM(file fileModels.File) gormModels.File
	MultipleToGORM(files []fileModels.File) []gormModels.File
}
