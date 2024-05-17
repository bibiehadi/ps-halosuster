package userrepository

import (
	"context"
	"errors"
	"halosuster/src/entities"

	"github.com/jackc/pgx/v5"
)

func (r *userRepository) FindById(userId string) (entities.User, error) {
	var user entities.User
	var query string = "SELECT id, nip, name, role, identity_card_scan_img, is_active FROM users WHERE id = $1"
	err := r.db.QueryRow(context.Background(), query, userId).Scan(&user.ID, &user.NIP, &user.Name, &user.Role, &user.IdentityCardScanImg, &user.IsActive)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entities.User{}, errors.New("USER NOT FOUND")
		}
	}
	return user, err
}

func (r *userRepository) FindByNIP(nip int) (entities.User, error) {
	var user entities.User
	var query string = "SELECT id, nip, name, role, password, identity_card_scan_img, is_active FROM users WHERE nip = $1"
	err := r.db.QueryRow(context.Background(), query, nip).Scan(&user.ID, &user.NIP, &user.Name, &user.Role, &user.Password, &user.IdentityCardScanImg, &user.IsActive)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entities.User{}, errors.New("USER NOT FOUND")
		}
	}
	return user, err
}

func (r *userRepository) NIPisExist(nip int) bool {
	var exist string
	var query string = "SELECT nip FROM users WHERE nip = $1 LIMIT 1"
	err := r.db.QueryRow(context.Background(), query, nip).Scan(&exist)
	if err != nil {
		if err == pgx.ErrNoRows {
			return false
		}
	}
	return true
}
