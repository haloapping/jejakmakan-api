package main

import (
	"fmt"
	"net/http"

	"github.com/MarceloPetrucio/go-scalar-api-reference"
	"github.com/haloapping/jejakmakan-api/config"
	"github.com/haloapping/jejakmakan-api/db"
	customMiddleware "github.com/haloapping/jejakmakan-api/middleware"
	"github.com/labstack/echo/v4"
)

// @securityDefinitions.apikey	BearerAuth
// @in							header
// @name						Authorization
func main() {
	// console and file log
	logFile, err := customMiddleware.MultiLogger()
	if err != nil {
		panic(err)
	}
	defer logFile.Close()

	// setup config
	dbconfig, err := db.NewDBConfig(db.ConnDBStr())
	if err != nil {
		panic(err)
	}

	// initiate database pooling
	pool, err := db.NewDBPool(dbconfig)
	if err != nil {
		panic(err)
	}
	defer pool.Close()

	// routing
	r := echo.New()
	customMiddleware.EchoLogger(r)
	Router(pool, r)

	// public route
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

	r.Start(config.APIUrl(".env"))
}
