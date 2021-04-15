package HelloWorld

import (
	"github.com/manik-12345/GoLang-Services-Template/pkg/API/HelloWorld/implementations"
	"github.com/manik-12345/GoLang-Services-Template/pkg/API/HelloWorld/services"
	"github.com/manik-12345/GoLang-Services-Template/pkg/DataProvider"
	"github.com/go-chi/chi"
	"github.com/rs/zerolog"
)

type ServiceTemplateComponent struct {
	DataProvider *DataProvider.HelloServiceDatabase
}

// Router is the Implementation to route CMAH member API calls
func (s *ServiceTemplateComponent) Router(router chi.Router) {
	// Set time format of logs
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	// Get a new implementation of the CMAH member Configuration
	implementation := implementations.Configure(s.DataProvider)
	implementation.Log.Info().Msg("Add DataProvider to the implementation")

	// Add CORS middleware around every request
	//implementation.Log.Info().Msg("Adding CORS to the CMAH member router")
	//router.Use(s.CorsOptions.Cors.Handler)

	// Tie the implementation with the services (each services is business use case)
	implementation.Log.Info().Msg("Configuration added to the implementation")
	serviceConfiguration := services.ImplementedBy(implementation, &implementation.Log)

	// the router answers to a get event with a Reaction on the path proxy
	// by default 404 are redirected to https://google.com
	router.NotFound(serviceConfiguration.NotFound)

	//Add route for the Add CMAH Member information
	implementation.Log.Info().Msg("Adding /hello/ POST route to component")
	router.Post("/", serviceConfiguration.PostHelloWorld)

	//Add route for the Update CMAH Member information
	implementation.Log.Info().Msg("Adding /hello/{title} GET route to component")
	router.Get("/{title}", serviceConfiguration.GetHelloWorld)
}
