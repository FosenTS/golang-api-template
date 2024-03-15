package application

import (
	"context"
	"fmt"
	"golang-api-template/internal/application/config"
	"golang-api-template/internal/application/product"
	"golang-api-template/pkg/advancedlog"

	"github.com/alitto/pond"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cache"
	"github.com/weaveworks/promrus"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type App interface {
	Run(ctx context.Context) error
}

type app struct {
	appCfg   config.AppConfig
	httpCfg  config.HTTPConfig
	endpoint product.Endpoint
	db       *gorm.DB
	log      *logrus.Entry
}

func NewApp(ctx context.Context) (App, error) {
	appConfig := config.App()

	log, err := createLogrus(&appConfig)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	db, err := createGormConnection(config.Gorm(), log.WithField("location", "gorm-connection"))
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	storage, err := product.NewStorage(db, log.WithField("location", "storage"))
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	services := product.NewServices(storage, config.Auth(), log.WithField("location", "services"))

	gateway := product.NewGateway(services)

	controller := product.NewController(gateway, config.Metrics(), log.WithField("location", "controller"))

	endpoint := product.NewEndpoint(controller)

	return &app{
		appCfg:   appConfig,
		httpCfg:  config.HTTP(),
		endpoint: *endpoint,
		db:       db,
		log:      log.WithField("location", "applicaiton"),
	}, nil
}

func (app *app) Run(ctx context.Context) error {
	pool := pond.New(5, 5)
	defer pool.StopAndWait()

	group, grpCtx := pool.GroupContext(ctx)
	group.Submit(func() error {
		err := app.runHTTP(grpCtx)
		if err != nil {
			app.log.Fatalln(err)
			return err
		}
		return nil
	})
	group.Submit(func() error {
		err := app.endpoint.ListenMetrics()
		if err != nil {
			app.log.Fatalln(err)
			return err
		}
		return nil
	})

	err := group.Wait()
	if err != nil {
		app.log.Fatalln(err)
		return err
	}

	return nil
}

func (app *app) runMetricsListen(ctx context.Context) error {
	err := app.endpoint.ListenMetrics()
	if err != nil {
		return err
	}
	return nil
}

func (app *app) runHTTP(ctx context.Context) error {
	fApp := fiber.New(fiber.Config{
		Concurrency: int(app.httpCfg.MaxConcurrentConnection),
	})

	if app.httpCfg.UseCache {
		fApp.Use(cache.New())
	}

	app.endpoint.ConfigureFiber(fApp)

	addr := fmt.Sprintf("%s:%s", app.httpCfg.Host, app.httpCfg.Port)
	fApp.Listen(addr)

	return nil
}

func createLogrus(config *config.AppConfig) (*logrus.Logger, error) {
	logger := logrus.New()

	level := logger.GetLevel()
	switch config.LogLevel {
	case logrus.InfoLevel.String():
		level = logrus.InfoLevel
	case logrus.DebugLevel.String():
		level = logrus.DebugLevel
	case logrus.ErrorLevel.String():
		level = logrus.ErrorLevel
	case logrus.FatalLevel.String():
		level = logrus.FatalLevel
	}

	logger.SetLevel(level)

	metricsHook, err := promrus.NewPrometheusHook()
	if err != nil {
		return nil, err
	}

	logger.AddHook(metricsHook)

	return logger, nil

}

func createGormConnection(config *config.GormConfig, log *logrus.Entry) (*gorm.DB, error) {
	logF := advancedlog.FunctionLog(log)
	sslmode := "disable"
	if config.UseCA {
		sslmode = "enable"
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", config.Host, config.Username, config.Password, config.DatabaseName, config.Port, sslmode, config.TimeZone)
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
