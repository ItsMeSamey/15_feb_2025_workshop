package main

import (
  "errors"
  "net/http"
  "strconv"

  "github.com/gofiber/fiber/v3"
)

var tasks []string

func getHandler(c fiber.Ctx) (err error) { return c.JSON(tasks) }

func addHandler(c fiber.Ctx) (err error) {
  task := string(c.Body())
  tasks = append(tasks, task)
  return c.SendStatus(http.StatusOK)
}

func removeHandler(c fiber.Ctx) (err error) {
  body := c.Body()
  idx, err := strconv.Atoi(string(body))

  if err != nil { return err }

  if len(tasks) < idx { return errors.New("Invalid index") }
  if idx == len(tasks) - 1 {
    tasks = tasks[:idx]
  } else {
    tasks = append(tasks[:idx], tasks[idx+1:]...)
  }

  return c.SendStatus(200)
}

func main() {
  app := fiber.New()
  app.Get("/tasks", getHandler)
  app.Put("/tasks", addHandler)
  app.Delete("/tasks", removeHandler)
  app.Listen(":3000")
}

