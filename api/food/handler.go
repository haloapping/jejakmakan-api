package food

import (
	"math"
	"net/http"
	"strconv"

	"github.com/haloapping/jejakmakan-api/api"
	"github.com/labstack/echo/v4"
	zlog "github.com/rs/zerolog/log"
)

type Handler struct {
	Service
}

func NewHandler(s Service) Handler {
	return Handler{
		Service: s,
	}
}

// Add new food
//
//	@Summary		Add new food
//	@Description	Add new food
//	@Tags			foods
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			food		body		AddReq	true	"Add request"
//	@Success		200			{object}	api.SingleDataResp[AddFood]
//	@Router			/foods   	[post]
func (h Handler) Add(c echo.Context) error {
	var reqBody AddReq
	err := c.Bind(&reqBody)
	if err != nil {
		zlog.Error().Msg(err.Error())

		return api.ErrorResponse(c, http.StatusBadRequest, err)
	}

	validation := AddValidation(reqBody)
	if len(validation) > 0 {
		zlog.Info().Interface("validation", validation).Msg("validation")

		return api.ValidationResponse(c, http.StatusBadRequest, validation)
	}

	f, err := h.Service.Add(c, reqBody)
	if err != nil {
		zlog.Error().Msg(err.Error())

		return api.ErrorResponse(c, http.StatusBadRequest, err)
	}

	zlog.Info().Msg("food is added")

	return c.JSON(
		http.StatusCreated,
		api.SingleDataResp[AddFood]{
			Message: "food is added",
			Data:    f,
		},
	)
}

// Get all foods
//
//	@Summary		Get all foods
//	@Description	Get all foods
//	@Tags			foods
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			limit	query		int	true	"limit"		default(15)
//	@Param			offset	query		int	true	"offset"	default(1)
//	@Success		200		{object}	api.MultipleDataResp[Food]
//	@Router			/foods 																																					[get]
func (h Handler) GetAll(c echo.Context) error {
	validation := make(map[string][]string)
	offset := c.QueryParam("offset")
	if offset == "" {
		validation["offset"] = append(validation["offset"], "cannot empty")
	}
	limit := c.QueryParam("limit")
	if limit == "" {
		validation["limit"] = append(validation["limit"], "cannot empty")
	}
	if len(validation) > 0 {
		zlog.Info().Interface("validation", validation).Msg("validation")

		return api.ValidationResponse(c, http.StatusBadRequest, validation)
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		zlog.Error().Msg(err.Error())

		return api.ErrorResponse(c, http.StatusBadRequest, err)
	}
	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		zlog.Error().Msg(err.Error())

		return api.ErrorResponse(c, http.StatusBadRequest, err)
	}

	l, total, err := h.Service.GetAll(c, limitInt, offsetInt)
	if err != nil {
		zlog.Error().Msg(err.Error())

		return api.ErrorResponse(c, http.StatusBadRequest, err)
	}

	zlog.Info().Msg("retrieve all foods")

	return c.JSON(
		http.StatusCreated,
		api.MultipleDataResp[Food]{
			Message: "retrieve all foods",
			Pagination: api.Pagination{
				Page:      (offsetInt / limitInt) + 1,
				PageSize:  limitInt,
				TotalPage: int(math.Ceil(float64(total)) / float64(limitInt)),
				TotalItem: total,
			},
			Data: l,
		},
	)
}

// Get food by id
//
//	@Summary		Get food by id
//	@Description	Get food by id
//	@Tags			foods
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id				path		string	true	"food id"
//	@Success		200				{object}	api.SingleDataResp[Food]
//	@Router			/foods/{id} 	[get]
func (h Handler) GetById(c echo.Context) error {
	id := c.Param("id")
	validation := make(map[string][]string)
	if id == "{id}" {
		validation["id"] = append(validation["id"], "cannot empty")

		return api.ValidationResponse(c, http.StatusBadRequest, validation)
	}

	f, err := h.Service.GetById(c, id)
	if err != nil {
		zlog.Error().Msg(err.Error())

		return api.ErrorResponse(c, http.StatusBadRequest, err)
	}

	zlog.Info().Msg("retrieve food by id")

	return c.JSON(
		http.StatusOK,
		api.SingleDataResp[Food]{
			Message: "retrieve food by id",
			Data:    f,
		},
	)
}

// Update food by id
//
//	@Summary		Update food by id
//	@Description	Update food by id
//	@Tags			foods
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id				path		string		true	"food id"
//	@Param			food			body		UpdateReq	true	"Update request"
//	@Success		200				{object}	api.SingleDataResp[Food]
//	@Router			/foods/{id} 	[patch]
func (h Handler) UpdateById(c echo.Context) error {
	id := c.Param("id")
	validation := make(map[string][]string)
	if id == "{id}" {
		validation["id"] = append(validation["id"], "cannot empty")

		return api.ValidationResponse(c, http.StatusBadRequest, validation)
	}

	var reqBody UpdateReq
	err := c.Bind(&reqBody)
	if err != nil {
		zlog.Error().Msg(err.Error())

		return api.ErrorResponse(c, http.StatusBadRequest, err)
	}

	f, err := h.Service.UpdateById(c, id, reqBody)
	if err != nil {
		zlog.Error().Msg(err.Error())

		return api.ErrorResponse(c, http.StatusBadRequest, err)
	}

	zlog.Info().Msg("update food by id")

	return c.JSON(
		http.StatusOK,
		api.SingleDataResp[Food]{
			Message: "update food by id",
			Data:    f,
		},
	)
}

// Delete food by id
//
//	@Summary		Delete food by id
//	@Description	Delete food by id
//	@Tags			foods
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id				path		string	true	"food id"
//	@Success		200				{object}	api.SingleDataResp[Food]
//	@Router			/foods/{id} 	[delete]
func (h Handler) DeleteById(c echo.Context) error {
	id := c.Param("id")
	validation := make(map[string][]string)
	if id == "{id}" {
		validation["id"] = append(validation["id"], "cannot empty")

		return api.ValidationResponse(c, http.StatusBadRequest, validation)
	}

	f, err := h.Service.DeleteById(c, id)
	if err != nil {
		zlog.Error().Msg(err.Error())

		return api.ErrorResponse(c, http.StatusBadRequest, err)
	}

	zlog.Info().Msg("delete food by id")

	return c.JSON(
		http.StatusOK,
		api.SingleDataResp[Food]{
			Message: "delete food by id",
			Data:    f,
		},
	)
}
