package userrepository

import (
	"context"
)

func (r *userRepository) Delete(userId string) error {
	var query string = "DELETE FROM users WHERE id = $1"
	_, err := r.db.Exec(context.Background(), query, userId)
	if err != nil {
		return err
	}
	return nil
}
