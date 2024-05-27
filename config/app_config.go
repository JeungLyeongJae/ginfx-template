package config

import (
	"fmt"
	"ginfx-template/pkg/logger"
	"github.com/spf13/viper"
	"log/slog"
)

type AppConfig struct {
	ServerConfig ServerConfig `mapstructure:"server"`
	DbConfig     DbConfig     `mapstructure:"database"`
	RedisConfig  RedisConfig  `mapstructure:"redis"`
	JwtConfig    JwtConfig    `mapstructure:"jwt"`
}

type JwtConfig struct {
	PrivateKey string `mapstructure:"privateKey"`
	PublicKey  string `mapstructure:"publicKey"`
}

// ServerConfig Http服务器配置
type ServerConfig struct {
	Mode string `mapstructure:"mode"` // 可选 dev,prod, 默认 dev
	Port int    `mapstructure:"port"`
}

type DbConfig struct {
	Host     string `mapstructure:"host"`
	Port     int64  `mapstructure:"port"`
	Dbname   string `mapstructure:"dbname"`
	SslMode  string `mapstructure:"sslmode"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type RedisConfig struct {
	Mode       string
	MasterName string
	Addrs      []string
	Addr       string `mapstructure:"addr"`
	Password   string `mapstructure:"password"`
	DB         int    `mapstructure:"db"`
}

func (p *DbConfig) dsn() string {
	if p.Port == 0 {
		p.Port = 3306
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?&charset=utf8&parseTime=true", p.Username, p.Password, p.Host, p.Port, p.Dbname)
}

func ProvideAppConfig() *AppConfig {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	var err error
	if err = viper.ReadInConfig(); err != nil {
		logger.Fatalf("fatal error config file: %v", err)
	}
	viper.AutomaticEnv()
	slog.Info("Load config file: " + viper.ConfigFileUsed())
	config := &AppConfig{}
	if err = viper.Unmarshal(config); err != nil {
		logger.Fatal(err)
	}
	return config
}
