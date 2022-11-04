package reading

import (
	"fmt"
)

type Sensor interface {
	fmt.Stringer

	Read() (Reading, error)
}
