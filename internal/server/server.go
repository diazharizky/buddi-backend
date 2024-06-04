package server

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/diazharizky/buddi-backend/config"
	ctlv1 "github.com/diazharizky/buddi-backend/internal/controllers/v1"
	"github.com/diazharizky/buddi-backend/pkg/apiresp"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

type server struct {
	engine *fiber.App
}

func init() {
	config.Global.SetDefault("app.host", "localhost")
	config.Global.SetDefault("app.port", 8080)
	config.Global.SetDefault("app.throttling.max.requests", 20)
	config.Global.SetDefault("app.throttling.expiration", 30)
}

func New() (svr server) {
	eng := fiber.New(fiber.Config{
		CaseSensitive: true,
	})

	eng.Use(cors.New())
	eng.Use(limiter.New(limiter.Config{
		// This rate limiter might not be working
		// when the app is running in multiple nodes.
		// Probably we need to utilise in-memory database
		// to store the metrics.

		Max:        config.Global.GetInt("app.throttling.max.requests"),
		Expiration: config.Global.GetDuration("app.throttling.expiration") * time.Second,
		LimitReached: func(fcx *fiber.Ctx) error {
			resp := apiresp.CommonError(
				errors.New("too many requests"),
			)

			return fcx.
				Status(http.StatusTooManyRequests).
				JSON(resp)
		},
	}))

	eng.Get("/healthcheck", func(fcx *fiber.Ctx) error {
		return fcx.
			Status(http.StatusOK).
			JSON(apiresp.Success(nil))
	})

	basePath := eng.Group("/api")
	basePath.Mount("/v1", ctlv1.NewRouter())

	svr.engine = eng

	return
}

func (svr server) Start() {
	addr := fmt.Sprintf("%s:%d",
		config.Global.GetString("app.host"),
		config.Global.GetInt("app.port"),
	)

	svr.engine.Listen(addr)
}
