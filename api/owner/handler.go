package owner

import (
	"net/http"

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

// Add new owner
//
//	@Summary		Add new owner
//	@Description	Add new owner
//	@Tags			owners
//	@Accept			json
//	@Produce		json
//	@Param			owner		body		AddReq	true	"Add request"
//	@Success		200			{object}	api.SingleDataResp[Owner]
//	@Router			/owners   	[post]
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

	o, err := h.Service.Add(c, reqBody)
	if err != nil {
		zlog.Error().Msg(err.Error())

		return api.ErrorResponse(c, http.StatusBadRequest, err)
	}

	zlog.Info().Msg("owner is added")

	return c.JSON(
		http.StatusCreated,
		api.SingleDataResp[Owner]{
			Message: "owner is added",
			Data:    o,
		},
	)
}

// Get all owners by owner id
//
//	@Summary		Get all owners by owner id
//	@Description	Get all owners by owner id
//	@Tags			owners
//	@Accept			json
//	@Produce		json
//	@Param			userId				path		string	true	"owner id"
//	@Param			offset				query		int		true	"offset"	default(1)
//	@Param			limit				query		int		true	"limit"		default(15)
//	@Success		200					{object}	api.MultipleDataResp[Owner]
//	@Router			/owners/{userId} 	[get]
func (h Handler) GetAll(c echo.Context) error {
	owners, err := h.Service.GetAll(c)
	if err != nil {
		zlog.Error().Msg(err.Error())

		return api.ErrorResponse(c, http.StatusBadRequest, err)
	}

	zlog.Info().Msg("retrieve all owners")

	return c.JSON(
		http.StatusOK,
		api.MultipleDataResp[Owner]{
			Message: "retrieve all owners",
			Data:    owners,
		},
	)
}

// Get owner by id
//
//	@Summary		Get owner by id
//	@Description	Get owner by id
//	@Tags			owners
//	@Accept			json
//	@Produce		json
//	@Param			id				path		string	true	"owner id"
//	@Success		200				{object}	api.SingleDataResp[Owner]
//	@Router			/owners/{id} 	[get]
func (h Handler) GetById(c echo.Context) error {
	id := c.Param("id")
	validation := make(map[string][]string)
	if id == "{id}" {
		validation["id"] = append(validation["id"], "cannot empty")

		return api.ValidationResponse(c, http.StatusBadRequest, validation)
	}

	o, err := h.Service.GetById(c, id)
	if err != nil {
		zlog.Error().Msg(err.Error())

		return api.ErrorResponse(c, http.StatusBadRequest, err)
	}

	zlog.Info().Msg("retrieve owner by id")

	return c.JSON(
		http.StatusOK,
		api.SingleDataResp[Owner]{
			Message: "retrieve owner by id",
			Data:    o,
		},
	)
}

// Update owner by id
//
//	@Summary		Update owner by id
//	@Description	Update owner by id
//	@Tags			owners
//	@Accept			json
//	@Produce		json
//	@Param			id				path		string		true	"owner id"
//	@Param			owner			body		UpdateReq	true	"Update request"
//	@Success		200				{object}	api.SingleDataResp[Owner]
//	@Router			/owners/{id} 	[patch]
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

	o, err := h.Service.UpdateById(c, id, reqBody)
	if err != nil {
		zlog.Error().Msg(err.Error())

		return api.ErrorResponse(c, http.StatusBadRequest, err)
	}

	zlog.Info().Msg("update owner by id")

	return c.JSON(
		http.StatusOK,
		api.SingleDataResp[Owner]{
			Message: "update owner by id",
			Data:    o,
		},
	)
}

// Delete owner by id
//
//	@Summary		Delete owner by id
//	@Description	Delete owner by id
//	@Tags			owners
//	@Accept			json
//	@Produce		json
//	@Param			id				path		string	true	"owner id"
//	@Success		200				{object}	api.SingleDataResp[Owner]
//	@Router			/owners/{id} 	[delete]
func (h Handler) DeleteById(c echo.Context) error {
	id := c.Param("id")
	validation := make(map[string][]string)
	if id == "{id}" {
		validation["id"] = append(validation["id"], "cannot empty")

		return api.ValidationResponse(c, http.StatusBadRequest, validation)
	}

	o, err := h.Service.DeleteById(c, id)
	if err != nil {
		zlog.Error().Msg(err.Error())

		return api.ErrorResponse(c, http.StatusBadRequest, err)
	}

	zlog.Info().Msg("delete owner by id")

	return c.JSON(
		http.StatusOK,
		api.SingleDataResp[Owner]{
			Message: "delete owner by id",
			Data:    o,
		},
	)
}
