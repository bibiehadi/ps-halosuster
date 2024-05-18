package entities

type Patient struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	PhoneNumber       string `json:"phone_number"`
	BirthDate         string `json:"birth_date"`
	Gender            string `json:"gender"`
	IdentityCardImage string `json:"identity_card_image"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
}
