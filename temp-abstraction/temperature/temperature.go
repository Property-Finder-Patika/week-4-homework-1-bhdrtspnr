package tempFunction

var K_TO_C_CONSTANT float64 = 273.15

type Temperature interface {
	GetName()
	GetTemperature()
	Convert()
}

type Kelvin struct {
	Name string
	Temp float64
}

type Celsius struct {
	Name string
	Temp float64
}

func (k Kelvin) GetName() string {
	return k.Name
}

func (c Celsius) GetName() string {
	return c.Name
}

func (k Kelvin) GetTemperature() float64 {
	return k.Temp
}

func (c Celsius) GetTemperature() float64 {
	return c.Temp
}

func (k Kelvin) Convert() float64 {
	return k.Temp - K_TO_C_CONSTANT
}

func (c Celsius) Convert() float64 {
	return c.Temp + K_TO_C_CONSTANT
}
