package configs

import (
	"fmt"
	"time"
)

type IDbConfig interface {
	Url() string
	MaxOpenConns() int
	MaxIdleConns() int
	ConnMaxLifeTime() time.Duration
	ConnMaxIdleTime() time.Duration
	RetryTime() time.Duration
	CountRetry() int
}

type db struct {
	host            string
	protocol        string
	username        string
	password        string
	database        string
	maxOpenConns    int
	maxIdleConns    int
	connMaxLifeTime time.Duration
	connMaxIdleTime time.Duration
	retryTime       time.Duration
	countRetry      int
}

func (c *config) Db() IDbConfig {
	return c.db
}

func (d *db) Url() string {
	// dsn := "myuser:mypassword@tcp(127.0.0.1:3306)/mydatabase?charset=utf8mb4&parseTime=True&loc=Local"
	dns := fmt.Sprintf("%s:%s@%s(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", d.username, d.password, d.protocol, d.database)
	return dns

}

func (d *db) MaxOpenConns() int              { return d.maxOpenConns }
func (d *db) MaxIdleConns() int              { return d.maxIdleConns }
func (d *db) ConnMaxLifeTime() time.Duration { return d.connMaxLifeTime }
func (d *db) ConnMaxIdleTime() time.Duration { return d.connMaxIdleTime }
func (d *db) RetryTime() time.Duration       { return d.retryTime }
func (d *db) CountRetry() int                { return d.countRetry }
