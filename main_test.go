package rabbitmqclient

import (
	"flag"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/fairyhunter13/amqpwrapper"
	"github.com/streadway/amqp"
)

const (
	uriDialTemplate = "amqp://guest:guest@%s:5672"
)

const (
	normalTimeSleep = 2 * time.Second
)

type TestSetup struct {
	amqpConnectionManager amqpwrapper.IConnectionManager
}

func newTestSetup() *TestSetup {
	return &TestSetup{}
}

func (t *TestSetup) InitConnection(uriHost string) *TestSetup {
	t.amqpConnectionManager = t.GetNewConnection(uriHost)
	return t
}

func (t *TestSetup) NewContainer() (*Container, error) {
	return NewContainer(t.GetConnection())
}

func (t *TestSetup) NewContainerAndConnection() (*Container, error) {
	return NewContainer(t.GetNewConnection(""))
}

func (t *TestSetup) GetConnection() amqpwrapper.IConnectionManager {
	return t.amqpConnectionManager
}

func (t *TestSetup) GetNewConnection(uriHost string) amqpwrapper.IConnectionManager {
	if uriHost == "" {
		uriHost = "localhost"
	}
	if envHost := os.Getenv("RABBITMQ_HOST"); envHost != "" {
		uriHost = envHost
	}
	uriHost = fmt.Sprintf(uriDialTemplate, uriHost)
	var (
		err  error
		conn amqpwrapper.IConnectionManager
	)
	conn, err = amqpwrapper.NewManager(uriHost, amqp.Config{})
	if err != nil {
		log.Panicf("Error: %+v", err)
		os.Exit(1)
	}
	return conn
}

var (
	testSetup *TestSetup
)

func initTestSetup() {
	testSetup = newTestSetup()
	testSetup = testSetup.InitConnection("localhost")
}

func teardownTestSetup() {
	testSetup.amqpConnectionManager.Close()
}

func TestMain(m *testing.M) {
	flag.Parse()
	initTestSetup()
	code := m.Run()
	teardownTestSetup()
	os.Exit(code)
}
