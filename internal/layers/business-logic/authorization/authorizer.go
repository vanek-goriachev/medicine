package authorization

import (
	"context"

	"medicine/internal/tooling/iam"
	userModels "medicine/pkg/user"
)

type Authorizer interface {
	// Authorize returns nil if authorization is successful, otherwise returns an error
	Authorize(
		ctx context.Context,
		user userModels.User,
		permission Permission,
		resourceType Resource,
		identifier string,
	) error
}

type AuthorizerImpl struct {
	IAMClient iam.IAM
}

func NewAuthorizer(
	iamClient iam.IAM,
) *AuthorizerImpl {
	return &AuthorizerImpl{
		IAMClient: iamClient,
	}
}

func (*AuthorizerImpl) Authorize(
	_ context.Context,
	_ userModels.User,
	_ Permission,
	_ Resource,
	_ string,
) error {
	return nil
}
