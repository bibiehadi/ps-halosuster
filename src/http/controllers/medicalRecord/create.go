package v1medicalrecord

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"halosuster/src/entities"
	"halosuster/src/helpers"
	"io"
	"net/http"
	"strconv"
)

func (controller *MedicalRecordController) CreateMedicalRecord(c echo.Context) error {
	var medicalRecordReq entities.MedicalRecordRequest
	fmt.Println("masuk create medical record", medicalRecordReq)
	bindError := c.Bind(&medicalRecordReq)
	fmt.Println("bind error", bindError)

	userId, _ := helpers.GetUserIDFromJWTClaims(c)

	medicalRecordReq.CreatedBy = strconv.Itoa(userId)
	fmt.Println("medical record req", medicalRecordReq)
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

	if err := controller.validator.Struct(medicalRecordReq); err != nil {
		var validationErrors []string
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, fmt.Sprintf("%s is %s", err.Field(), err.Tag()))
		}
		return c.JSON(http.StatusBadRequest, entities.ErrorResponse{
			Status:  false,
			Message: validationErrors,
		})
	}
	fmt.Println("identity number", len(strconv.Itoa(medicalRecordReq.IdentityNumber)))
	if len(strconv.Itoa(medicalRecordReq.IdentityNumber)) < 16 || len(strconv.Itoa(medicalRecordReq.IdentityNumber)) > 16 {
		return c.JSON(http.StatusBadRequest, entities.ErrorResponse{
			Status:  false,
			Message: "identityNumber must be 16 characters",
		})

	}

	patient := controller.patientService.IDisExist(int64(medicalRecordReq.IdentityNumber))
	if !patient {
		return c.JSON(http.StatusNotFound, entities.ErrorResponse{
			Status:  false,
			Message: "identityNumber is not exist",
		})
	}

	medicalRecord, err := controller.medicalRecordService.CreateMedicalRecord(medicalRecordReq)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			entities.ErrorResponse{
				Status:  false,
				Message: err.Error(),
			},
		)

	}
	return c.JSON(
		http.StatusCreated,
		entities.SuccessResponse{
			Message: "success",
			Data:    medicalRecord,
		},
	)
}
