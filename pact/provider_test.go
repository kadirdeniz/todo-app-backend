package pact

import (
	"fmt"
	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
	"github.com/pact-foundation/pact-go/utils"
	"testing"
	"todo/server"
)

type Settings struct {
	Host            string
	ProviderName    string
	BrokerBaseURL   string
	BrokerUsername  string // Basic authentication
	BrokerPassword  string // Basic authentication
	ConsumerName    string
	ConsumerVersion string // a git sha, semantic version number
	ConsumerTag     string // dev, staging, prod
	ProviderVersion string
}

func (s *Settings) create() {
	s.Host = "127.0.0.1"
	s.ProviderName = "CampaignService"
	s.ConsumerName = "TodoBackend"
	//s.BrokerBaseURL = "http://localhost"
	s.ConsumerTag = "main"
	s.ProviderVersion = "1.0.0"
	s.ConsumerVersion = "1.0.0"
}

func TestProvider(t *testing.T) {
	port, _ := utils.GetFreePort()
	go server.StartServer(port)

	settings := Settings{}
	settings.create()

	pact := dsl.Pact{
		Host:                     settings.Host,
		Provider:                 settings.ProviderName,
		Consumer:                 settings.ConsumerName,
		DisableToolValidityCheck: true,
	}

	verifyRequest := types.VerifyRequest{
		ProviderBaseURL: fmt.Sprintf("http://%s:%d", settings.Host, port),
		ProviderVersion: settings.ProviderVersion,
		//BrokerUsername:             settings.BrokerUsername,
		//BrokerURL:                  settings.BrokerBaseURL,
		//BrokerPassword:             settings.BrokerPassword,
		Tags:                       []string{settings.ConsumerTag},
		PactURLs:                   []string{"./TodoFrontend-TodoBackend.json"},
		PublishVerificationResults: false,
		FailIfNoPactsFound:         true,
	}

	verifyResponses, err := pact.VerifyProvider(t, verifyRequest)
	if err != nil {
		fmt.Println("Error on VerifyProvider: ", err)
		t.Fatal(err)
	}

	fmt.Println(len(verifyResponses), "pact tests run")
}
