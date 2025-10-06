package dto

type FilterTermsRequest struct {
	// From query
	Ref string `query:"ref" validate:"omitempty,uuid4"`

	// From path
	OrgID string `param:"org_id" validate:"required,uuid4"`

	// From JSON body
	SchoolID string `json:"school_id" validate:"required,uuid4"`
	UserID   string `json:"user_id" validate:"required,uuid4"`
	Name     string `json:"name" validate:"required,min=2"`
	Email    string `json:"email" validate:"required,email"`
}
