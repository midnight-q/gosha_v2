package core

import (
	"gorm.io/gorm"
	"skeleton-app/settings"
	"gorm.io/driver/postgres"
    "gorm.io/gorm/logger"
)

const DbConnectString =
    "host='" +          settings.DbHost +
    "' port='" +        settings.DbPort +
    "' user='" +	    settings.DbUser +
    "' password='" +    settings.DbPass +
    "' dbname='" +    	settings.DbName +
    "' sslmode='disable'"


var	config = gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}

var Db, DbErr = gorm.Open(postgres.Open(DbConnectString), &config)

func EnableSqlLog() {
	Db.Logger.LogMode(logger.Info)
}

func DisableSqlLog() {
	Db.Logger.LogMode(logger.Error)
}

func GetTableName(model interface{}) string {
	stmt := &gorm.Statement{DB: Db}
	_ = stmt.Parse(model)
	return stmt.Schema.Table
}
