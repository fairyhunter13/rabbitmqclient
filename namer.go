package rabbitmqclient

import (
	"fmt"
)

// GenerateExchangeName generates the default exchange name of this library.
func GenerateExchangeName(isPrefix bool, name string) string {
	if isPrefix {
		return fmt.Sprintf("%s%s", DefaultPrefixExchange, name)
	}
	return name
}
