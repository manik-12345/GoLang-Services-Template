package implementations

import (
	"github.com/manik-12345/GoLang-Services-Template/pkg/Model"
	"github.com/jinzhu/gorm"
)

func (s *HelloWorldConfiguration) ReadHelloWorldByTitle(helloWorldTitle string) (Model.HelloWorldModel, error) {
	var dbHelloWorld Model.HelloWorldModel
	var dbError error

	if dbError = s.DataProvider.HelloServiceDatabase.GetConnection().Where(Model.HelloWorldModel{HelloTitle: helloWorldTitle}).Find(&dbHelloWorld).Error; dbError != nil {
		// Record doesn't exist, add it to the database
		if gorm.IsRecordNotFoundError(dbError) {
			s.Log.Info().Msg("No hello world found by email")
			return Model.HelloWorldModel{}, nil
		}
		return Model.HelloWorldModel{}, dbError
	}

	return dbHelloWorld, nil
}