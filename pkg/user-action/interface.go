package user_action

import (
	"context"

	userModels "medicine/pkg/user"
)

// UserAction is always an atomic action which can be performed by user.
type UserAction[Input any, Output any] interface {
	Act(ctx context.Context, user userModels.User, input Input) (Output, error)
}

// Decorator is some additional logic wraps UserAction.
type Decorator[Input any, InnerInput any, InnerOutput any, Output any] interface {
	UserAction[Input, Output]
	DecoratedAction() UserAction[InnerInput, InnerOutput]
}
