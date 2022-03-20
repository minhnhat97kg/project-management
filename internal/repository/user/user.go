package repository

import (
	"project-management/internal/model"

	"github.com/sirupsen/logrus"
)

type Repository struct {
	log *logrus.Logger
	db  interface{}
}

func NewUserRepository(l *logrus.Logger, db interface{}) *Repository {
	return &Repository{
		log: l,
		db:  db,
	}
}

func (r Repository) CreateNewUser(user model.RegisterRequest) error {
	return nil
}
