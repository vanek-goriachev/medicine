package authorization

const (
	CreateTagPermission Permission = "medicine.CreateTagPermission"
	DeleteTagPermission Permission = "medicine.DeleteTagPermission"

	ReadTagsSpacePermission   Permission = "medicine.ReadTagsSpacePermission"
	CreateTagsSpacePermission Permission = "medicine.CreateTagsSpacePermission"
	DeleteTagsSpacePermission Permission = "medicine.DeleteTagsSpacePermission"
)

const (
	TagResource       Resource = "medicine.tag"
	TagsSpaceResource Resource = "medicine.tagsSpace"
)
