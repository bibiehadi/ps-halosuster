package medicalrecordrepository

import (
	"context"
	"fmt"
	"halosuster/src/entities"
	"strconv"
	"strings"
)

func (r *medicalRecordRepository) GetAllMedicalRecord(params entities.MedicalRecordQueryParams) ([]entities.MedicalRecordResponse, error) {
	var query string = "SELECT patients.id , patients.phone_number, patients.name, patients.birth_date, patients.gender, patients.identity_card_scan_img, medical_records.sympthoms, medical_records.medications, medical_records.created_at, users.nip, users.name , users.id FROM patients INNER JOIN medical_records ON patients.id = medical_records.patient_id INNER JOIN users ON medical_records.created_by = users.id "
	conditions := ""

	// Filter by ID
	if string(params.IdentityNumber) != "" {
		conditions += " patiens.id = '" + string(params.IdentityNumber) + "' AND"
	}
	if params.UserId != "" {
		conditions += " users.id = '" + params.UserId + "' AND"
	}
	if params.NIP != "" {
		conditions += " users.nip = '" + params.NIP + "' AND"
	}
	if conditions != "" {
		conditions = " WHERE " + strings.TrimSuffix(conditions, " AND")
	}
	query += conditions
	var orderBy []string
	if params.CreatedAt != "" {
		orderBy = append(orderBy, "medical_records.created_at "+params.CreatedAt)
	}
	if len(orderBy) > 0 {
		query += " ORDER BY " + strings.Join(orderBy, ", ")
	} else {
		query += " ORDER BY medical_records.created_at DESC"
	}

	query += " LIMIT " + strconv.Itoa(params.Limit) + " OFFSET " + strconv.Itoa(params.Offset)
	rows, err := r.db.Query(context.Background(), query)

	fmt.Println(query)

	if err != nil {
		fmt.Println(err.Error())
		return []entities.MedicalRecordResponse{}, err
	}
	defer rows.Close()
	var MedicalRecords []entities.MedicalRecordResponse
	for rows.Next() {
		var medicalRecord entities.MedicalRecordResponse
		err := rows.Scan(&medicalRecord.IdentityDetail.IdentityNumber, &medicalRecord.IdentityDetail.PhoneNumber, &medicalRecord.IdentityDetail.Name, &medicalRecord.IdentityDetail.BirthDate, &medicalRecord.IdentityDetail.Gender, &medicalRecord.IdentityDetail.IdentityCardScanImg, &medicalRecord.Sympthoms, &medicalRecord.Medications, &medicalRecord.CreatedAt, &medicalRecord.CreatedBy.Nip, &medicalRecord.CreatedBy.Nip, &medicalRecord.CreatedBy.UserId)
		if err != nil {
			return []entities.MedicalRecordResponse{}, err
		}
		MedicalRecords = append(MedicalRecords, medicalRecord)
	}

	fmt.Println(MedicalRecords)
	return MedicalRecords, nil

}
