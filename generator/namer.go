package generator

import (
	"fmt"

	"github.com/fairyhunter13/rabbitmqclient/constant"
)

// GenerateExchangeName generates the default exchange name of this library.
func GenerateExchangeName(isPrefix bool, name string) string {
	if isPrefix {
		return fmt.Sprintf("%s%s", constant.DefaultPrefixExchange, name)
	}
	return name
}
