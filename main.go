package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v3"
)

func CustomErrorHandler(c fiber.Ctx, err error) error {
	code := http.StatusBadRequest
	res := map[string]any{
		"code":         code,
		"errorHandler": "defaultErrorHandler",
		"message":      err.Error(),
	}

	// Send a custom error response
	return c.Status(code).JSON(res)
}

func getUser(c fiber.Ctx) error {
	// It always fails
	return errors.New("some error occurred")
}

func userRoutes() *fiber.App {
	app := fiber.New()
	app.Get("/", getUser)

	return app
}

func V1Routes() *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: CustomErrorHandler,
	})

	app.Use("/users", userRoutes())
	return app
}

func main() {
	port := 8080
	app := fiber.New()
	app.Use("/api/v1", V1Routes())

	fmt.Printf("running on port %d...\n", port)

	err := app.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	// curl http://localhost:8080/api/v1/users
}
