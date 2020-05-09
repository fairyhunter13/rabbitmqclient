package rabbitmqclient

import "errors"

// List of all errors for rabbitmqclient.
var (
	ErrConnectionAlreadyClosed   = errors.New("Connection is already closed")
	ErrContainerMustBeSavedFirst = errors.New("Container must be saved first before used")
	ErrTopologyMustNotBeNil      = errors.New("Topology mustn't be nil")
	ErrArgumentsMusntBeEmpty     = errors.New("Arguments of function cannot be empty")
	ErrMethodNotFound            = errors.New("Method/Function is not found in the given struct")
	ErrInvalidFunctionCalled     = errors.New("Invalid function when called for declaration")
	ErrInvalidReturnValues       = errors.New(`Invalid return values for the function called`)
)
