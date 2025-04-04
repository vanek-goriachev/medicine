package tag

type UserActions struct {
	simpleActions SimpleActions
}

func NewUserActions(
	simpleActions SimpleActions,
) *UserActions {
	return &UserActions{
		simpleActions: simpleActions,
	}
}
