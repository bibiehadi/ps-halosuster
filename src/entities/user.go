package entities

import "time"

type User struct {
	ID                  string    `json:"id"`
	NIP                 int       `json:"nip" validate:"required"`
	Name                string    `json:"name" validate:"required,min=5,max=50"`
	Role                Role      `json:"role" `
	Password            string    `json:"password"`
	IdentityCardScanImg string    `json:"identityCardScanImg" validate:"required"`
	IsActive            bool      `json:"isActive"`
	CreatedAt           time.Time `json:"createdAt"`
	UpdatedAt           time.Time `json:"updatedAt"`
}

type NurseRequest struct {
	NIP                 int    `json:"nip" validate:"required"`
	Name                string `json:"name" validate:"required,min=5,max=50"`
	IdentityCardScanImg string `json:"identityCardScanImg" validate:"required"`
}

type NurseUpdateRequest struct {
	NIP  int    `json:"nip" validate:"required"`
	Name string `json:"name" validate:"required,min=5,max=50"`
}

type NurseActivate struct {
	Password string `json:"password" validate:"required,min=5, max=33"`
}

type NurseResponse struct {
	ID   string `json:"userId" validate:"required"`
	NIP  int    `json:"nip" validate:"required"`
	Name string `json:"name" validate:"required,min=5,max=50"`
}

type Role string

const (
	IT    Role = "it"
	Nurse Role = "nurse"
)
