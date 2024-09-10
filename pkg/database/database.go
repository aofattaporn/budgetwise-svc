package database

import (
	"github.com/goproject/configs"
	"github.com/goproject/pkg/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type IDatabase interface {
	GetDb() *gorm.DB
	Close() error
}

type mysqlDatabase struct {
	Db *gorm.DB
}

func NewMysqlDatabase(cfg configs.IDbConfig, logger log.ILogger) (IDatabase, error) {

	mysqlDb, err := gorm.Open(mysql.Open(cfg.Url()), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &mysqlDatabase{
		Db: mysqlDb,
	}, nil
}

func (m *mysqlDatabase) GetDb() *gorm.DB {
	return m.Db
}

func (m *mysqlDatabase) Close() error {
	sql, err := m.Db.DB()
	if err != nil {
		return err
	}
	return sql.Close()
}
