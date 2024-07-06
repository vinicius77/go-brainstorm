package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	// "github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"github.com/vinicius77/go-brainstorm/internal/handlers"
)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
	Body  string `json:"body"`
}

func main() {
	/**
	=== Uncomment that snippet to test Fiber ======

	PS.: For obvious reasons running Fiber and Chi concurrently
	triggers crazy stuff

	log.SetReportCaller(true)
	app := fiber.New()

	todos := []Todo{}

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo := &Todo{}
		err := c.BodyParser(todo)

		if err != nil {
			return err
		}

		todo.ID = len(todos) + 1
		todos = append(todos, *todo)
		return c.JSON(todos)
	})

	app.Patch("api/todos/:id/done", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")

		if err != nil {
			return c.Status(404).SendString("Id not found")
		}

		for index, todo := range todos {
			if todo.ID == id {
				todos[index].Done = true
				break
			}
		}

		return c.JSON(todos)
	})

	app.Get("api/todos", func(c *fiber.Ctx) error {
		return c.JSON(todos)
	})

	log.Fatal(app.Listen(":8000"))

		=== Uncomment that snippet to test Fiber ======
	*/

	var router *chi.Mux = chi.NewRouter()
	handlers.Handler(router)

	fmt.Println("Starting API service ...")
	fmt.Println("Server listening at http://localhost:8000 ...")

	err := http.ListenAndServe(":8000", router)

	if err != nil {
		log.Error(err)
	}

}
