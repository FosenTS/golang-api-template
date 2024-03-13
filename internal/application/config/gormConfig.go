package config

import (
	"sync"
)

type GormConfig struct {
	Host         string
	Port         string
	Username     string
	Password     string
	DatabaseName string
	UseCA        bool
	CaPath       string
	TimeZone     string

	// GORM perform single create, update, delete operations in transactions by default to ensure database data integrity
	// You can disable it by setting `SkipDefaultTransaction` to true
	SkipDefaultTransaction bool

	// FullSaveAssociations full save associations
	FullSaveAssociations bool

	// DryRun generate sql without execute
	DryRun bool
	// PrepareStmt executes the given query in cached statement
	PrepareStmt bool
	// DisableAutomaticPing
	DisableAutomaticPing bool
	// DisableForeignKeyConstraintWhenMigrating
	DisableForeignKeyConstraintWhenMigrating bool
	// IgnoreRelationshipsWhenMigrating
	IgnoreRelationshipsWhenMigrating bool
	// DisableNestedTransaction disable nested transaction
	DisableNestedTransaction bool
	// AllowGlobalUpdate allow global update
	AllowGlobalUpdate bool
	// QueryFields executes the SQL query with all fields of the table
	QueryFields bool
	// CreateBatchSize default create batch size
	CreateBatchSize int
	// TranslateError enabling error translation
	TranslateError bool
}

var (
	gormConfigInst     = &GormConfig{}
	loadGormConfigOnce = sync.Once{}
)

func Gorm() *GormConfig {
	loadGormConfigOnce.Do(func() {
		env := Env()
		gormConfigInst.Host = env.PostgresHost
		gormConfigInst.Port = env.PostgresPort
		gormConfigInst.Username = env.PostgresUsername
		gormConfigInst.Password = env.PostgresPassword
		gormConfigInst.DatabaseName = env.PostgresDatabaseName
		gormConfigInst.UseCA = env.PostgresUseCA
		gormConfigInst.CaPath = env.PostgresCaPath
		gormConfigInst.TimeZone = env.PostgresTimeZone
	})

	return gormConfigInst
}
