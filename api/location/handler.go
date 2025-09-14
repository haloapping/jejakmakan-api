package location

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

// Add new location
//
//	@Summary		Add new location
//	@Description	Add new location
//	@Tags			locations
//	@Accept			json
//	@Produce		json
//	@Param			location		body		AddReq	true	"Add request"
//	@Success		200				{object}	api.SingleDataResp[Location]
//	@Router			/locations   	[post]
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

	l, err := h.Service.Add(c, reqBody)
	if err != nil {
		zlog.Error().Msg(err.Error())

		return api.ErrorResponse(c, http.StatusBadRequest, err)
	}

	zlog.Info().Msg("location is added")

	return c.JSON(
		http.StatusCreated,
		api.SingleDataResp[Location]{
			Message: "location is added",
			Data:    l,
		},
	)
}

// Get all locations
//
//	@Summary		Get all locations
//	@Description	Get all locations
//	@Tags			locations
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			limit		query		int	true	"limit"		default(15)
//	@Param			offset		query		int	true	"offset"	default(1)
//	@Success		200			{object}	api.MultipleDataResp[Location]
//	@Router			/locations 																																					[get]
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

	zlog.Info().Msg("retrieve all locations")

	return c.JSON(
		http.StatusCreated,
		api.MultipleDataResp[Location]{
			Message: "retrieve all locations",
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

// Get location by id
//
//	@Summary		Get location by id
//	@Description	Get location by id
//	@Tags			locations
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id					path		string	true	"location id"
//	@Success		200					{object}	api.SingleDataResp[Location]
//	@Router			/locations/{id} 	[get]
func (h Handler) GetById(c echo.Context) error {
	id := c.Param("id")
	validation := make(map[string][]string)
	if id == "{id}" {
		validation["id"] = append(validation["id"], "cannot empty")

		return api.ValidationResponse(c, http.StatusBadRequest, validation)
	}

	l, err := h.Service.GetById(c, id)
	if err != nil {
		zlog.Error().Msg(err.Error())

		return api.ErrorResponse(c, http.StatusBadRequest, err)
	}

	zlog.Info().Msg("retrieve location by id")

	return c.JSON(
		http.StatusOK,
		api.SingleDataResp[Location]{
			Message: "retrieve location by id",
			Data:    l,
		},
	)
}

// Update location by id
//
//	@Summary		Update location by id
//	@Description	Update location by id
//	@Tags			locations
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id					path		string		true	"location id"
//	@Param			location			body		UpdateReq	true	"Update request"
//	@Success		200					{object}	api.SingleDataResp[Location]
//	@Router			/locations/{id} 	[patch]
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

	l, err := h.Service.UpdateById(c, id, reqBody)
	if err != nil {
		zlog.Error().Msg(err.Error())

		return api.ErrorResponse(c, http.StatusBadRequest, err)
	}

	zlog.Info().Msg("update location by id")

	return c.JSON(
		http.StatusOK,
		api.SingleDataResp[Location]{
			Message: "update location by id",
			Data:    l,
		},
	)
}

// Delete location by id
//
//	@Summary		Delete location by id
//	@Description	Delete location by id
//	@Tags			locations
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id					path		string	true	"location id"
//	@Success		200					{object}	api.SingleDataResp[Location]
//	@Router			/locations/{id} 	[delete]
func (h Handler) DeleteById(c echo.Context) error {
	id := c.Param("id")
	validation := make(map[string][]string)
	if id == "{id}" {
		validation["id"] = append(validation["id"], "cannot empty")

		return api.ValidationResponse(c, http.StatusBadRequest, validation)
	}

	l, err := h.Service.DeleteById(c, id)
	if err != nil {
		zlog.Error().Msg(err.Error())

		return api.ErrorResponse(c, http.StatusBadRequest, err)
	}

	zlog.Info().Msg("delete location by id")

	return c.JSON(
		http.StatusOK,
		api.SingleDataResp[Location]{
			Message: "delete location by id",
			Data:    l,
		},
	)
}
