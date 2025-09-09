package main

import (
	"fmt"
	"net/http"

	"github.com/MarceloPetrucio/go-scalar-api-reference"
	"github.com/haloapping/jejakmakan-api/api/owner"
	"github.com/haloapping/jejakmakan-api/api/user"
	"github.com/haloapping/jejakmakan-api/config"
	"github.com/haloapping/jejakmakan-api/db"
	customMiddleware "github.com/haloapping/jejakmakan-api/middleware"
	"github.com/labstack/echo/v4"
)

func main() {
	// console and file log
	logFile, err := customMiddleware.MultiLog()
	if err != nil {
		panic(err)
	}
	defer logFile.Close()

	// initiate database pooling
	connStr := config.DBConnStr(".env")
	pool := db.NewConnection(connStr)
	defer pool.Close()

	// routing
	r := echo.New()
	customMiddleware.EchoLogger(r)

	userRepo := user.NewRepository(pool)
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)
	user.Router(r.Group("/users"), userHandler)

	ownerRepo := owner.NewRepository(pool)
	ownerService := owner.NewService(ownerRepo)
	ownerHandler := owner.NewHandler(ownerService)
	owner.Router(r.Group("/owners"), ownerHandler)

	r.GET("/", func(c echo.Context) error {
		htmlContent, err := scalar.ApiReferenceHTML(&scalar.Options{
			SpecURL: "./docs/swagger.json",
			CustomOptions: scalar.CustomOptions{
				PageTitle: "Jejak Makan API",
			},
			DarkMode:   true,
			Theme:      scalar.ThemeDeepSpace,
			HideModels: true,
		})

		if err != nil {
			fmt.Printf("%v", err)
		}

		return c.HTML(http.StatusOK, htmlContent)
	})

	r.Start(":3000")
}
