package main

import (
	"fmt"
	"github.com/Xavier577/fiber-bolierplate/config/env"
	"github.com/Xavier577/fiber-bolierplate/users"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"log"
)

func setAppRoutes(a *fiber.App) {
	api := a.Group("/api")

	userRouter := api.Group(users.UserRoute)
	users.SetRoutes(userRouter)

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

	setAppRoutes(app)

	ADDRESS := fmt.Sprintf(":%s", env.Get[string]("PORT"))

	err := app.Listen(ADDRESS)

	if err != nil {
		log.Fatal(err)
	}

}
