package iam

import (
	"context"

	userModels "medicine/pkg/user"
)

type Permission string
type Resource string

type IAM interface {
	// Authorize returns nil if authorization is successful, otherwise returns an error
	Authorize(
		ctx context.Context,
		user userModels.User,
		permission Permission,
		resourceType Resource,
		identifier string,
	) error
}

// TODO remove inteface from here

type Impl struct{}

func NewIAM(
	_ context.Context,
	_ Config,
) (*Impl, error) {
	return &Impl{}, nil
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
