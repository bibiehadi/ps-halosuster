package userrepository

import (
	"context"
	"fmt"
	"halosuster/src/entities"
)

func (r *userRepository) Update(userId string, user entities.User) error {

	fmt.Println("error id")
	var query string = "UPDATE users SET nip = $1, name = $2, role = $3, identity_card_scan_img = $4, is_active = $5, updated_at = $6 WHERE id = $7"
	_, err := r.db.Exec(context.Background(), query, user.NIP, user.Name, user.Role, user.IdentityCardScanImg, user.IsActive, user.UpdatedAt, userId)

	if err != nil {
		return err
	}
	return nil
}
