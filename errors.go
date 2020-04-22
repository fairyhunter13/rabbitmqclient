package rabbitmqclient

import "errors"

// List of all errors for connection.
var (
	ErrConnectionAlreadyClosed = errors.New("Connection is already closed")
)
