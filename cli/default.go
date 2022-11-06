package cli

import (
	"os"
)

var defaultInstance = New(nil, os.Stdin, os.Stdout)

func Default() *CLI {
	return defaultInstance
}
