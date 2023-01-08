package pact

import (
	"fmt"
	"os"
	"testing"
	"todo/server"

	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
	"github.com/pact-foundation/pact-go/utils"
)

type Settings struct {
	Host            string
	ProviderName    string
	BrokerBaseURL   string
	BrokerUsername  string // Basic authentication
	BrokerToken     string // Basic authentication
	ConsumerName    string
	ConsumerVersion string // a git sha, semantic version number
	ConsumerTag     string // dev, staging, prod
	ProviderVersion string
}

func (s *Settings) create() {
	s.Host = "127.0.0.1"
	s.ProviderName = "TodoBackend"
	s.ConsumerName = "TodoFrontend"
	s.BrokerBaseURL = os.Getenv("PACT_FLOW_BASE_URL")
	s.ConsumerTag = "main"
	s.ProviderVersion = "1.0.0"
	s.ConsumerVersion = "1.0.0"
	s.BrokerUsername = os.Getenv("PACT_FLOW_USERNAME")
	s.BrokerToken = os.Getenv("PACT_FLOW_API_KEY")
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
		//Tags:                       []string{settings.ConsumerTag},
		PactURLs:                   []string{"https://kadirdenz.pactflow.io/pacts/provider/TodoBackend/consumer/TodoFrontend/version/1.0.0"},
		PublishVerificationResults: true,
		FailIfNoPactsFound:         true,
		BrokerURL:                  settings.BrokerBaseURL,
		BrokerUsername:             settings.BrokerUsername,
		BrokerToken:                settings.BrokerToken,
	}

	verifyResponses, err := pact.VerifyProvider(t, verifyRequest)
	if err != nil {
		fmt.Println("Error on VerifyProvider: ", err)
		data, _ := os.ReadFile("/log/pact.log")
		fmt.Println(string(data))
		t.Fatal(err)
	}

	fmt.Println(len(verifyResponses), "pact tests run")
}
