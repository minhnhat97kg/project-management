package handler

import (
	"project-management/internal/handler/user"
	"project-management/internal/service"

	"github.com/sirupsen/logrus"
)

type Handler struct {
	User *user.Handler
}

func NewHandler(l *logrus.Logger, s service.Service) *Handler {
	return &Handler{
		User: user.NewUserHandler(l, s.User),
	}
}
