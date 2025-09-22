package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"

	"github.com/MarceloPetrucio/go-scalar-api-reference"
	"github.com/haloapping/jejakmakan-api/db"
	customMiddleware "github.com/haloapping/jejakmakan-api/middleware"
	"github.com/labstack/echo/v4"
	zlog "github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

// @securityDefinitions.apikey	BearerAuth
// @in							header
// @name						Authorization
func main() {
	// load env
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.SetConfigName("dev.config")
	err := viper.ReadInConfig()
	if err != nil {
		zlog.Error().Msg(err.Error())
	}
	viper.AutomaticEnv()

	// console and file log
	logFile, err := customMiddleware.MultiLogger()
	if err != nil {
		panic(err)
	}
	defer logFile.Close()

	// setup config
	connStr := fmt.Sprintf(
		"postgresql://%s:%s@%s:%d/%s?sslmode=%s",
		viper.GetString("DB_USER"),
		viper.GetString("DB_PASSWORD"),
		viper.GetString("DB_HOST"),
		viper.GetInt("DB_PORT"),
		viper.GetString("DB_NAME"),
		viper.GetString("DB_SSLMODE"),
	)
	dbconfig, err := db.NewDBConfig(connStr)
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

	r.Start(fmt.Sprintf(":%d", viper.GetInt("APP_PORT")))
}
