package definitions

import (
	"github.com/manik-12345/GoLang-Services-Template/pkg/Model"
)

// this is the contract that allows us to bind the business logic (use cases) in the component package
// to the implementation in the implementationOf package
type HelloWorldProviderInterface interface {

	// *********************************************************
	// found in implementations/ReadHelloWorldByTitle.go
	// *********************************************************
	ReadHelloWorldByTitle(helloWorldTitle string) (Model.HelloWorldModel, error)

	// *********************************************************
	// found in implementations/PublishHelloWorld.go
	// *********************************************************
	PublishHelloWorld(helloWorld Model.HelloWorldModel) error

}

