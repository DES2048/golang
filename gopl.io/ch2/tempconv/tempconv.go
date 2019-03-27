package tempconv

import "fmt"

type Fahrengeit float64
type Celcius float64
type Kalvin float64

const (
	AbsoluteZeroC Celcius = -273.15
	WaterBoilingPointC Celcius = 100
	WaterFreezePointC Celcius= 0
)

const (
	AbsoluteZeroF Fahrengeit = -459.67
	WaterBoilingPointF Fahrengeit = 212
	WaterFreezePointF Fahrengeit= 32
)

const (
	AbsoluteZeroK Kalvin = 0
	WaterFreezePointK Kalvin = 273.15
	WaterBoilingPointK Kalvin = 373.15
)

// специальные методы на подобии toString()
func (c Celcius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (f Fahrengeit) String() string {
	return fmt.Sprintf("%g°F", f)
}

func (k Kalvin) String() string {
	return fmt.Sprintf("%g°K", k)
}