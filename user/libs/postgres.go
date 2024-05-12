package libs

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ConfigDB struct {
	PostgresHost     string
	PostgresPort     string
	PostgresName     string
	PostgresUsername string
	PostgresPassword string
	PostgresTimeout  int
}

func NewDBPostgres(config ConfigDB) (db *gorm.DB, err error) {
	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, err
}
