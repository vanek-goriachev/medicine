package authorization

const (
	CreateTagPermission Permission = "medicine.CreateTagPermission"

	ReadTagsSpacePermission   Permission = "medicine.ReadTagsSpacePermission"
	CreateTagsSpacePermission Permission = "medicine.CreateTagsSpacePermission"
)

const (
	TagResource       Resource = "medicine.tag"
	TagsSpaceResource Resource = "medicine.tagsSpace"
)
