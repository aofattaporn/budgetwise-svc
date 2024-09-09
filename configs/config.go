package configs

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type IConfig interface {
	App() IAppConfig
	Db() IDbConfig
	Log() ILogConfig
}

type config struct {
	app *app
	db  *db
	log *log
}

func LoadConfig(path string) (IConfig, error) {

	if path == "" {
		return nil, fmt.Errorf("config file is empty")
	}

	viper.SetConfigFile(path)
	viper.AutomaticEnv()

	// read config file
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed read config file: %v", err)
	}

	var viperConfig Config

	if err := viper.Unmarshal(&viperConfig); err != nil {
		return nil, fmt.Errorf("unable to decode into struct: %v", err)
	}

	// mapping config
	config := &config{
		app: &app{
			port:         viperConfig.Server.Port,
			contextPath:  viperConfig.Server.ContextPath,
			name:         viperConfig.Server.Name,
			version:      viperConfig.Server.Version,
			readTimeout:  time.Duration(viperConfig.Server.ReadTimeout) * time.Second,
			writeTimeout: time.Duration(viperConfig.Server.WrtieTimeout) * time.Second,
			bodyLimit:    viperConfig.Server.BodyLimit,
		},
		db: &db{
			host:            viperConfig.Database.Host,
			protocol:        viperConfig.Database.Protocol,
			username:        viperConfig.Database.User,
			password:        viperConfig.Database.Password,
			database:        viperConfig.Database.DBName,
			maxOpenConns:    viperConfig.Database.MaxOpenConns,
			maxIdleConns:    viperConfig.Database.MaxIdleConns,
			connMaxLifeTime: time.Duration(viperConfig.Database.ConnMaxLifeTime) * time.Second,
			connMaxIdleTime: time.Duration(viperConfig.Database.ConnMaxIdleTime) * time.Second,
			retryTime:       time.Duration(viperConfig.Database.RetryTimeInSec) * time.Second,
			countRetry:      viperConfig.Database.CountRetry,
		},
		log: &log{
			console: console{
				level:  viperConfig.Log.Console.Level,
				isJson: viperConfig.Log.Console.IsJson,
				color:  viperConfig.Log.Console.Color,
			},
		},
	}

	return config, nil
}
