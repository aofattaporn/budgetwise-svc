package configs

type Config struct {
	Azure    AzureConfig    `mapstructure:"azure"`
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
	Log      LogConfig      `mapstructure:"log"`
	Customer CustomerConfig `mapstructure:"customer"`
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
	Port         int    `mapstructure:"port" validate:"required"`
	ContextPath  string `mapstructure:"contextPath"`
	Name         string `mapstructure:"name" validate:"required"`
	Version      string `mapstructure:"version" validate:"required"`
	BodyLimit    int    `mapstructure:"bodyLimit"`
	ReadTimeout  int    `mapstructure:"readTimeoutSec"`
	WrtieTimeout int    `mapstructure:"wrtieTimeoutSec"`
}

type DatabaseConfig struct {
	Host            string `mapstructure:"host" validate:"required"`
	Protocol        string `mapstructure:"protocol" validate:"required"`
	User            string `mapstructure:"dbuser" validate:"required"`
	Password        string `mapstructure:"dbpwd"`
	DBName          string `mapstructure:"dbName" validate:"required"`
	MaxOpenConns    int    `mapstructure:"maxOpenConns"`
	MaxIdleConns    int    `mapstructure:"maxIdleConns"`
	ConnMaxLifeTime int    `mapstructure:"connMaxLifeTimeSec"`
	ConnMaxIdleTime int    `mapstructure:"connMaxIdleTimeSec"`
	RetryTimeInSec  int    `mapstructure:"retryTimeInSec"`
	CountRetry      int    `mapstructure:"countRetry"`
}

type LogConfig struct {
	Console ConsoleConfig `mapstructure:"console" validate:"required"`
	File    FileConfig    `mapstructure:"file" validate:"required"`
}

type ConsoleConfig struct {
	Level  string `mapstructure:"level" validate:"required"`
	IsJson bool   `mapstructure:"isJson"`
	Color  bool   `mapstructure:"color"`
}

type FileConfig struct {
	Enable bool   `mapstructure:"enable"`
	Name   string `mapstructure:"name"`
	Level  string `mapstructure:"level"`
	IsJson bool   `mapstructure:"isJson"`
}

type CustomerConfig struct {
	BatchSleepTimeMinute int    `mapstructure:"batchSleepTimeMinute" validate:"required"`
	SplitInsertNum       int    `mapstructure:"splitInsertNum" validate:"required"`
	LimitConcurrent      int    `mapstructure:"limitConcurrent" validate:"required"`
	FolderCcomFile       string `mapstructure:"folderCcomFile" validate:"required"`
	LimitDaysRollback    int    `mapstructure:"limitDaysRollback" validate:"required"`
	CronTabCiBatch       string `mapstructure:"cronTabCiBatch" validate:"required"`
}
