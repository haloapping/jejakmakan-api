package user

import (
	"net/http"

	"github.com/haloapping/jejakmakan-api/api"
	"github.com/haloapping/jejakmakan-api/jwt"
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

// Register user
//
//	@Summary		Register user
//	@Description	Register user
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			register			user		body	UserRegisterReq	true	"Register user request"
//	@Success		200					{object}	api.SingleDataResp[UserRegister]
//	@Router			/users/register   	[post]
func (h Handler) Register(c echo.Context) error {
	var reqBody UserRegisterReq
	err := c.Bind(&reqBody)
	if err != nil {
		zlog.Error().Msg(err.Error())

		return api.ErrorResponse(c, http.StatusBadRequest, err)
	}

	validation := RegisterValidation(reqBody)
	if len(validation) > 0 {
		zlog.Info().Interface("validation", validation).Msg("validation")

		return api.ValidationResponse(c, http.StatusBadRequest, validation)
	}

	ur, err := h.Service.Register(c, reqBody)
	if err != nil {
		zlog.Error().Msg(err.Error())

		return api.ErrorResponse(c, http.StatusBadRequest, err)
	}

	zlog.Info().Msg("registration is successfully")

	return c.JSON(
		http.StatusCreated,
		api.SingleDataResp[UserRegister]{
			Message: "registration is successfully",
			Data:    ur,
		},
	)
}

// Login user
//
//	@Summary		Login user
//	@Description	Login user
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			login			user		body	UserLoginReq	true	"Login user request"
//	@Success		200				{object}	api.SingleDataResp[UserLogin]
//	@Router			/users/login   	[post]
func (h *Handler) Login(c echo.Context) error {
	var reqBody UserLoginReq
	err := c.Bind(&reqBody)
	if err != nil {
		zlog.Error().Msg(err.Error())

		return api.ErrorResponse(c, http.StatusBadRequest, err)
	}

	validation := LoginValidation(reqBody)
	if len(validation) > 0 {
		zlog.Info().Interface("validation", validation).Msg("validation")

		return api.ValidationResponse(c, http.StatusBadRequest, validation)
	}

	ul, err := h.Service.Login(c, reqBody)
	if err != nil {
		zlog.Error().Msg(err.Error())

		return api.ErrorResponse(c, http.StatusBadRequest, err)
	}

	token, err := jwt.GenerateToken(ul.Id, ul.Username)
	if err != nil {
		zlog.Error().Msg(err.Error())

		return api.ErrorResponse(c, http.StatusBadRequest, err)
	}

	zlog.Info().Msg("login is successfully")

	return c.JSON(
		http.StatusOK,
		map[string]string{
			"token": token,
		},
	)
}
