package domain

type Config struct {
	App    App          `yaml:"app" env-prefix:"SEGMENTATION_" json:"app"`
	Listen ListenConfig `yaml:"listen" env-prefix:"SEGMENTATION_LISTEN_" json:"listen"`
	Log    Logger       `yaml:"log" env-prefix:"SEGMENTATION_" json:"log"`
	Grpc   Grpc         `yaml:"grpc" env-prefix:"SEGMENTATION_" json:"grpc"`
}

type App struct {
	Cors bool `yaml:"cors" env:"CORS"`
}

type ListenConfig struct {
	Ports portConfig `yaml:"ports" env-prefix:"PORTS_"`
	Host  string     `yaml:"host" env:"HOSTNAME" env-default:"0.0.0.0"`
}

type portConfig struct {
	Http string `yaml:"http" env:"HTTP"`
}

type Logger struct {
	Level map[string]string `yaml:"level" env:"LOG_LEVEL"`
}

type Grpc struct {
	Clients Clients `yaml:"clients" env-prefix:"GRPC_CLIENTS_"`
}

type Clients struct {
	SegmentService string `yaml:"segmentservice" env:"SEGMENTSERVICE"`
}
