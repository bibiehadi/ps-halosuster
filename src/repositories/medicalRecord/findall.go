package medicalrecordrepository

import (
	"fmt"
	"halosuster/src/entities"
	"strconv"
	"strings"
)

func (r *medicalRecordRepository) GetAllMedicalRecord(params entities.MedicalRecordQueryParams) ([]entities.MedicalRecordResponse, error) {
	var query string = "SELECT o.id AS transactionId, o.customer_id AS customerId, od.product_id AS productId, od.quantity AS quantity, o.paid AS paid, o.change AS change, o.created_at AS createdAt FROM orders o INNER JOIN order_details od ON o.id = od.order_id "
	conditions := ""
	//SELECT
	//  p.identity_number AS identityNumber, -- Assuming "identity_number" exists in patients table
	//  p.phone_number AS phoneNumber,
	//  p.name,
	//  p.birth_date AS birthDate,
	//  p.gender,
	//  p.identity_card_scan_img AS identityCardScanImg,
	//  m.symptoms,
	//  m.medications,
	//  m.created_at,
	//  u.nip,
	//  u.name AS created_by_name,
	//  u.user_id
	//FROM
	//  patients p
	//INNER JOIN medical_records m ON p.id = m.patient_id
	//INNER JOIN users u ON m.created_by = u.id
	//-- Specify your filtering condition here (e.g., WHERE p.id = specific_id)
	//ORDER BY m.created_at DESC;

	// Filter by ID
	if params.CustomerId != "" {
		conditions += " customer_id = '" + params.CustomerId + "' AND"
	}
	if conditions != "" {
		conditions = " WHERE " + strings.TrimSuffix(conditions, " AND")
	}
	query += conditions
	var orderBy []string
	if params.CreatedAt != "" {
		orderBy = append(orderBy, "o.created_at "+params.CreatedAt)
	}
	if len(orderBy) > 0 {
		query += " ORDER BY " + strings.Join(orderBy, ", ")
	} else {
		query += " ORDER BY o.created_at DESC"
	}

	query += " LIMIT " + strconv.Itoa(params.Limit) + " OFFSET " + strconv.Itoa(params.Offset)
	rows, err := r.db.Query(context.Background(), query)

	fmt.Println(query)

	if err != nil {
		fmt.Println(err.Error())
		return []entities.HistoryResponse{}, err
	}
	defer rows.Close()
	var Histories []entities.HistoryResponse
	for rows.Next() {
		var history entities.HistoryResponse
		err := rows.Scan(&history.TransactionId, &history.CustomerId, &history.ProductDetails.ProductId, &history.ProductDetails.Quantity, &history.Paid, &history.Change, &history.CreatedAt)
		if err != nil {
			return []entities.HistoryResponse{}, err
		}
		Histories = append(Histories, history)
	}

	fmt.Println(Histories)
	return Histories, nil

}
