package patientrepository

import (
	"context"
	"errors"
	"halosuster/src/entities"

	"github.com/jackc/pgx/v5"
)

func (r *patientRepository) FindById(patientId int64) (entities.Patient, error) {
	var patient entities.Patient
	var query string = "SELECT identity_number FROM patiens WHERE identity_number = $1"
	err := r.db.QueryRow(context.Background(), query, patientId).Scan(&patient.IdentityNumber)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entities.Patient{}, errors.New("PATIENT NOT FOUND")
		}
	}
	return patient, err
}

func (r *patientRepository) IDisExist(patientId int64) bool {
	var exist string
	var query string = "SELECT identity_number FROM patiens WHERE identity_number = $1 LIMIT 1"
	err := r.db.QueryRow(context.Background(), query, patientId).Scan(&exist)
	if err != nil {
		if err == pgx.ErrNoRows {
			return false
		}
	}
	return true
}
