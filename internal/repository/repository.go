package repository

import (
	user "project-management/internal/repository/user"

	"github.com/sirupsen/logrus"
)

type Repository struct {
	User user.Repository
}

func NewRepository(l *logrus.Logger, db interface{}) *Repository {
	return &Repository{
		User: *user.NewUserRepository(l, db),
	}
}
