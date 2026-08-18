package infra

import (
	"fmt"
	"net/url"

	"github.com/amacneil/dbmate/v2/pkg/dbmate"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GetSQLStore(
	databaseHost string,
	databaseUser string,
	databasePassword string,
	databaseName string,
	databasePort string,
) *gorm.DB {
	dsn := "host=" + databaseHost +
		" user=" + databaseUser +
		" password=" + databasePassword +
		" dbname=" + databaseName +
		" port=" + databasePort +
		" sslmode=disable" +
		" TimeZone=Asia/Shanghai"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		panic(err)
	}
	return db
}

func GetSQLStoreMigration(
	databaseHost string,
	databaseUser string,
	databasePassword string,
	databaseName string,
	databasePort string,
) *dbmate.DB {
	database_url := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		databaseUser, databasePassword, databaseHost, databasePort, databaseName)

	u, _ := url.Parse(database_url)
	db := dbmate.New(u)
	return db
}
