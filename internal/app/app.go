package app

import (
	"context"
	"fmt"

	"project-management/internal/handler"
	"project-management/internal/repository"
	"project-management/internal/service"
	"project-management/pkg/config"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

type ApplicationContext struct {
	server fiber.App
	conf   config.Config
	log    *log.Logger
}

func New(ctx context.Context) (*ApplicationContext, error) {
	conf, _ := config.ReadConfig("config/dev.yaml")
	log := log.New()
	server := fiber.New(fiber.Config{
		// ErrorHandler: func(c *fiber.Ctx, err error) error {
		// 	return c.Status(200).SendString("Error handler")
		// },
	})
	repositories := repository.NewRepository(log, nil)
	services := service.NewService(log, *repositories)
	handlers := handler.NewHandler(log, *services)

	NewRoute(server, handlers)

	return &ApplicationContext{
		server: *server,
		conf:   conf,
		log:    log,
	}, nil
}

func (app *ApplicationContext) Start() {
	app.log.Fatal(app.server.Listen(fmt.Sprintf(":%d", app.conf.Port)))
}

func (app *ApplicationContext) Stop() {
	app.log.Info("Server is stopping....")
}
