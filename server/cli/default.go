package cli

import (
	"os"
)

var defaultInstance = Start(nil, os.Stdin, os.Stdout)

func Default() *CLI {
	return defaultInstance
}
