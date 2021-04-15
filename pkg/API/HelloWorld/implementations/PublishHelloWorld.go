package implementations

import (
	"encoding/json"
	"errors"
	"github.com/manik-12345/GoLang-Services-Template/pkg/Model"
)

func (s *HelloWorldConfiguration) PublishHelloWorld(helloWorld Model.HelloWorldModel) error {
	// Marshaling result to json
	s.Log.Info().Msg("PublishHelloWorld: Marshaling hello world data to json")
	jsonObject, err := json.Marshal(helloWorld)

	// Check for marshalling errors and log them
	if err != nil {
		s.Log.Error().Err(err).Msg("PublishHelloWorld: Marshaling to json error received")
		return errors.New("unable to marshal object")
	}

	//Publishing message to pub/sub
	s.Log.Info().Msg("PublishHelloWorld: publishing data information to pub/sub")
	err = s.HelloWorldPublisher.Publish(string(jsonObject))

	if err != nil {
		// Error sending the message to pub/sub
		s.Log.Error().Err(err).Msg("Er. " + string(jsonObject))
		return err
	}

	//Message published to pub/sub
	s.Log.Info().Msg("PublishHelloWorld: Message published to pub/sub")
	return nil
}