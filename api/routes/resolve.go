package routes

import (
	"api-shortner/database"

	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/v2"
)

func ResolveURL(c *fiber.Ctx) error {

	url := c.Params("url")

	r := database.CreateClient(0)

	defer r.Close()

	// Find the short URL in the database
	value, err := r.Get(database.Ctx, url).Result()

	if err == redis.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "short not found in the db!"})
	} else if err != nil {
         return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "cannot connect to db"})
	}

	rInr := database.CreateClient(1);
	defer rInr.Close();
	
	_ = rInr.Iner(database.Ctx, "counter");

	return c.Redirect(value, 301);
}
