package medical_file

import entityID "medicine/pkg/entity-id"

// UploadedMedicalFile is a struct representing MedicalFile after upload, but before saving.
type UploadedMedicalFile struct {
	Data []byte
	Name string
}

type Extension string

type MedicalFileInfo struct {
	Extension Extension
	Name      string
	ID        entityID.EntityID
}

type DataType []byte

type MedicalFileData struct {
	Data DataType
}

type MedicalFile struct {
	MedicalFileData
	MedicalFileInfo
}
