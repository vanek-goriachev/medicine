package visit_record

import (
	"context"
	"fmt"
	"time"

	"medicine/internal/layers/business-logic/authorization"
	visitRecordModels "medicine/internal/layers/business-logic/models/visit-record"
	entityID "medicine/pkg/entity-id"
	userModels "medicine/pkg/user"
)

type VisitRecordCreateIn struct {
	Name     string
	Datetime time.Time

	TagIDs []entityID.EntityID
}

type VisitRecordCreateOut struct {
	VisitRecordLinkedEntities visitRecordModels.VisitRecordLinkedEntities
	VisitRecord               visitRecordModels.VisitRecord
}

type CreateUA struct {
	authorizer    authorization.Authorizer
	simpleActions SimpleActions
}

func NewCreateUA(
	authorizer authorization.Authorizer,
	simpleActions SimpleActions,
) *CreateUA {
	return &CreateUA{
		simpleActions: simpleActions,
		authorizer:    authorizer,
	}
}

func (ua *CreateUA) Act(
	ctx context.Context,
	user userModels.User,
	in *VisitRecordCreateIn,
) (VisitRecordCreateOut, error) {
	err := ua.checkPermissions(ctx, user, in.TagIDs)
	if err != nil {
		return VisitRecordCreateOut{}, err
	}

	visitRecord, visitRecordLinkedEntities, err := ua.simpleActions.CreateWithEntities(
		ctx,
		in.Name,
		in.Datetime,
		in.TagIDs,
	)
	if err != nil {
		return VisitRecordCreateOut{}, fmt.Errorf("can't create visitRecord (ua): %w", err)
	}

	return VisitRecordCreateOut{
		VisitRecord:               visitRecord,
		VisitRecordLinkedEntities: visitRecordLinkedEntities,
	}, nil
}

func (ua *CreateUA) checkPermissions(
	ctx context.Context,
	user userModels.User,
	tagIDs []entityID.EntityID,
) error {
	actions := make([]authorization.Action, 0)

	for _, tagID := range tagIDs {
		actions = append(
			actions, authorization.NewAction(
				authorization.AttachTagPermission,
				authorization.TagResource,
				tagID,
			),
		)
	}

	actions = append(
		actions,
		authorization.NewAction(
			authorization.CreateVisitRecordPermission,
			authorization.VisitRecordResource,
			entityID.EntityID{},
		),
		authorization.NewAction(
			authorization.UploadFilePermission,
			authorization.MedicalFileResource,
			entityID.EntityID{},
		),
	)

	err := ua.authorizer.BatchAuthorize(
		ctx,
		user,
		actions,
	)

	if err != nil {
		return authorization.NewUnauthorizedError(err)
	}

	return nil
}
