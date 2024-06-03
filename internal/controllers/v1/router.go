package ctlv1

import (
	ctlusers "github.com/diazharizky/buddi-backend/internal/controllers/v1/users"
	"github.com/gofiber/fiber/v2"
)

func NewRouter() (router *fiber.App) {
	router = fiber.New()
	router.Mount("/users", ctlusers.NewRouter())

	return
}
