package users

import "github.com/gofiber/fiber/v2"

const UserRoute = "/users"

func SetRoutes(r fiber.Router) {
	r.Use("/", getUser)

}
