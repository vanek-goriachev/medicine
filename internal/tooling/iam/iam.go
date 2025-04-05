package iam

import (
	"context"

	userModels "medicine/pkg/user"
)

type Permission string
type Resource string

type IAM interface {
	// Authenticate returns a user if authentication is successful, otherwise returns an error
	Authenticate(
		ctx context.Context,
	) (userModels.User, error)
	// Authorize returns nil if authorization is successful, otherwise returns an error
	Authorize(
		ctx context.Context,
		user userModels.User,
		permission Permission,
		resourceType Resource,
		identifier string,
	) error
}

type Impl struct{}

func NewIAM(
	_ context.Context,
	_ Config,
) (*Impl, error) {
	return &Impl{}, nil
}

//nolint:unparam // Error is needed in signature
func (*Impl) Authenticate(
	_ context.Context,
) (userModels.User, error) {
	return userModels.User{
		ID:          VanekID,
		IsAnonymous: false,
	}, nil
}

func (*Impl) Authorize(
	_ context.Context,
	_ userModels.User,
	_ Permission,
	_ Resource,
	_ string,
) error {
	return nil
}
