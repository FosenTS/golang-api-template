package application

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"golang-api-template/internal/application/config"
	"golang-api-template/pkg/advancedlog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type App interface {
}

type app struct {
	cfg  config.AppConfig
	gorm *gorm.DB
	log  *logrus.Entry
}

func createGormConnection(config *config.GormConfig, log *logrus.Entry) (*gorm.DB, error) {
	logF := advancedlog.FunctionLog(log)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%t TimeZone=%s", config.Host, config.Username, config.Password, config.DatabaseName, config.Port, config.UseCA, config.TimeZone)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction:                   config.SkipDefaultTransaction,
		FullSaveAssociations:                     config.FullSaveAssociations,
		DryRun:                                   config.DryRun,
		PrepareStmt:                              config.PrepareStmt,
		DisableAutomaticPing:                     config.DisableAutomaticPing,
		DisableForeignKeyConstraintWhenMigrating: config.DisableForeignKeyConstraintWhenMigrating,
		IgnoreRelationshipsWhenMigrating:         config.IgnoreRelationshipsWhenMigrating,
		DisableNestedTransaction:                 config.DisableNestedTransaction,
		AllowGlobalUpdate:                        config.AllowGlobalUpdate,
		QueryFields:                              config.QueryFields,
		CreateBatchSize:                          config.CreateBatchSize,
		TranslateError:                           config.TranslateError,
	})
	if err != nil {
		logF.Fatalln(err)
		return nil, err
	}

	return db, nil
}
