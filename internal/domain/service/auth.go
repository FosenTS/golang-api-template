package service

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"golang-api-template/internal/domain/entity"
	"golang-api-template/internal/domain/storage"
	"golang-api-template/internal/domain/storage/dto"
	"golang-api-template/internal/domain/storage/gorm/scheme"
	"golang-api-template/internal/infrastructure/controllers/safeobject"
	"golang-api-template/pkg/advancedlog"
	"golang-api-template/pkg/ajwt"
	"golang-api-template/pkg/passlib"
)

type Auth interface {
	Register(ctx context.Context, register *dto.UserCreate) error
	Login(ctx context.Context, login *dto.Login) (*safeobject.User, *safeobject.PairToken, error)
}

var ErrNotFound = errors.New("not found user")

type auth struct {
	userStorage         storage.User
	refreshTokenStorage storage.RefreshToken

	hashManager passlib.HashManager
	jwtManager  ajwt.JWTManager

	log *logrus.Entry
}

func (a *auth) createUserRefreshToken(ctx context.Context, login string) (*entity.RefreshToken, error) {
	logF := advancedlog.FunctionLog(a.log)
	token, createTime, expireTime, err := a.jwtManager.NewRefreshToken()
	if err != nil {
		logF.Errorln(err)
		return nil, err
	}
	err = a.refreshTokenStorage.DeleteByLogin(login)
	if err != nil {
		logF.Errorln(err)
		return nil, err
	}

	rTC := dto.NewRefreshTokenCreate(token, login, createTime, expireTime)
	rT, err := a.refreshTokenStorage.InsertRefreshToken(rTC)
	if err != nil {
		return nil, err
	}

	return rT, nil
}

func (a *auth) Login(ctx context.Context, login *dto.Login) (*safeobject.User, *safeobject.PairToken, error) {
	logF := advancedlog.FunctionLog(a.log)

	user, err := a.userStorage.Find(&scheme.User{
		Login: login.Login,
	})

	err = a.hashManager.Compare(user.Password, login.Password)
	if err != nil {
		logF.Errorln(err)
		return nil, nil, err
	}

	accessT, err := a.jwtManager.NewUser(user.Login)
	if err != nil {
		logF.Errorln(err)
		return nil, nil, err
	}

	refreshT, err := a.createUserRefreshToken(ctx, user.Login)
	if err != nil {
		logF.Errorln(err)
		return nil, nil, err
	}

	if err != nil {
		logF.Errorln(err)
		return nil, nil, err
	}

	pair := safeobject.NewPairToken(accessT, refreshT)

	userSafe := safeobject.NewUser(user.Login, user.Permission)

	return userSafe, pair, nil
}

func (a *auth) Refresh(ctx context.Context, token string) (*safeobject.PairToken, error) {
	_, err := a.jwtManager.ParseRefreshToken(ctx, token)
	if err != nil {
		return nil, err
	}
	rt, err := a.refreshTokenStorage.Find(&scheme.RefreshToken{Token: token})
	if err != nil {
		return nil, err
	}
	if rt == nil {
		return nil, ErrNotFound
	}

	err = a.refreshTokenStorage.DeleteByID(rt.ID)
	if err != nil {
		return nil, err
	}

	user, err := a.userStorage.Find(&scheme.User{Login: rt.Login})

	if user == nil {
		return nil, ErrNotFound
	}

	accessToken, err := a.jwtManager.NewUser(rt.Login)
	if err != nil {
		return nil, err
	}
	refreshToken, err := a.createUserRefreshToken(ctx, rt.Login)
	if err != nil {
		return nil, err
	}

	return safeobject.NewPairToken(accessToken, refreshToken), nil
}

func (a *auth) Register(ctx context.Context, register *dto.UserCreate) error {
	logF := advancedlog.FunctionLog(a.log)
	err := a.userStorage.InsertUser(register)
	if err != nil {
		logF.Errorln(err)
		return err
	}

	return nil
}

func (a *auth) Check(ctx context.Context) {

}
