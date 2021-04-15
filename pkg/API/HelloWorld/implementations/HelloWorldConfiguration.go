package implementations

import (
	"github.com/BurntSushi/toml"
	"github.com/manik-12345/GoLang-Services-Template/pkg/DataProvider"
	//PubSub "github.com/Joule-CMA/PubSub"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

//Configuration items
type HelloWorldConfiguration struct {
	DataProvider              		*DataProvider.HelloServiceDatabase
	Log                       		zerolog.Logger
	HelloWorldPublisher 	  		PubSub.Publisher
}


// Configuration for the UuidApi implementation.  It will create the database configuration
func Configure(helloWorldDataProvider *DataProvider.HelloServiceDatabase) *HelloWorldConfiguration {
	// Create configuration object
	var configuration HelloWorldConfiguration

	// Set the output of the logger of zerolog
	configuration.Log = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	//Set the component of zerolog for this package
	configuration.Log = log.Logger.With().Str("component", "HelloWorldConfiguration").Logger()
	configuration.Log.Info().Msg("Component value set for HelloWorld router logger")

	// Read Toml File
	if _, tomlErr := toml.DecodeFile(os.Getenv("ConfigPath"), &configuration); tomlErr != nil {
		configuration.Log.Error().Msgf("Error reading the configuration file %s.  Error is: %s", os.Getenv("ConfigPath"), tomlErr.Error())
		panic(configuration)
	}

	// Set dataProvider
	configuration.Log.Info().Msg("CmahMemberConfiguration: setting the data provider")
	configuration.DataProvider = helloWorldDataProvider

	//Creating the publisher for the pubsub
	configuration.Log.Info().Msg("CreditConfiguration: Initializing POST publisher")
	configuration.HelloWorldPublisher = PubSub.CreatePublisher("HELLO.WORLD")

	// Return configuration object created
	configuration.Log.Info().Msg("Implementation configuration object created")

	return &configuration
}

