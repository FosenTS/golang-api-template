package service

import (
	"context"
	"errors"
	"golang-api-template/internal/domain/entity"
	"golang-api-template/internal/domain/storage"
	"golang-api-template/internal/domain/storage/dto"
	"golang-api-template/internal/infrastructure/controllers/safeobject"
	"golang-api-template/pkg/advancedlog"
	"golang-api-template/pkg/ajwt"
	"golang-api-template/pkg/passlib"
	"golang-api-template/pkg/qrgen"

	"github.com/sirupsen/logrus"
)

type Auth interface {
	Policy(ctx context.Context, token string) (*safeobject.User, error)

	Register(ctx context.Context, register *dto.UserCreate) error
	Login(ctx context.Context, login *dto.Login) (*safeobject.PairToken, error)
	Check(ctx context.Context, token string) (*safeobject.User, error)
	Refresh(ctx context.Context, token string) (*safeobject.PairToken, error)
}

var ErrNotFound = errors.New("not found user")

type auth struct {
	userStorage         storage.User
	refreshTokenStorage storage.RefreshToken

	hashManager passlib.HashManager
	jwtManager  ajwt.JWTManager

	log *logrus.Entry
}

func NewAuth(userStorage storage.User, refreshTokenStorage storage.RefreshToken, hashManaher passlib.HashManager, jwtManager ajwt.JWTManager, log *logrus.Entry) Auth {

	return &auth{
		userStorage:         userStorage,
		refreshTokenStorage: refreshTokenStorage,
		hashManager:         hashManaher,
		jwtManager:          jwtManager,
		log:                 log,
	}
}

func (a *auth) CreateQRAuth() ([]byte, error) {
	qr, err := qrgen.Encode("", qrgen.Medium, 256)
	if err != nil {
		return nil, err
	}

	return qr, nil
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
		logF.Errorln(err)
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return rT, nil
}

func (a *auth) Login(ctx context.Context, login *dto.Login) (*safeobject.PairToken, error) {
	logF := advancedlog.FunctionLog(a.log)

	user, err := a.userStorage.FindByLogin(login.Login)
	if err != nil {
		logF.Errorln(err)
		return nil, err
	}

	err = a.hashManager.Compare(user.Password, login.Password)
	if err != nil {
		logF.Errorln(err)
		return nil, err
	}

	accessT, err := a.jwtManager.NewUser(user.Login)
	if err != nil {
		logF.Errorln(err)
		return nil, err
	}

	refreshT, err := a.createUserRefreshToken(ctx, user.Login)
	if err != nil {
		logF.Errorln(err)
		return nil, err
	}

	if err != nil {
		logF.Errorln(err)
		return nil, err
	}

	pair := safeobject.NewPairToken(accessT, refreshT.Token)

	return pair, nil
}

func (a *auth) Check(ctx context.Context, token string) (*safeobject.User, error) {
	logF := advancedlog.FunctionLog(a.log)

	userClaims, err := a.jwtManager.ParseUser(token)
	if err != nil {
		logF.Errorln(err)
		return nil, err
	}

	user, err := a.userStorage.FindByLogin(userClaims.Login)
	if err != nil {
		logF.Errorln(err)
		return nil, err
	}

	return &safeobject.User{
		Login:      user.Login,
		Permission: user.Permission,
	}, nil
}

func (a *auth) Policy(ctx context.Context, token string) (*safeobject.User, error) {
	logF := advancedlog.FunctionLog(a.log)
	userClaims, err := a.jwtManager.ParseUser(token)
	if err != nil {
		logF.Errorln(err)
		return nil, err
	}

	user, err := a.userStorage.FindByLogin(userClaims.Login)
	if err != nil {
		logF.Errorln(err)
		return nil, err
	}

	return &safeobject.User{
		Login:      user.Login,
		Permission: user.Permission,
	}, nil
}

func (a *auth) Refresh(ctx context.Context, token string) (*safeobject.PairToken, error) {
	logF := advancedlog.FunctionLog(a.log)

	_, err := a.jwtManager.ParseRefreshToken(ctx, token)
	if err != nil {
		logF.Errorln(err)
		return nil, err
	}
	rt, err := a.refreshTokenStorage.FindByToken(token)
	if err != nil {
		logF.Errorln(err)
		return nil, err
	}
	if rt == nil {
		logF.Errorln(ErrNotFound)
		return nil, ErrNotFound
	}

	err = a.refreshTokenStorage.DeleteByID(rt.ID)
	if err != nil {
		logF.Errorln(err)
		return nil, err
	}

	user, err := a.userStorage.FindByLogin(rt.Login)
	if user == nil {
		logF.Errorln(ErrNotFound)
		return nil, ErrNotFound
	}

	accessToken, err := a.jwtManager.NewUser(rt.Login)
	if err != nil {
		logF.Errorln(err)
		return nil, err
	}
	refreshToken, err := a.createUserRefreshToken(ctx, rt.Login)
	if err != nil {
		logF.Errorln(err)
		return nil, err
	}

	return safeobject.NewPairToken(accessToken, refreshToken.Token), nil
}

func (a *auth) Register(ctx context.Context, register *dto.UserCreate) error {
	logF := advancedlog.FunctionLog(a.log)

	pass, err := a.hashManager.Hash(register.Password)
	if err != nil {
		logF.Errorln(err)
		return err
	}

	err = a.userStorage.InsertUser(dto.NewUserCreate(register.Login, pass, register.Permission))
	if err != nil {
		logF.Errorln(err)
		return err
	}

	return nil
}
