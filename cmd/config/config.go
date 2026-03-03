package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/go-playground/validator/v10"
)

type Configuracion struct {
	AppEnv          string `env:"ENVIRONMENT" envDefault:"dev" validate:"oneof=dev qa production"`
	Port            int    `env:"PORT" envDefault:"3200"`
	LogLevel        string `env:"LOG_LEVEL" envDefault:"info"`
	ShutDownTimeout int    `env:"SHUTDOWN_TIMEOUT" envDefault:"30"`
	// Se definen las variables de entorno necesarias para el proyecto
}

func CargarVariables() (*Configuracion, error) {
	cfg := &Configuracion{}

	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	v := validator.New()

	if err := v.Struct(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
