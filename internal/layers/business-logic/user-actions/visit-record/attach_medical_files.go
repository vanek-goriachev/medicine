package visit_record

import (
	"context"
	"fmt"

	"medicine/internal/layers/business-logic/authorization"
	medicalFileModels "medicine/internal/layers/business-logic/models/medical-file"
	entityID "medicine/pkg/entity-id"
	userModels "medicine/pkg/user"
)

type VisitRecordAttachMedicalFilesIn struct {
	UploadedMedicalFiles []medicalFileModels.UploadedMedicalFile
	VisitRecordID        entityID.EntityID
}

type VisitRecordAttachMedicalFilesOut struct{}

type AttachMedicalFilesUA struct {
	authorizer    authorization.Authorizer
	simpleActions SimpleActions
}

func NewAttachMedicalFilesUA(
	authorizer authorization.Authorizer,
	simpleActions SimpleActions,
) *AttachMedicalFilesUA {
	return &AttachMedicalFilesUA{
		simpleActions: simpleActions,
		authorizer:    authorizer,
	}
}

func (ua *AttachMedicalFilesUA) Act(
	ctx context.Context,
	user userModels.User,
	in *VisitRecordAttachMedicalFilesIn,
) (VisitRecordAttachMedicalFilesOut, error) { //nolint:unparam // Signature requires return value
	err := ua.checkPermissions(ctx, user, in.VisitRecordID)
	if err != nil {
		return VisitRecordAttachMedicalFilesOut{}, err
	}

	err = ua.simpleActions.AttachMedicalFiles(
		ctx,
		in.VisitRecordID,
		in.UploadedMedicalFiles,
	)
	if err != nil {
		return VisitRecordAttachMedicalFilesOut{}, fmt.Errorf(
			"can't attach medical files to visitRecord (ua): %w",
			err,
		)
	}

	return VisitRecordAttachMedicalFilesOut{}, nil
}

func (ua *AttachMedicalFilesUA) checkPermissions(
	ctx context.Context,
	user userModels.User,
	visitRecordID entityID.EntityID,
) error {
	actions := []authorization.Action{
		authorization.NewAction(
			authorization.AttachMedicalFilePermission,
			authorization.VisitRecordResource,
			visitRecordID,
		),
		authorization.NewAction(
			authorization.UploadMedicalFilePermission,
			authorization.MedicalFileResource,
			entityID.EntityID{},
		),
	}

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
