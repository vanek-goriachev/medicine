package tags_space

type UserActions struct {
	Create *CreateUA
}

func NewUserActions(
	simpleActions SimpleActions,
) *UserActions {
	create := NewCreateUA(simpleActions)

	return &UserActions{
		Create: create,
	}
}
