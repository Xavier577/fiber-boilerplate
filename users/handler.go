package users

import "github.com/gofiber/fiber/v2"

func getUser(c *fiber.Ctx) error {
	return c.JSON(map[string]any{"id": 1, "email": "johndoe@mail.com"})
}
