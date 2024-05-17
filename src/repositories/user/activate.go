package userrepository

import (
	"context"
)

func (r *userRepository) Activate(userId string, password string) error {
	var query string = "UPDATE users SET password = $2, is_active = true WHERE id = $1"
	_, err := r.db.Exec(context.Background(), query, userId, password)
	if err != nil {
		return err
	}
	return nil
}
