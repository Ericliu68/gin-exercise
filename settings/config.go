package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

type ZapConfig struct {
	Level         string `mapstructure:"level"`
	Format        string `mapstructure:"format"`
	Prefix        string `mapstructure:"prefix"`
	Director      string `mapstructure:"director"`
	LinkName      string `mapstructure:"link-name"`
	ShowLine      bool   `mapstructure:"show-line"`
	EncodeLevel   string `mapstructure:"encode-level"`
	StacktraceKey string `mapstructure:"stacktrace-key"`
	LogInConsole  bool   `mapstructure:"log-in-console"`
}
type JwtConfig struct {
	SignKey string `mapstructure:"sign-key"`
}

type RedisConfig struct {
	Db0      int64  `mapstructure:"db0"`
	Addr     string `mapstructure:"addr"`
	Password string `mapstructure:"password"`
}

type MysqlConfig struct {
	Username     string `mapstructure:"username"`
	Password     string `mapstructure:"password"`
	Path         string `mapstructure:"path"`
	DbName       string `mapstructure:"db-name"`
	Config       string `mapstructure:"config"`
	MaxIdleConns int    `mapstructure:"max-idle-conns"`
	MaxOpenConns int    `mapstructure:"max-open-conns"`
	LogMode      bool   `mapstructure:"log-mode"`
}

type Config struct {
	Jwt   JwtConfig   `mapstructure:"jwt"`
	Zap   ZapConfig   `mapstructure:"zap"`
	Redis RedisConfig `mapstructure:"redis"`
	Mysql MysqlConfig `mapstructure:"mysql"`
	Salt  string      `mapstructure:"salt"`
}

var Setting *Config

// 配置文件初始化
func GetConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("config error: %s", err))
	}
	viper.Unmarshal(&Setting)
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		// 配置文件发生变更之后会调用的回调函数
		fmt.Println("Config file changed:", e.Name)
	})

}
