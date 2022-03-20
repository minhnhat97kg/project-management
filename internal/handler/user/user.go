package user

import (
	"project-management/internal/model"
	"project-management/pkg/validator"

	fiber "github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

type UserService interface {
	Register(user model.RegisterRequest) error
}

type Handler struct {
	log *log.Logger
	srv UserService
}

func NewUserHandler(l *log.Logger, s UserService) *Handler {
	return &Handler{
		log: l,
		srv: s,
	}

}

func (h *Handler) Login(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.SendString(id)
}

func (h *Handler) Register(c *fiber.Ctx) error {
	user := model.RegisterRequest{}
	if err := c.BodyParser(&user); err != nil {
		return c.SendString("Type mismatch" + err.Error())
	}

	if errs := validator.ValidateStruct(user); len(errs) > 0 {
		h.log.Error("Validate error", errs)
		return errs[0]
	}

	if err2 := h.srv.Register(user); err2 != nil {
		return err2
	}

	return c.SendString("Register successfully")
}
