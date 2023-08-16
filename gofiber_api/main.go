package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

type Person struct {
	Name string
	Age  int
}

func getPerson(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(person)
}

var person Person

func createPerson(ctx *fiber.Ctx) error {
	body := new(Person)
	err := ctx.BodyParser(body)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
		return err
	}

	person = Person{
		Name: body.Name,
		Age:  body.Age,
	}
	return ctx.Status(fiber.StatusOK).JSON(person)
}
func main() {
	app := fiber.New()

	// Call Use method before setting up any endpoints
	app.Use(logger.New())
	app.Use(requestid.New())
	personApp := app.Group("/person")

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello World!")
	})
	personApp.Get("", getPerson)
	personApp.Post("", createPerson)

	app.Listen(":80")
}
