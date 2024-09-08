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
	return fmt.Sprintf(
		"%s:%s@%s(%s)/%s?tls=true&parseTime=true&charset=utf8mb4&loc=Local",
		d.username,
		d.password,
		d.protocol,
		d.host,
		d.database,
	)
}

func (d *db) MaxOpenConns() int              { return d.maxOpenConns }
func (d *db) MaxIdleConns() int              { return d.maxIdleConns }
func (d *db) ConnMaxLifeTime() time.Duration { return d.connMaxLifeTime }
func (d *db) ConnMaxIdleTime() time.Duration { return d.connMaxIdleTime }
func (d *db) RetryTime() time.Duration       { return d.retryTime }
func (d *db) CountRetry() int                { return d.countRetry }
