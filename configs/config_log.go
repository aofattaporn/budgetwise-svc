package configs

type log struct {
	console
	file
}

type console struct {
	level  string
	isJson bool
	color  bool
}

type file struct {
	enable bool
	name   string
	level  string
	isJson bool
}

type ILogConfig interface {
	ConsoleLevel() string
	ConsoleIsJson() bool
	ConsoleColor() bool
	FileEnable() bool
	FileName() string
	FileLevel() string
	FileIsJson() bool
}

func (c *config) Log() ILogConfig {
	return c.log
}

func (l *log) ConsoleLevel() string { return l.console.level }
func (l *log) ConsoleIsJson() bool  { return l.console.isJson }
func (l *log) ConsoleColor() bool   { return l.console.color }
func (l *log) FileEnable() bool     { return l.file.enable }
func (l *log) FileName() string     { return l.file.name }
func (l *log) FileLevel() string    { return l.file.level }
func (l *log) FileIsJson() bool     { return l.file.isJson }
