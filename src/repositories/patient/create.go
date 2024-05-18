package patientrepository

import (
	"context"
	"fmt"
	"halosuster/src/entities"
	"strconv"
)

func (r *patientRepository) Create(patient entities.Patient) (entities.Patient, error) {
	var query string = "INSERT INTO patiens (identity_number, phone_number, name, birth_date, gender, identity_card_scan_img) values ($1,$2,$3,$4,$5,$6) RETURNING identity_number"
	var patientId string
	err := r.db.QueryRow(context.Background(), query, patient.IdentityNumber, patient.PhoneNumber, patient.Name, patient.BirthDate, patient.Gender, patient.IdentityCardScanImg).Scan(
		&patientId,
	)
	fmt.Println(err)
	if err != nil {
		return entities.Patient{}, err
	}
	identityNumber, err := strconv.ParseInt(patientId, 10, 64)
	if err != nil {
		fmt.Println("Error converting identity number:", err)
	}
	patient.IdentityNumber = identityNumber
	return patient, err
}
