package services

import (
	"github.com/manik-12345/GoLang-Services-Template/pkg/API/HelloWorld/definitions"
	"github.com/rs/zerolog"
)

type HelloWorldServiceConfig struct {
	definitions.HelloWorldProviderInterface
	Log			*zerolog.Logger
}

// ImplementedBy, creates a configuration object from the definition
func ImplementedBy(implementation definitions.HelloWorldProviderInterface, componentLogger *zerolog.Logger) HelloWorldServiceConfig {
	return HelloWorldServiceConfig{
		implementation,
		componentLogger,
	}
}
