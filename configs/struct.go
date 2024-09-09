package configs

type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
	Log      LogConfig      `mapstructure:"log"`
}

type AzureConfig struct {
	Vault Vault  `mapstructure:"vault"`
	Blob  Blob   `mapstructure:"blob"`
	Id    string `mapstructure:"id"`
}

type Vault struct {
	Enable bool   `mapstructure:"enable"`
	Name   string `mapstructure:"name"`
}

type Blob struct {
	Enable        bool   `mapstructure:"enable"`
	Name          string `mapstructure:"name" validate:"required"`
	ContainerName string `mapstructure:"containerName" validate:"required"`
}

type ServerConfig struct {
	Port         int    `mapstructure:"port"`
	ContextPath  string `mapstructure:"contextPath"`
	Name         string `mapstructure:"name"`
	Version      string `mapstructure:"version"`
	BodyLimit    int    `mapstructure:"bodyLimit"`
	ReadTimeout  int    `mapstructure:"readTimeoutSec"`
	WrtieTimeout int    `mapstructure:"wrtieTimeoutSec"`
}

type DatabaseConfig struct {
	Host            string `mapstructure:"host"`
	Protocol        string `mapstructure:"protocol"`
	User            string `mapstructure:"dbuser"`
	Password        string `mapstructure:"dbpwd"`
	DBName          string `mapstructure:"dbName"`
	MaxOpenConns    int    `mapstructure:"maxOpenConns"`
	MaxIdleConns    int    `mapstructure:"maxIdleConns"`
	ConnMaxLifeTime int    `mapstructure:"connMaxLifeTimeSec"`
	ConnMaxIdleTime int    `mapstructure:"connMaxIdleTimeSec"`
	RetryTimeInSec  int    `mapstructure:"retryTimeInSec"`
	CountRetry      int    `mapstructure:"countRetry"`
}

type LogConfig struct {
	Console ConsoleConfig `mapstructure:"console" validate:"required"`
}

type ConsoleConfig struct {
	Level  string `mapstructure:"level" validate:"required"`
	IsJson bool   `mapstructure:"isJson"`
	Color  bool   `mapstructure:"color"`
}
