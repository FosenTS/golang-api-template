package storage

import (
	"golang-api-template/internal/domain/entity"
	"golang-api-template/internal/domain/storage/dto"
	"golang-api-template/internal/domain/storage/gormDB/scheme"
)

type RefreshToken interface {
	InsertRefreshToken(tCreate *dto.RefreshTokenCreate) (*entity.RefreshToken, error)
	Find(refreshToken *scheme.RefreshToken) (*entity.RefreshToken, error)
	DeleteByLogin(string) error
	DeleteByID(id uint) error
}
