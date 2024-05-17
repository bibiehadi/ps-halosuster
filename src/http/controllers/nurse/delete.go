package nursecontroller

import (
	"halosuster/src/entities"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (controller *nurseController) Delete(c echo.Context) error {
	userId := c.Param("id")
	err := controller.userService.Delete(userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entities.ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, entities.SuccessResponse{
		Message: "Nurse deleted successfull",
		Data: entities.NurseResponse{
			ID: userId,
		},
	})
}
