package apihandler

import (
	"golang-api-template/internal/infrastructure/controllers/fiberHTTP"

	"github.com/sirupsen/logrus"
)

type handlerApi struct {
	log *logrus.Entry
}

func NewHandlerApi(log *logrus.Entry) fiberHTTP.HandlerFiber {
	return &handlerApi{
		log: log,
	}
}
