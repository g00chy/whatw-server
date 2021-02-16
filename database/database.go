package database

import (
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"whatw/config"
)

var d *gorm.DB
var c *viper.Viper

// Init initializes database
func Init() {
	c = config.GetConfig()

	err := openDB()
	if err != nil {
		panic(err)
	}
}

func openDB() error {
	databaseUrl := c.GetString("db.url")
	databaseType := c.GetString("db.provider")
	enableSqlLog := c.GetBool("db.enable_log")

	var err error
	gConfig := &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Silent),
		DisableForeignKeyConstraintWhenMigrating: true,
	}

	if databaseType == "sqlite3" {
		d, err = gorm.Open(sqlite.Open(databaseUrl), gConfig)
	}
	if databaseType == "postgres" {
		d, err = gorm.Open(postgres.Open(databaseUrl), gConfig)
	}
	if databaseType == "mysql" {
		d, err = gorm.Open(mysql.Open(databaseUrl), gConfig)
	}

	if enableSqlLog {
		d.Logger = d.Logger.LogMode(logger.Info)
	}

	return err
}

// GetDB returns database connection
func GetDB() *gorm.DB {
	return d
}

//
//// Close closes database
//func Close() {
//
//}
