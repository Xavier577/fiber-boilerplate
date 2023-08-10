package main

import (
	"fmt"
	"github.com/Xavier577/fiber-bolierplate/config/env"
	_ "github.com/Xavier577/fiber-bolierplate/docs"
	"github.com/Xavier577/fiber-bolierplate/users"
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"log"
)

// @title Fiberapp Swagger Api
// @version 1.0
// @description This is a sample swagger doc.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2

func setAppRoutes(a *fiber.App) {
	api := a.Group("/api")

	userRouter := api.Group(users.UserRoute)
	users.SetRoutes(userRouter)

}

func setUpSwagger(a *fiber.App) {

	cfg := swagger.Config{
		BasePath: "/api",
		FilePath: "./docs/swagger.json",
	}

	a.Use("/", swagger.New(cfg))
}

func main() {

	app := fiber.New(fiber.Config{
		AppName: "fiberapp",
	})

	app.Use(requestid.New())

	app.Use(cors.New(cors.Config{}))

	app.Use(logger.New(logger.Config{
		Format:     "[${time}] [${locals:requestid}] ${host} - ${latency} | ${status] | ${path} | ${method} | { headers -> {${reqHeaders} } | { body -> ${body} | { res -> ${resBody} }\n",
		TimeFormat: "02-Jan-2006 15:04",
	}))

	setUpSwagger(app)

	setAppRoutes(app)

	ADDRESS := fmt.Sprintf(":%s", env.Get[string]("PORT"))

	err := app.Listen(ADDRESS)

	if err != nil {
		log.Fatal(err)
	}

}
