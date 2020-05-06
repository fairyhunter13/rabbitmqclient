package rabbitmqclient

import "errors"

// List of all errors for connection.
var (
	ErrConnectionAlreadyClosed   = errors.New("Connection is already closed")
	ErrContainerMustBeSavedFirst = errors.New("Container must be saved first before used")
	ErrTopologyMustNotBeNil      = errors.New("Topology mustn't be nil")
)
