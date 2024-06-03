package ctlusers

import (
	ctlwallets "github.com/diazharizky/buddi-backend/internal/controllers/v1/wallets"
	"github.com/gofiber/fiber/v2"
)

func NewRouter() (router *fiber.App) {
	router = fiber.New()
	router.Get("/", nil)
	router.Post("/", nil)

	userPath := router.Group("/:user_id")
	userPath.Mount("/wallets", ctlwallets.NewRouter())

	return
}
