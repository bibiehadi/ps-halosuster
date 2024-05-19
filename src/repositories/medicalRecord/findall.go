package medicalrecordrepository

import (
	"context"
	"fmt"
	"halosuster/src/entities"
	"strconv"
	"strings"
)

func (r *medicalRecordRepository) GetAllMedicalRecord(params entities.MedicalRecordQueryParams) ([]entities.MedicalRecordResponse, error) {
	var query string = "SELECT patiens.identity_number , patiens.phone_number, patiens.name, patiens.birth_date, patiens.gender, patiens.identity_card_scan_img, medical_records.symptoms, medical_records.medications, medical_records.created_at, users.nip, users.name , users.id FROM patiens INNER JOIN medical_records ON patiens.identity_number = medical_records.patien_id INNER JOIN users ON medical_records.created_by = users.id "
	conditions := ""

	// Filter by ID
	fmt.Println("params", strconv.Itoa(params.IdentityNumber))
	if params.IdentityNumber != 0 {
		conditions += " patiens.identity_number = '" + string(params.IdentityNumber) + "' AND"
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
		err := rows.Scan(&medicalRecord.IdentityDetail.IdentityNumber, &medicalRecord.IdentityDetail.PhoneNumber, &medicalRecord.IdentityDetail.Name, &medicalRecord.IdentityDetail.BirthDate, &medicalRecord.IdentityDetail.Gender, &medicalRecord.IdentityDetail.IdentityCardScanImg, &medicalRecord.Symptoms, &medicalRecord.Medications, &medicalRecord.CreatedAt, &medicalRecord.CreatedBy.Nip, &medicalRecord.CreatedBy.Name, &medicalRecord.CreatedBy.UserId)
		fmt.Println("error", err)
		fmt.Println("medical record", medicalRecord)
		if err != nil {
			return []entities.MedicalRecordResponse{}, err
		}
		MedicalRecords = append(MedicalRecords, medicalRecord)
	}

	fmt.Println(MedicalRecords)
	return MedicalRecords, nil

}
