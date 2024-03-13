package storage

import (
	"golang-api-template/internal/domain/entity"
	"golang-api-template/internal/domain/storage/dto"
	"golang-api-template/internal/domain/storage/gorm/scheme"
)

type User interface {
	InsertUser(user *dto.UserCreate) error
	Find(user *scheme.User) (*entity.User, error)
	DeleteByID(id uint) error
}
