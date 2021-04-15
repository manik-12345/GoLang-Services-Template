package services

import (
	"encoding/json"
	//"github.com/Joule-CMA/Error"
	"github.com/manik-12345/GoLang-Services-Template/pkg/Model"
	"github.com/go-chi/chi"
	"net/http"
)

func (s *HelloWorldServiceConfig) GetHelloWorld(w http.ResponseWriter, r *http.Request) {
	// Reading the title from path
	s.Log.Info().Msg("GetHelloWorld: Reading the title from the path")
	title := chi.URLParam(r, "title")

	//Reading user from WSO2
	s.Log.Info().Msg("GetHelloWorld: Reading hello world from db")
	helloWorld, err := s.ReadHelloWorldByTitle(title)
	if err != nil {
		//Error reading information in database
		s.Log.Error().Err(err).Msg("GetHelloWorld: Error reading the object from the database")
		jsonError := Error.JsonError{ErrorCode: "error.unknown", Message: "Unknown error"}
		http.Error(w, jsonError.Error(), http.StatusInternalServerError)
		return
	}

	if (helloWorld == Model.HelloWorldModel{}) {
		//Entry doesn't exists
		s.Log.Error().Err(err).Msg("GetHelloWorld: Title not found in database")
		jsonError := Error.JsonError{ErrorCode: "hello.world.not.found", Message: "Title not found"}
		http.Error(w, jsonError.Error(), http.StatusNotFound)
		return
	}

	jsonObject, err := json.Marshal(helloWorld)
	if err != nil {
		//Entry doesn't exists
		s.Log.Error().Err(err).Msg("GetHelloWorld: Title not found in database")
		jsonError := Error.JsonError{ErrorCode: "hello.world.marshalling.error", Message: "Marshalling error"}
		http.Error(w, jsonError.Error(), http.StatusInternalServerError)
		return
	}

	// Write json to response
	s.Log.Info().Msg("GetHelloWorld: Writing json to response")
	w.Header().Add("Content-Type", "application/json")
	_, _ = w.Write(jsonObject)
}
