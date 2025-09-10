package food

import "github.com/labstack/echo/v4"

func Router(g *echo.Group, h Handler) {
	g.POST("", h.Add)
}
