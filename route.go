package main

import (
	"github.com/haloapping/jejakmakan-api/api/food"
	"github.com/haloapping/jejakmakan-api/api/location"
	"github.com/haloapping/jejakmakan-api/api/owner"
	"github.com/haloapping/jejakmakan-api/api/user"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

func Router(pool *pgxpool.Pool, r *echo.Echo) {
	userRepo := user.NewRepository(pool)
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)
	user.Router(r.Group("/users"), userHandler)

	ownerRepo := owner.NewRepository(pool)
	ownerService := owner.NewService(ownerRepo)
	ownerHandler := owner.NewHandler(ownerService)
	owner.Router(r.Group("/owners"), ownerHandler)

	locationRepo := location.NewRepository(pool)
	locationService := location.NewService(locationRepo)
	locationHandler := location.NewHandler(locationService)
	location.Router(r.Group("/locations"), locationHandler)

	foodRepo := food.NewRepository(pool)
	foodService := food.NewService(foodRepo)
	foodHandler := food.NewHandler(foodService)
	food.Router(r.Group("/foods"), foodHandler)
}
