package ctltransactions

import (
	"github.com/gofiber/fiber/v2"
)

func NewRouter() (router *fiber.App) {
	router = fiber.New()
	router.Get("/", nil)
	router.Post("/", nil)

	transactionPath := router.Group("/:transaction_id")
	transactionPath.Get("/", nil)
	transactionPath.Patch("/", nil)

	return
}
