package nursecontroller

import (
	"halosuster/src/entities"
	"net/http"
	"reflect"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (controller *nurseController) GetAll(c echo.Context) error {
	params := entities.UserQueryParams{}

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

	if id := c.QueryParam("id"); id != "" {
		params.Id = id
	}

	if name := c.QueryParam("name"); name != "" {
		params.Name = name
	}

	if NIP := c.QueryParam("nip"); NIP != "" {
		params.NIP = NIP
	}

	if role := c.QueryParam("role"); role != "" {
		params.Role = role
	}

	if createdAt := c.QueryParam("createdAt"); createdAt != "" {
		params.CreatedAt = createdAt
	}

	users, err := controller.userService.GetAll(params)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, entities.ErrorResponse{
			Status:  false,
			Message: "Failed to fetch products",
		})
	}

	if users == nil || reflect.ValueOf(users).IsNil() {
		return c.JSON(http.StatusOK, entities.SuccessResponse{
			Message: "success",
			Data:    []entities.UserResponse{},
		})
	}

	return c.JSON(http.StatusOK, entities.SuccessResponse{
		Message: "success",
		Data:    users,
	})
}
