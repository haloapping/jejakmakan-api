package api

import (
	"github.com/labstack/echo/v4"
)

type ValidationResp struct {
	Validation map[string][]string `json:"validation" binding:"required" extensions:"x-order=1"`
}

type ErrorResp struct {
	Error string `json:"error" binding:"required" extensions:"x-order=1"`
}

type Pagination struct {
	Page      int `json:"page" binding:"required" extensions:"x-order=1"`
	PageSize  int `json:"pageSize" binding:"required" extensions:"x-order=2"`
	TotalPage int `json:"totalPage" binding:"required" extensions:"x-order=3"`
	TotalItem int `json:"totalItem" binding:"required" extensions:"x-order=4"`
}

type SingleDataResp[data any] struct {
	Message string `json:"message" binding:"required" extensions:"x-order=1"`
	Data    data   `json:"data" binding:"required" extensions:"x-order=2"`
}

type MultipleDataResp[data any] struct {
	Message    string     `json:"message" binding:"required" extensions:"x-order=1"`
	Pagination Pagination `json:"pagination" binding:"required" extensions:"x-order=2"`
	Data       []data     `json:"data" binding:"required" extensions:"x-order=2"`
}

func ErrorResponse(c echo.Context, status int, err error) error {
	return c.JSON(
		status,
		ErrorResp{
			Error: err.Error(),
		},
	)
}

func ValidationResponse(c echo.Context, status int, validation map[string][]string) error {
	return c.JSON(
		status,
		ValidationResp{
			Validation: validation,
		},
	)
}
