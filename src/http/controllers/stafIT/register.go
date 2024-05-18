package itcontroller

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

func (controller *stafItController) Register(c echo.Context) error {
	var stafITRequest entities.ITRequest
	bindError := c.Bind(&stafITRequest)

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

	if err := controller.validator.Struct(stafITRequest); err != nil {
		var validationErrors []string
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, fmt.Sprintf("%s is %s", err.Field(), err.Tag()))
		}
		return c.JSON(http.StatusBadRequest, entities.ErrorResponse{
			Status:  false,
			Message: validationErrors,
		})
	}

	if !helpers.ValidateNIP(stafITRequest.NIP) {
		return c.JSON(http.StatusBadRequest, entities.ErrorResponse{
			Status:  false,
			Message: "INVALID NIP FORMAT",
		})
	}

	if strconv.Itoa(stafITRequest.NIP)[0:3] != "615" {
		return c.JSON(http.StatusNotFound, entities.ErrorResponse{
			Status:  false,
			Message: "INVALID NIP FORMAT",
		})
	}

	stafIT := entities.User{
		NIP:      stafITRequest.NIP,
		Name:     stafITRequest.Name,
		Password: stafITRequest.Password,
		Role:     entities.IT,
		IsActive: true,
	}
	user, err := controller.userService.Register(stafIT, false)
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

	authRequest := entities.AuthRequest{
		NIP:      stafIT.NIP,
		Password: stafIT.Password,
	}

	token, _, err := controller.userService.Login(authRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entities.ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, entities.SuccessResponse{
		Message: "IT user registered successfull",
		Data: entities.AuthResponse{
			ID:          user.ID,
			NIP:         user.NIP,
			Name:        user.Name,
			AccessToken: token,
		},
	})
}
