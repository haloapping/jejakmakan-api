package main

import (
	"github.com/haloapping/jejakmakan-api/api/food"
	"github.com/haloapping/jejakmakan-api/api/location"
	"github.com/haloapping/jejakmakan-api/api/owner"
	"github.com/haloapping/jejakmakan-api/api/user"
	"github.com/haloapping/jejakmakan-api/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

func Router(pool *pgxpool.Pool, r *echo.Echo) {
	userRoute := r.Group("/users")
	userRepo := user.NewRepository(pool)
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)
	user.Router(userRoute, userHandler)

	ownerRoute := r.Group("/owners")
	ownerRoute.Use(middleware.JWTAuth)
	ownerRepo := owner.NewRepository(pool)
	ownerService := owner.NewService(ownerRepo)
	ownerHandler := owner.NewHandler(ownerService)
	owner.Router(ownerRoute, ownerHandler)

	locationRoute := r.Group("/locations")
	locationRoute.Use(middleware.JWTAuth)
	locationRepo := location.NewRepository(pool)
	locationService := location.NewService(locationRepo)
	locationHandler := location.NewHandler(locationService)
	location.Router(locationRoute, locationHandler)

	foodRoute := r.Group("/foods")
	foodRoute.Use(middleware.JWTAuth)
	foodRepo := food.NewRepository(pool)
	foodService := food.NewService(foodRepo)
	foodHandler := food.NewHandler(foodService)
	food.Router(foodRoute, foodHandler)
}
