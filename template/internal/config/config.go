package config

type Config struct {
	LogLevel string `env:"LOG_LEVEL" envDefault:"info"`
	GRPCAddr string `env:"GRPC_ADDR" envDefault:":9099"`
	HTTPAddr string `env:"HTTP_ADDR" envDefault:":8099"`
	APIAddr  string `env:"HTTP_ADDR" envDefault:":9010"`
}
