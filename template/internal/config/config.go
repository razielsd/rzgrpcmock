package config

type Config struct {
	LogLevel string `env:"LOG_LEVEL" envDefault:info`
	GRPCAddr string `env:"GRPC_ADDR" envDefault:0.0.0.0:9099`
	HTTPAddr string `env:"HTTP_ADDR" envDefault:0.0.0.0:8099`
	APIAddr  string `env:"API_ADDR" envDefault:0.0.0.0:9010`
}
