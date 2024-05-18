package patientcontroller

import (
	"encoding/json"
	"fmt"
	"halosuster/src/entities"
	"halosuster/src/helpers"
	"io"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func (controller *patientController) Create(c echo.Context) error {
	var patientRequest entities.Patient
	bindError := c.Bind(&patientRequest)

	if bindError != nil {
		switch bindError.(type) {
		case validator.ValidationErrors:
			var errorMessages string
			for _, e := range bindError.(validator.ValidationErrors) {
				errorMessage := fmt.Sprintf("Error on field: %s, condition: %s", e.Field(), e.ActualTag())
				errorMessages = fmt.Sprintf(errorMessages + errorMessage)
			}
			return c.JSON(
				http.StatusBadRequest,
				entities.ErrorResponse{
					Status:  false,
					Message: errorMessages,
				},
			)

		case *json.UnmarshalTypeError:
			return c.JSON(http.StatusBadRequest, entities.ErrorResponse{
				Status:  false,
				Message: bindError.Error(),
			})

		default:
			if bindError == io.EOF {
				return c.JSON(http.StatusBadRequest, entities.ErrorResponse{
					Status:  false,
					Message: "Request body is empty",
				})

			}
			return c.JSON(http.StatusBadRequest, entities.ErrorResponse{
				Status:  false,
				Message: bindError.Error(),
			})

		}
	}

	if err := controller.validator.Struct(patientRequest); err != nil {
		var validationErrors []string
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, fmt.Sprintf("%s is %s", err.Field(), err.Tag()))
		}
		return c.JSON(http.StatusBadRequest, entities.ErrorResponse{
			Status:  false,
			Message: validationErrors,
		})
	}

	if !helpers.ValidateUrl(patientRequest.IdentityCardScanImg) {
		c.JSON(http.StatusBadRequest, entities.ErrorResponse{
			Status:  false,
			Message: "URL FORMAT IS NOT VALID",
		})
		return nil
	}

	patient, err := controller.patientService.Create(patientRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entities.ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, entities.SuccessResponse{
		Message: "Patient registered successfull",
		Data: entities.PatientResponse{
			IdentityNumber: patient.IdentityNumber,
			Name:           patient.Name,
		},
	})
}
