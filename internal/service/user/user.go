package user

import (
	"project-management/internal/model"

	"github.com/sirupsen/logrus"
)

type Repository interface {
	CreateNewUser(user model.RegisterRequest) error
}

type Service struct {
	log *logrus.Logger
	r   Repository
}

func NewUserService(l *logrus.Logger, r Repository) *Service {
	return &Service{
		log: l,
		r:   r,
	}
}

func (srv Service) Register(user model.RegisterRequest) error {
	return srv.r.CreateNewUser(user)
}
