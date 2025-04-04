package tags_space

type Service struct {
	mapper   userActionsMapper
	createUA createTagsSpaceUserAction
}

func NewService(
	mapper userActionsMapper,
	createUA createTagsSpaceUserAction,
) *Service {
	return &Service{
		mapper:   mapper,
		createUA: createUA,
	}
}
