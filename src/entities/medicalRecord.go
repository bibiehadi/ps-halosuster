package entities

type MedicalRecord struct {
	ID          int     `json:"id"`
	PatientID   Patient `json:"patient_id"`
	Sympthoms   string  `json:"sympthoms"`
	Medications string  `json:"medications"`
	CreatedBy   User    `json:"created_by"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

type MedicalRecordRequest struct {
	IdentityNumber int    `json:"identityNumber" validate:"required min=16 max=16"`
	Symptoms       string `json:"symptoms" validate:"required min=1 max=2000"`
	Medications    string `json:"medications" validate:"required min=1 max=2000"`
	CreatedBy      string `json:"createdBy" validate:"required"`
}

type MedicalRecordQueryParams struct {
	IdentityNumber int    `json:"identityNumber"`
	UserId         string `json:"userId"`
	NIP            string `json:"nip"`
	Limit          int    `json:"limit"`
	Offset         int    `json:"offset"`
	CreatedAt      string `json:"createdAt"`
}

type MedicalRecordResponse struct {
	IdentityDetail IdentityDetail `json:"identityDetail"`
	Sympthoms      string         `json:"sympthoms"`
	Medications    string         `json:"medications"`
	CreatedAt      string         `json:"createdAt"`
	CreatedBy      CreatedBy      `json:"createdBy"`
}

type IdentityDetail struct {
	IdentityNumber      int    `json:"identityNumber"`
	PhoneNumber         string `json:"phoneNumber"`
	Name                string `json:"name"`
	BirthDate           string `json:"birthDate"`
	Gender              string `json:"gender"`
	IdentityCardScanImg string `json:"identityCardScanImg"`
}
type CreatedBy struct {
	Nip    int64  `json:"nip"`
	Name   string `json:"name"`
	UserId string `json:"userId"`
}
