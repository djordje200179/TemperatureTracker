package reading

import "fmt"

type Temperature float64

func (temp Temperature) String() string {
	return fmt.Sprintf("%d°C", temp.Round())
}

func (temp Temperature) Celsius() float64 {
	return float64(temp)
}

func (temp Temperature) Fahrenheit() float64 {
	return temp.Celsius()*1.8 + 32
}

func (temp Temperature) Kelvin() float64 {
	return temp.Celsius() + 273.15
}

func (temp Temperature) Round() int {
	return int(temp + 0.5)
}
