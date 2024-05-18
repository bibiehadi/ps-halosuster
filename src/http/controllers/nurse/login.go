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

func (controller *nurseController) Login(c echo.Context) error {
	var authRequest entities.AuthRequest
	err := c.Bind(&authRequest)

	if err != nil {
		fmt.Println(err.Error())
		switch err.(type) {
		case validator.ValidationErrors:
			var errorMessages string
			for _, e := range err.(validator.ValidationErrors) {
				errorMessage := fmt.Sprintf("Error on field: %s, condition: %s", e.Field(), e.ActualTag())
				errorMessages = fmt.Sprintf(errorMessages + errorMessage)
			}
			c.JSON(
				http.StatusBadRequest,
				entities.ErrorResponse{
					Status:  false,
					Message: errorMessages,
				},
			)
			return nil
		case *json.UnmarshalTypeError:
			c.JSON(http.StatusBadRequest, entities.ErrorResponse{
				Status:  false,
				Message: err.Error(),
			})
			return nil

		default:
			if err == io.EOF {
				c.JSON(http.StatusBadRequest, entities.ErrorResponse{
					Status:  false,
					Message: "Request body is empty",
				})
				return nil
			}
			c.JSON(http.StatusBadRequest, entities.ErrorResponse{
				Status:  false,
				Message: err.Error(),
			})
			return nil
		}
	}

	// Validasi input menggunakan validator
	if err := controller.validator.Struct(authRequest); err != nil {
		var validationErrors []string
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, fmt.Sprintf("%s is %s", err.Field(), err.Tag()))
		}
		return c.JSON(http.StatusBadRequest, entities.ErrorResponse{
			Status:  false,
			Message: validationErrors,
		})
	}

	if !helpers.ValidateNIP(authRequest.NIP) {
		return c.JSON(http.StatusBadRequest, entities.ErrorResponse{
			Status:  false,
			Message: "INVALID NIP FORMAT",
		})
	}
	if strconv.Itoa(authRequest.NIP)[0:3] != "303" {
		return c.JSON(http.StatusNotFound, entities.ErrorResponse{
			Status:  false,
			Message: "INVALID NIP FORMAT",
		})
	}

	tokenString, userData, err := controller.userService.Login(authRequest)
	if err != nil {
		if err.Error() == "INVALID NIP OR PASSWORD" {
			return c.JSON(http.StatusBadRequest, entities.ErrorResponse{
				Status:  false,
				Message: err.Error(),
			})

		}
		fmt.Println("Error : %s", err)
		return c.JSON(http.StatusInternalServerError, entities.ErrorResponse{
			Status:  false,
			Message: "Internal server error",
		})
	}

	return c.JSON(http.StatusOK, entities.SuccessResponse{
		Message: "User login successfully",
		Data: entities.AuthResponse{
			ID:          userData.ID,
			NIP:         userData.NIP,
			Name:        userData.Name,
			AccessToken: tokenString,
		},
	})
}
