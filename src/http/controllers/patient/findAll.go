package patientcontroller

import (
	"halosuster/src/entities"
	"net/http"
	"reflect"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (controller *patientController) FindAll(c echo.Context) error {
	params := entities.PatientQueryParams{}

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

	if identityNumber := c.QueryParam("identityNumber"); identityNumber != "" {
		params.IdentityNumber = identityNumber
	}

	if name := c.QueryParam("name"); name != "" {
		params.Name = name
	}

	if phoneNumber := c.QueryParam("phoneNumber"); phoneNumber != "" {
		params.PhoneNumber = phoneNumber
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

	// Call service to find patients
	patients, err := controller.patientService.FindAll(params)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, entities.ErrorResponse{
			Status:  false,
			Message: "Failed to fetch patients",
		})
	}

	if patients == nil || reflect.ValueOf(patients).IsNil() {
		return c.JSON(http.StatusOK, entities.SuccessResponse{
			Message: "success",
			Data:    []entities.Patient{},
		})
	}

	return c.JSON(http.StatusOK, entities.SuccessResponse{
		Message: "success",
		Data:    patients,
	})
}
