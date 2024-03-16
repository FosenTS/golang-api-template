package storage

import (
	"golang-api-template/internal/domain/entity"
	"golang-api-template/internal/domain/storage/dto"
)

type User interface {
	InsertUser(user *dto.UserCreate) error
	FindByLogin(login string) (*entity.User, error)
	DeleteByID(id uint) error
}
