package nursecontroller

import (
	"encoding/json"
	"fmt"
	"halosuster/src/entities"
	"halosuster/src/helpers"
	"io"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func (controller *nurseController) Register(c echo.Context) error {
	var nurseRequest entities.NurseRequest
	bindError := c.Bind(&nurseRequest)

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

	if err := controller.validator.Struct(nurseRequest); err != nil {
		var validationErrors []string
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, fmt.Sprintf("%s is %s", err.Field(), err.Tag()))
		}
		return c.JSON(http.StatusBadRequest, entities.ErrorResponse{
			Status:  false,
			Message: validationErrors,
		})
	}

	if !helpers.ValidateNIP(nurseRequest.NIP) {
		return c.JSON(http.StatusBadRequest, entities.ErrorResponse{
			Status:  false,
			Message: "INVALID FORMAT NURSE NIP",
		})
	}

	if strconv.Itoa(nurseRequest.NIP)[0:3] != "303" {
		return c.JSON(http.StatusNotFound, entities.ErrorResponse{
			Status:  false,
			Message: "INVALID NIP FORMAT",
		})
	}

	if !helpers.ValidateUrl(nurseRequest.IdentityCardScanImg) {
		c.JSON(http.StatusBadRequest, entities.ErrorResponse{
			Status:  false,
			Message: "URL FORMAT IS NOT VALID",
		})
		return nil
	}

	nurse := entities.User{
		NIP:                 nurseRequest.NIP,
		Name:                nurseRequest.Name,
		IdentityCardScanImg: nurseRequest.IdentityCardScanImg,
		Role:                entities.Nurse,
		IsActive:            false,
	}

	user, err := controller.userService.Register(nurse, true)
	if err != nil {
		if err.Error() == "NIP ALREADY EXIST" {
			return c.JSON(http.StatusConflict, entities.ErrorResponse{
				Status:  false,
				Message: err.Error(),
			})
		}

		return c.JSON(http.StatusBadRequest, entities.ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, entities.SuccessResponse{
		Message: "Nurse registered successfull",
		Data: entities.NurseResponse{
			ID:   user.ID,
			NIP:  user.NIP,
			Name: user.Name,
		},
	})
}
