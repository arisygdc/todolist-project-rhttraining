package core

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type errJson struct {
	Status string `json:"status"`
	Msg    string `json:"message"`
}

type createJson struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func ResponseBadRequest(c echo.Context, msg string) error {
	httpCode := http.StatusBadRequest

	return c.JSON(httpCode, errJson{
		Status: http.StatusText(httpCode),
		Msg:    msg,
	})
}

func ResponseCreated(c echo.Context, data interface{}) error {
	httpCode := http.StatusCreated

	return c.JSON(httpCode, createJson{
		Status: http.StatusText(httpCode),
		Data:   data,
	})
}

func ResponseOk(c echo.Context, data interface{}) error {
	httpCode := http.StatusOK

	return c.JSON(httpCode, createJson{
		Status: http.StatusText(httpCode),
		Data:   data,
	})
}
