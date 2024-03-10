package advancedlog

import (
	"github.com/sirupsen/logrus"
	"runtime"
)

const (
	FuncField     = "function"
	LocationField = "location"
)

func FunctionLog(log *logrus.Entry) *logrus.Entry {
	pc, _, _, _ := runtime.Caller(1)
	funcName := runtime.FuncForPC(pc).Name()

	return log.WithField(FuncField, funcName)
}
