package entities

import "time"

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
	IdentityNumber int64     `json:"identityNumber"`
	PhoneNumber    string    `json:"phoneNumber"`
	Name           string    `json:"name"`
	BirthDate      string    `json:"birthDate"`
	Gender         string    `json:"gender"`
	CreatedAt      time.Time `json:"createdAt"`
}
