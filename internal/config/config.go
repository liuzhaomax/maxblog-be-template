package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync"
)

var cfg *Config
var once sync.Once

func init() {
	once.Do(func() {
		cfg = &Config{}
	})
}

func GetInstanceOfConfig() *Config {
	return cfg
}

type Config struct {
	mux     sync.Mutex
	RunMode string `mapstructure:"run_mode"`
	App     App    `mapstructure:"app"`
	Server  Server `mapstructure:"server"`
	DB      DB     `mapstructure:"db"`
	Mysql   Mysql  `mapstructure:"mysql"`
}

type App struct {
	AppName string `mapstructure:"app_name"`
	Version string `mapstructure:"version"`
}

type Server struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type DB struct {
	Type         string `mapstructure:"type"`
	Debug        bool   `mapstructure:"debug"`
	DSN          string `mapstructure:"dsn"`
	MaxLifetime  int    `mapstructure:"max_life_time"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

type Mysql struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	UserName string `mapstructure:"user_name"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"db_name"`
	Params   string `mapstructure:"params"`
}

func (mysql *Mysql) DSN() string {
	if mysql.Password == "" {
		mysql.Password = "123456"
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s",
		mysql.UserName, mysql.Password, mysql.Host, mysql.Port, mysql.DBName, mysql.Params)
}

func (cfg *Config) Load(configDir string, configName string) error {
	dataJson, err := ioutil.ReadFile(configDir + "/" + configName)
	if err != nil {
		return err
	}
	err = json.Unmarshal(dataJson, cfg)
	if err != nil {
		return err
	}
	return nil
}
