package patientrepository

import (
	"context"
	"fmt"
	"halosuster/src/entities"
	"strconv"
	"strings"
)

func (r *patientRepository) FindAll(params entities.PatientQueryParams) ([]entities.PatientResponse, error) {
	var query string = "SELECT identity_number, phone_number, name, birth_date, gender, created_at FROM patiens "
	conditions := ""

	// Filter by ID
	if string(params.IdentityNumber) != "" {
		conditions += " identity_number = '" + string(params.IdentityNumber) + "' AND"
	}
	if params.Name != "" {
		conditions += " LOWER(name) ILIKE '%" + strings.ToLower(params.Name) + "%' AND"
	}
	if params.PhoneNumber != "" {
		conditions += " phone_number LIKE '%" + string(params.PhoneNumber) + "%' AND"
	}
	if conditions != "" {
		conditions = " WHERE " + strings.TrimSuffix(conditions, " AND")
	}
	query += conditions
	var orderBy []string
	if params.CreatedAt != "" {
		orderBy = append(orderBy, "created_at "+params.CreatedAt)
	}
	if len(orderBy) > 0 {
		query += " ORDER BY " + strings.Join(orderBy, ", ")
	} else {
		query += " ORDER BY created_at DESC"
	}

	query += " LIMIT " + strconv.Itoa(params.Limit) + " OFFSET " + strconv.Itoa(params.Offset)
	rows, err := r.db.Query(context.Background(), query)

	fmt.Println(query)
	fmt.Println(rows)
	fmt.Println(err)

	if err != nil {
		fmt.Println(err.Error())
		return []entities.PatientResponse{}, err
	}
	defer rows.Close()
	var Patients []entities.PatientResponse
	for rows.Next() {
		var patient entities.PatientResponse
		err := rows.Scan(&patient.IdentityNumber, &patient.PhoneNumber, &patient.Name, &patient.BirthDate, &patient.Gender, &patient.CreatedAt)
		fmt.Println(err)
		if err != nil {
			return []entities.PatientResponse{}, err
		}
		Patients = append(Patients, patient)
	}

	fmt.Println(Patients)
	return Patients, nil

}
