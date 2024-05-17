package userrepository

import (
	"context"
	"fmt"
	"halosuster/src/entities"
	"strconv"
	"strings"
)

func (r *userRepository) GetAll(params entities.UserQueryParams) ([]entities.UserResponse, error) {
	var query string = "SELECT id, nip, name, created_at FROM users"

	conditions := ""

	if params.Id != "" {
		conditions += " id = '" + params.Id + "' AND"
	}

	// Filter by Name (wildcard search, case insensitive)
	if params.Name != "" {
		conditions += " LOWER(name) LIKE '%" + strings.ToLower(params.Name) + "%' AND"
	}

	if params.NIP != "" {
		conditions += " nip LIKE '" + params.NIP + "%' AND"
	}

	if params.Role != "" {
		conditions += " role = '" + params.Role + "' AND"
	}

	// Remove trailing "AND"
	if conditions != "" {
		conditions = " WHERE " + strings.TrimSuffix(conditions, " AND")
	}

	// Apply conditions
	query += conditions

	if strings.ToLower(params.CreatedAt) == "asc" {
		query += " ORDER BY created_at ASC"
	} else {
		query += " ORDER BY created_at DESC"
	}

	// Apply limit and offset
	query += " LIMIT " + strconv.Itoa(params.Limit) + " OFFSET " + strconv.Itoa(params.Offset)
	fmt.Printf("QUERY : ", query)
	fmt.Println()
	rows, err := r.db.Query(context.Background(), query)

	if err != nil {
		return nil, err
	}

	var users []entities.UserResponse

	for rows.Next() {
		var user entities.UserResponse
		err := rows.Scan(&user.ID, &user.NIP, &user.Name, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, err
}
