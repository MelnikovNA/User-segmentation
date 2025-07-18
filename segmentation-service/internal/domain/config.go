package domain

import "time"

type Config struct {
	AppName    string     `yaml:"appname" env-default:"segmentation-service"`
	GrpcServer GrpcServer `yaml:"grpcserver" env-prefix:"SEGMENTATION_SERVICE_"`
	Sqlite     Sqlite     `yaml:"sqlite" env-prefix:"SEGMENTATION_SERVICE_"`
}

type GrpcServer struct {
	Host string `yaml:"host" env-default:"0.0.0.0" env:"HOST"`
	Port string `yaml:"port" env-default:"9001" env:"PORT"`
}

type Sqlite struct {
	DSN             string        `yaml:"dsn" env:"MYSQL_DSN"`
	ConnMaxLifetime time.Duration `yaml:"conn_max_lifetime" env-default:"5m"`
	MaxOpenConns    int           `yaml:"max_open_conns" env-default:"10"`
	MaxIdleConns    int           `yaml:"max_idle_conns" env-default:"10"`
}
