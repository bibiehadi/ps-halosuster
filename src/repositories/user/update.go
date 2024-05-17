package userrepository

import (
	"context"
	"halosuster/src/entities"
)

func (r *userRepository) Update(userId string, user entities.User) error {
	var query string = "UPDATE users SET nip = $1, nama = $2, role = $3, identity_card_scan_img = $4, is_active = $5 WHERE id = $6"
	_, err := r.db.Exec(context.Background(), query, user.NIP, user.Name, user.Role, user.IdentityCardScanImg, user.IsActive, userId)
	if err != nil {
		return err
	}
	return nil
}
