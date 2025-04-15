package authorization

import (
	"context"
	entityID "medicine/pkg/entity-id"

	"medicine/internal/tooling/iam"
	userModels "medicine/pkg/user"
)

type Action struct {
	permission   Permission
	resourceType Resource
	identifier   entityID.EntityID
}

func NewAction(permission Permission, resourceType Resource, identifier entityID.EntityID) Action {
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

	// AvailableResources returns the list of available resources for the given user and permission
	AvailableResources(
		ctx context.Context,
		user userModels.User,
		resourceType Resource,
		permission Permission,
	) ([]entityID.EntityID, error)
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

func (*AuthorizerImpl) AvailableResources(
	_ context.Context,
	_ userModels.User,
	_ Resource,
	_ Permission,
) ([]entityID.EntityID, error) {
	return nil, nil
}
