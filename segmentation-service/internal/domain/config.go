package domain

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
	DSN string `yaml:"dsn" env:"DSN"`
}
