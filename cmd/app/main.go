package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/redanthrax/mspinventory/db"
	"github.com/redanthrax/mspinventory/handlers"
	"github.com/redanthrax/mspinventory/util"
)

func main() {
  if err := godotenv.Load(".env.local"); err != nil {
    log.Fatal(err)
  }

  db.Init()

  app := fiber.New(fiber.Config{
    ErrorHandler: handlers.ErrorHandler,
    DisableStartupMessage: true,
    PassLocalsToViews: true,
    Views: setupEngine(),
  })

  initRoutes(app)
  listenAddress := os.Getenv("HTTP_LISTEN_ADDR")
  log.Printf("App running in %s and listening on %s\n",
    util.AppEnv(),
    listenAddress)
  log.Fatal(app.Listen(listenAddress))
}

func initRoutes(app *fiber.App) {
  app.Static("/public", "./public")
  app.Get("/", handlers.HandleHome)
  app.Use(handlers.NotFoundMiddleware)
}
