package reading

import "fmt"

type Humidity float64

func (humidity Humidity) String() string {
	if humidity <= 0 {
		return "NaN"
	}

	return fmt.Sprintf("%d%%", humidity.Round())
}

func (humidity Humidity) Percent() float64 {
	return float64(humidity)
}

func (humidity Humidity) Round() int {
	return int(humidity + 0.5)
}
