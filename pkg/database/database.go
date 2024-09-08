package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/goproject/configs"
	"github.com/goproject/pkg/log"
)

type IDatabase interface {
	GetDb() *sql.DB
	Close() error
}

type mysqlDatabase struct {
	Db *sql.DB
}

func NewMysqlDatabase(cfg configs.IDbConfig, logger log.ILogger) (*sql.DB, error) {

	db, err := sql.Open("mysql", cfg.Url())
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(cfg.MaxIdleConns())
	db.SetConnMaxIdleTime(cfg.ConnMaxIdleTime())

	db.SetMaxOpenConns(cfg.MaxOpenConns())
	db.SetConnMaxLifetime(cfg.ConnMaxLifeTime())

	return db, nil
}

func (m *mysqlDatabase) GetDb() *sql.DB {
	return m.Db
}

func (m *mysqlDatabase) Close() error {
	return m.Db.Close()
}
