package userrepository

import (
	"context"
	"errors"
	"halosuster/src/entities"

	"github.com/jackc/pgx/v5"
)

func (r *userRepository) FindById(userId string) (entities.User, error) {
	var user entities.User
	var query string = "SELECT id, nip, name, role, identity_card_scan_img, is_active FROM user WHERE id = $1"
	err := r.db.QueryRow(context.Background(), query, userId).Scan(&user.ID, &user.NIP, &user.Name, &user.Role, &user.IdentityCardScanImg, &user.IsActive)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entities.User{}, errors.New("USER NOT FOUND")
		}
	}
	return user, err
}

func (r *userRepository) NIPisExist(nip int) bool {
	var user entities.User
	var query string = "SELECT id, nip, name, role, identity_card_scan_img, is_active FROM user WHERE nip = $1"
	err := r.db.QueryRow(context.Background(), query, nip).Scan(&user.ID, &user.NIP, &user.Name, &user.Role, &user.IdentityCardScanImg, &user.IsActive)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false
		}
	}
	return true
}
