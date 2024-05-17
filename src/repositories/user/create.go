package userrepository

import (
	"context"
	"halosuster/src/entities"
)

func (r *userRepository) Create(user entities.User) (entities.User, error) {
	var query string = "INSERT INTO users (nip, name,role, identity_card_scan_img, is_active) values ($1,$2,$3,$4,$5) RETURNING id"
	var userId string
	err := r.db.QueryRow(context.Background(), query, user.NIP, user.Name, user.Role, user.IdentityCardScanImg, user.IsActive).Scan(
		&userId,
	)
	if err != nil {
		return entities.User{}, err
	}
	user.ID = userId
	return user, err
}
