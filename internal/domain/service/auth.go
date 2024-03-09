package service

import (
	"context"
	"golang-api-template/internal/domain/storage"
)

type Auth interface {
	Login(ctx context.Context)
}

type auth struct {
	AuthStorage storage.Auth
}

func (a *auth) Login(ctx context.Context) {

}

func (a *auth) Register(ctx context.Context) {

}
