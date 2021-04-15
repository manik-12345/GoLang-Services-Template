package DataProvider

import (
	"github.com/BurntSushi/toml"
	//"github.com/Joule-CMA/DatabaseUtil"
	"github.com/rs/zerolog/log"
	"os"
)

//Member object to hold the gorm connection
type HelloServiceDatabase struct {
	HelloServiceDatabase DatabaseUtil.DBConnection
}

//Member object to hold the gorm connection
type HelloListenerDatabase struct {
	HelloListenerDatabase DatabaseUtil.DBConnection
}

func (s *HelloServiceDatabase) Configure(){
	//Set the component of zerolog for this package
	log.Info().Msg("Component value set for DataProvider logger")

	if _, tomlErr := toml.DecodeFile(os.Getenv("ConfigPath"), &s); tomlErr != nil {
		log.Error().Msgf("Error reading the configuration file %s.  Error is: %s", os.Getenv("ConfigPath"), tomlErr.Error())
		panic(s)
	}

	//Configuring DB connection pool for listener
	log.Info().Msg("Configuring DB connection pool for listeners")
	s.HelloServiceDatabase.Configure()
}



func (s *HelloListenerDatabase) Configure(){
	//Set the component of zerolog for this package
	log.Info().Msg("Component value set for DataProvider logger")

	if _, tomlErr := toml.DecodeFile(os.Getenv("ConfigPath"), &s); tomlErr != nil {
		log.Error().Msgf("Error reading the configuration file %s.  Error is: %s", os.Getenv("ConfigPath"), tomlErr.Error())
		panic(s)
	}

	//Configuring DB connection pool for listener
	log.Info().Msg("Configuring DB connection pool for listeners")
	s.HelloListenerDatabase.Configure()
}