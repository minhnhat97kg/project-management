package service

import (
	"project-management/internal/repository"
	"project-management/internal/service/user"

	"github.com/sirupsen/logrus"
)

type Service struct {
	User *user.Service
}

func NewService(l *logrus.Logger, r repository.Repository) *Service {
	return &Service{
		User: user.NewUserService(l, r.User),
	}

}
