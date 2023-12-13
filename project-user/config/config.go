package config

import (
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"log"
	"os"
	"test.com/project-common/logs"
)

var C = InitConfig()

type Config struct {
	viper       *viper.Viper
	SC          *ServerConfig
	GC          *GrpcConfig
	MysqlConfig *MysqlConfig
}

type ServerConfig struct {
	Name string
	Adds string
}

type GrpcConfig struct {
	Adds string
	Name string
}

type MysqlConfig struct {
	Username string
	Password string
	Host     string
	Port     int
	Db       string
}

func InitConfig() *Config {
	//读取 到yaml文件的配置
	v := viper.New()
	conf := &Config{viper: v}
	workDir, _ := os.Getwd()
	conf.viper.SetConfigName("config")
	conf.viper.SetConfigType("yaml")
	conf.viper.AddConfigPath("/etc/ms_project/user")
	conf.viper.AddConfigPath(workDir + "/config")
	// 创建出一个可以根据key读取 配置文件的对象
	err := conf.viper.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	//创建结构体，然后创建函数去把配置都初始化进去， 并赋值非一个可以引用的对象
	//拿到配置可以对 结构体经行初始化， 如果不初始化，也可以直接拿conf引用变量就行
	conf.ReadServerConfig()
	conf.InitZapLog()
	conf.InitRedisOptions()
	conf.ReadGrpcConfigConfig()
	conf.InitMysqlConfig()
	return conf
}

func (c *Config) InitZapLog() {
	//从配置中读取日志配置，初始化日志
	lc := &logs.LogConfig{
		DebugFileName: c.viper.GetString("zap.debugFileName"),
		InfoFileName:  c.viper.GetString("zap.infoFileName"),
		WarnFileName:  c.viper.GetString("zap.warnFileName"),
		MaxSize:       c.viper.GetInt("zap.maxSize"),
		MaxAge:        c.viper.GetInt("zap.maxAge"),
		MaxBackups:    c.viper.GetInt("zap.maxBackups"),
	}
	err := logs.InitLogger(lc)
	if err != nil {
		log.Fatalln(err)
	}
}

func (c *Config) ReadGrpcConfigConfig() {
	c.GC = &GrpcConfig{
		Name: c.viper.GetString("grpc.name"),
		//你写erro word了，根据env会下滑行提示你 显示 曲线下划线
		Adds: c.viper.GetString("grpc.adds"),
	}
}

func (c *Config) ReadServerConfig() {
	c.SC = &ServerConfig{
		Name: c.viper.GetString("server.name"),
		Adds: c.viper.GetString("server.adds"),
	}
}

func (c *Config) InitRedisOptions() *redis.Options {
	return &redis.Options{
		Addr:     c.viper.GetString("redis.host") + ":" + c.viper.GetString("redis.port"),
		Password: c.viper.GetString("redis.password"), // no password set
		DB:       c.viper.GetInt("db"),                // use default DB
	}
}

func (c *Config) InitMysqlConfig() {
	mc := &MysqlConfig{
		Username: c.viper.GetString("mysql.username"),
		Password: c.viper.GetString("mysql.password"),
		Host:     c.viper.GetString("mysql.host"),
		Port:     c.viper.GetInt("mysql.port"),
		Db:       c.viper.GetString("mysql.db"),
	}
	c.MysqlConfig = mc
}
