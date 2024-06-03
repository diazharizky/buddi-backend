package ctlwallets

import (
	ctltransactions "github.com/diazharizky/buddi-backend/internal/controllers/v1/transactions"
	"github.com/gofiber/fiber/v2"
)

func NewRouter() (router *fiber.App) {
	router = fiber.New()
	router.Get("/", nil)
	router.Post("/", nil)

	walletPath := router.Group("/:wallet_id")
	walletPath.Mount("/transactions", ctltransactions.NewRouter())

	return
}
