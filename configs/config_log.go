package configs

type log struct {
	console
}

type console struct {
	level  string
	isJson bool
	color  bool
}

type ILogConfig interface {
	ConsoleLevel() string
	ConsoleIsJson() bool
	ConsoleColor() bool
}

func (c *config) Log() ILogConfig {
	return c.log
}

func (l *log) ConsoleLevel() string { return l.console.level }
func (l *log) ConsoleIsJson() bool  { return l.console.isJson }
func (l *log) ConsoleColor() bool   { return l.console.color }
