package owner

import "github.com/labstack/echo/v4"

func Router(g *echo.Group, h Handler) {
	g.POST("", h.Add)
	g.GET("", h.GetAll)
	g.GET("/:id", h.GetById)
	g.PATCH("/:id", h.UpdateById)
	g.DELETE("/:id", h.DeleteById)
}
