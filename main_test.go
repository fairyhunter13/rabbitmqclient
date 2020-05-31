package rabbitmqclient

import (
	"flag"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/fairyhunter13/amqpwrapper"
	"github.com/streadway/amqp"
)

const (
	uriDialTemplate = "amqp://guest:guest@%s:5672"
)

type TestSetup struct {
	amqpConnectionManager amqpwrapper.IConnectionManager
}

func newTestSetup() *TestSetup {
	return &TestSetup{}
}

func (t *TestSetup) InitConnection(uriHost string) *TestSetup {
	if uriHost == "" {
		uriHost = "localhost"
	}
	if envHost := os.Getenv("RABBITMQ_HOST"); envHost != "" {
		uriHost = envHost
	}
	uriHost = fmt.Sprintf(uriDialTemplate, uriHost)
	var err error
	t.amqpConnectionManager, err = amqpwrapper.NewManager(uriHost, amqp.Config{})
	if err != nil {
		log.Panicf("Error: %+v", err)
		os.Exit(1)
	}
	return t
}

func (t *TestSetup) GetConnection() amqpwrapper.IConnectionManager {
	return t.amqpConnectionManager
}

var (
	testSetup *TestSetup
)

func initTestSetup() {
	testSetup = newTestSetup()
	testSetup = testSetup.InitConnection("localhost")
}

func TestMain(m *testing.M) {
	flag.Parse()
	initTestSetup()
	os.Exit(m.Run())
}
