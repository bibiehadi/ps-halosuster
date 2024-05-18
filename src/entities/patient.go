package entities

type Patient struct {
	IdentityNumber      int64  `json:"identityNumber" validate:"required"`
	PhoneNumber         string `json:"phoneNumber" validate:"required,min=10,max=15,e164"`
	Name                string `json:"name" validate:"required,min=3,max=30"`
	BirthDate           string `json:"birthDate" validate:"required"`
	Gender              string `json:"gender" validate:"required,oneof=male female"`
	IdentityCardScanImg string `json:"identityCardScanImg" validate:"required,url"`
}

type PatientQueryParams struct {
	IdentityNumber string
	Name           string
	PhoneNumber    string
	CreatedAt      string
	Limit          int
	Offset         int
}

type PatientResponse struct {
	IdentityNumber int64  `json:"identityNumber" validate:"required,len=16"`
	PhoneNumber    string `json:"phoneNumber" validate:"required,e164,min=10,max=15"`
	Name           string `json:"name" validate:"required,min=3,max=30"`
	BirthDate      string `json:"birthDate" validate:"required,datetime=2006-01-02"`
	Gender         string `json:"gender" validate:"required,oneof=male female"`
	CreatedAt      string `json:"createdAt" validate:"required,datetime=2006-01-02T15:04:05Z07:00"`
}
