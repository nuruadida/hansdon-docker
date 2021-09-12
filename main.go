package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"os"
)

func main(){

	app := fiber.New()

	nama := os.Getenv("NAMA")

	app.Get("/", func(d *fiber.Ctx) error {
		return d.SendString("Hello, " + nama)
	})

	app.Listen(":3002")
	fmt.Println("Hello" + nama)
}
