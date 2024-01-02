package main

import (
	"restApiFiber/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	//use router movies
	routes.UseMoviesRoutes(app)

	app.Listen(":4000")
}


