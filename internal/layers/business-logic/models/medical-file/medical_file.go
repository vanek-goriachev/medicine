package medical_file

import entityID "medicine/pkg/entity-id"

// UploadedMedicalFile is a struct representing MedicalFile after upload, but before saving
type UploadedMedicalFile struct {
	Name string
	Data *[]byte
}

type Extension string

type MedicalFileInfo struct {
	ID        entityID.EntityID
	Extension Extension
	Name      string
}

type DataType *[]byte

type MedicalFileData struct {
	Data DataType
}

type MedicalFile struct {
	MedicalFileInfo
	MedicalFileData
}
