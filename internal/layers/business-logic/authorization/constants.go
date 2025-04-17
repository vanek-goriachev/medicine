package authorization

const (
	AttachTagPermission Permission = "medicine.AttachTagPermission"
	CreateTagPermission Permission = "medicine.CreateTagPermission"
	DeleteTagPermission Permission = "medicine.DeleteTagPermission"

	ReadTagsSpacePermission   Permission = "medicine.ReadTagsSpacePermission"
	CreateTagsSpacePermission Permission = "medicine.CreateTagsSpacePermission"
	DeleteTagsSpacePermission Permission = "medicine.DeleteTagsSpacePermission"

	ReadVisitRecordPermission   Permission = "medicine.ReadVisitRecordPermission"
	CreateVisitRecordPermission Permission = "medicine.CreateVisitRecordPermission"
	AttachMedicalFilePermission Permission = "medicine.AttachMedicalFilePermission"
	DeleteVisitRecordPermission Permission = "medicine.DeleteVisitRecordPermission"

	UploadMedicalFilePermission Permission = "medicine.UploadMedicalFilePermission"
)

const (
	TagResource       Resource = "medicine.tag"
	TagsSpaceResource Resource = "medicine.tagsSpace"

	VisitRecordResource Resource = "medicine.visitRecord"

	MedicalFileResource Resource = "medicine.medicalFile"
)
