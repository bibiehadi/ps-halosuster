package v1medicalrecord

import (
	"github.com/labstack/echo/v4"
	"halosuster/src/entities"
	"net/http"
	"reflect"
	"strconv"
)

func (controller *MedicalRecordController) GetAll(c echo.Context) error {
	params := entities.MedicalRecordQueryParams{}

	limitStr := c.QueryParam("limit")
	if limitStr != "" {
		limit, err := strconv.Atoi(limitStr)
		if err == nil && limit > 0 {
			params.Limit = limit
		} else {
			return c.JSON(http.StatusBadRequest, entities.ErrorResponse{
				Status:  false,
				Message: "Invalid limit parameter",
			})
		}
	} else {
		params.Limit = 5
	}

	offsetStr := c.QueryParam("offset")
	if offsetStr != "" {
		offset, err := strconv.Atoi(offsetStr)
		if err == nil && offset >= 0 {
			params.Offset = offset
		} else {
			return c.JSON(http.StatusBadRequest, entities.ErrorResponse{
				Status:  false,
				Message: "Invalid offset parameter",
			})
		}
	} else {
		params.Offset = 0
	}
	if id := c.QueryParam("identityNumber"); id != "" {
		params.IdentityNumber, _ = strconv.Atoi(id)
	}
	if userId := c.QueryParam("userId"); userId != "" {
		params.UserId = userId
	}
	if nip := c.QueryParam("nip"); nip != "" {
		params.NIP = nip

	}
	if createdAt := c.QueryParam("createdAt"); createdAt != "" {
		if createdAt != "asc" && createdAt != "desc" {
			// return c.JSON(http.StatusBadRequest, entities.ErrorResponse{
			// 	Status:  false,
			// 	Message: "Invalid createdAt parameter",
			// })
			// params.CreatedAt = createdAt
		} else {
			params.CreatedAt = createdAt
		}
	}
	// Call service to find products
	MedicalRecords, err := controller.medicalRecordService.GetAll(params)
	if err != nil {
		// fmt.Println("ERROR: %s", err)
		// if err.Error() == "PRODUCTID IS NOT FOUND" {
		// 	return c.JSON(http.StatusNotFound, entities.ErrorResponse{
		// 		Status:  false,
		// 		Message: "Product is not found",
		// 	})
		// }
		return c.JSON(http.StatusInternalServerError, entities.ErrorResponse{
			Status:  false,
			Message: "Failed to fetch Medical Records",
		})
	}

	if MedicalRecords == nil || reflect.ValueOf(MedicalRecords).IsNil() {
		return c.JSON(http.StatusOK, entities.SuccessResponse{
			Message: "success",
			Data:    []entities.MedicalRecordResponse{},
		})
	}

	return c.JSON(http.StatusOK, entities.SuccessResponse{
		Message: "success",
		Data:    MedicalRecords,
	})
}
