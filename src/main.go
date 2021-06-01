package main

import (
	"os"
	"strconv"

	fiber "github.com/gofiber/fiber/v2"
)

type Result struct {
	Value int `json:"value"`
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Get("/:operation/:num1/:num2", OperationHandler)

	port := os.Getenv("PORT")
	app.Listen(":" + port)
}

func OperationHandler(c *fiber.Ctx) error {
	operation := c.Params("operation")
	num1, _ := strconv.Atoi(c.Params("num1"))
	num2, _ := strconv.Atoi(c.Params("num2"))

	result := Result{}

	switch operation {
	case "sum":
		result.Value = add(num1, num2)
		return c.JSON(result)
	case "multiply":
		result.Value = multiply(num1, num2)
		return c.JSON(result)
	default:
		return c.SendFile("./views/404.html")
	}

}

func add(n1, n2 int) int {
	return n1 + n2
}

func multiply(n1 int, n2 int) int {
	return n1 * n2
}
