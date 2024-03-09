package gorm

import (
	"golang-api-template/internal/domain/storage"
	"gorm.io/gorm"
)

type AuthRepository storage.Auth

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db: db}
}
