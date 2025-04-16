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
	DeleteVisitRecordPermission Permission = "medicine.DeleteVisitRecordPermission"

	UploadFilePermission Permission = "medicine.UploadFilePermission"
)

const (
	TagResource       Resource = "medicine.tag"
	TagsSpaceResource Resource = "medicine.tagsSpace"

	VisitRecordResource Resource = "medicine.visitRecord"

	MedicalFileResource Resource = "medicine.medicalFile"
)
