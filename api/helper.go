package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type jsonResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

var validate = validator.New()

func (app *Config) validateBody(c *gin.Context, data any) error {
	if err := c.ShouldBindJSON(&data); err != nil {
		app.ErrorResponse(c, err)
		return err
	}
	if err := validate.Struct(data); err != nil {
		app.ErrorResponse(c, err)
		return err
	}
	return nil
}

func (app *Config) SuccessResponse(c *gin.Context, status int, data any) {
	c.JSON(status, jsonResponse{
		Status:  status,
		Message: "success",
		Data:    data,
	})
}

func (app *Config) ErrorResponse(c *gin.Context, err error, status ...int) {
	statusKode := http.StatusBadRequest
	if len(status) > 0 {
		statusKode = status[0]
	}
	c.JSON(statusKode, jsonResponse{
		Status:  statusKode,
		Message: err.Error(),
	})
}
