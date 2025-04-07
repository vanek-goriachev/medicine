package authorization

import (
	"context"

	"medicine/internal/tooling/iam"
	userModels "medicine/pkg/user"
)

type Action struct {
	permission   Permission
	resourceType Resource
	identifier   string
}

func NewAction(permission Permission, resourceType Resource, identifier string) Action {
	return Action{
		permission:   permission,
		resourceType: resourceType,
		identifier:   identifier,
	}
}

type Authorizer interface {
	// Authorize returns nil if authorization is successful, otherwise returns an error
	Authorize(
		ctx context.Context,
		user userModels.User,
		action Action,
	) error
	BatchAuthorize(
		ctx context.Context,
		user userModels.User,
		actions []Action,
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
	_ Action,
) error {
	return nil
}

func (*AuthorizerImpl) BatchAuthorize(
	_ context.Context,
	_ userModels.User,
	_ []Action,
) error {
	return nil
}
