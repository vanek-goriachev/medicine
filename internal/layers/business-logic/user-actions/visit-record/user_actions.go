package visit_record

import "medicine/internal/layers/business-logic/authorization"

type UserActions struct {
	Create            *CreateUA
	AttachMedicalFile *AttachMedicalFilesUA
}

func NewUserActions(
	authorizer authorization.Authorizer,
	simpleActions SimpleActions,
) *UserActions {
	createUA := NewCreateUA(authorizer, simpleActions)
	attachMedicalFilesUA := NewAttachMedicalFilesUA(authorizer, simpleActions)

	return &UserActions{
		Create:            createUA,
		AttachMedicalFile: attachMedicalFilesUA,
	}
}
