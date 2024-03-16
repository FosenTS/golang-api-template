package storage

import (
	"golang-api-template/internal/domain/entity"
	"golang-api-template/internal/domain/storage/dto"
)

type RefreshToken interface {
	InsertRefreshToken(tCreate *dto.RefreshTokenCreate) (*entity.RefreshToken, error)
	FindByToken(token string) (*entity.RefreshToken, error)
	DeleteByLogin(string) error
	DeleteByID(id uint) error
}
