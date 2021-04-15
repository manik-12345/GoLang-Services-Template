package services

import (
	"encoding/json"
	//"github.com/Joule-CMA/Error"
	"github.com/manik-12345/GoLang-Services-Template/pkg/Model"
	"net/http"
)


// Create survey service to create a new .
func (s *HelloWorldServiceConfig) PostHelloWorld(w http.ResponseWriter, r *http.Request) {
	var helloWorld Model.HelloWorldModel

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	s.Log.Info().Msg("Decoding data from json")
	err := json.NewDecoder(r.Body).Decode(&helloWorld)
	if err != nil || helloWorld.HelloTitle == "" {
		// If Err is not nil, print it in the logs
		if err != nil {
			s.Log.Error().Err(err)
		}

		//Error marshaling data
		s.Log.Error().Msg("Unable to decode body")
		jsonError := Error.JsonError{Message: "Body does not contain a valid object", ErrorCode: "json.marshal.error"}
		http.Error(w, jsonError.Error(), http.StatusBadRequest)
		return
	}

	helloWorldCheck, _ := s.ReadHelloWorldByTitle(helloWorld.HelloTitle)
	if helloWorldCheck.HelloTitle != "" {
		// Error marshaling data
		s.Log.Error().Msg("Entry already exist for title: " + helloWorldCheck.HelloTitle )
		jsonError := Error.JsonError{Message: "Entry already exist", ErrorCode: "hello.world.exists"}
		http.Error(w, jsonError.Message, http.StatusBadRequest)
		return
	}

	// Creating survey with data passed
	s.Log.Info().Msg("Publishing message to pubsub")
	err = s.PublishHelloWorld(helloWorld)

	if err != nil {
		//Error updating information in database
		s.Log.Error().Msg("PostHelloWorld: Unable to marshal body")
		jsonError := Error.JsonError{Message: err.Error(), ErrorCode: "json.marshal.error"}
		http.Error(w, jsonError.Error(), http.StatusInternalServerError)
		return
	}
}