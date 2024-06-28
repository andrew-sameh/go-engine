package server

import (
	"github.com/gofiber/fiber/v2"

	"engine/internal/database"
)

type FiberServer struct {
	*fiber.App

	db database.Service
}

func New() *FiberServer {
	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "engine",
			AppName:      "engine",
		}),

		db: database.New(),
	}

	return server
}
