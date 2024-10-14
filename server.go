package main

import (
	"Mereb3/config"
	"Mereb3/routes"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Errorf("Error Happend Loding Env")
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "5000"
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(config.CorsConfig)
	routes.PersonRoutes(e)

	e.Logger.Fatal(e.Start(":" + port))
}
