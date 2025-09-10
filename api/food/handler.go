package food

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

// Add new food
//
//	@Summary		Add new food
//	@Description	Add new food
//	@Tags			foods
//	@Accept			json
//	@Produce		json
//	@Param			food		body		AddReq	true	"Add request"
//	@Success		200			{object}	api.SingleDataResp[Food]
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
		api.SingleDataResp[Food]{
			Message: "food is added",
			Data:    f,
		},
	)
}
