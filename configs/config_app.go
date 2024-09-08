package configs

import "time"

type IAppConfig interface {
	Name() string
	ContextPath() string
	Version() string
	ReadTimeout() time.Duration
	WriteTimeout() time.Duration
	BodyLimit() int
	Port() int
}

type app struct {
	port         int
	contextPath  string
	name         string
	version      string
	readTimeout  time.Duration
	writeTimeout time.Duration
	bodyLimit    int
}

func (c *config) App() IAppConfig {
	return c.app
}

func (a *app) Name() string                { return a.name }
func (a *app) ContextPath() string         { return a.contextPath }
func (a *app) Version() string             { return a.version }
func (a *app) ReadTimeout() time.Duration  { return a.readTimeout }
func (a *app) WriteTimeout() time.Duration { return a.writeTimeout }
func (a *app) BodyLimit() int              { return a.bodyLimit }
func (a *app) Port() int                   { return a.port }