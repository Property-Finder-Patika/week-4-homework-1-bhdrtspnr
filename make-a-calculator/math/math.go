package mathFunction

import (
	"math"
)

var validFunctions = []string{"sin", "cos", "log"}

func CheckValidFunction(name string) bool {
	for i := 0; i < len(validFunctions); i++ {
		if validFunctions[i] == name {
			return true
		}
	}
	return false
}

func GetValidFunctions() []string {
	return validFunctions
}

type MathFunction interface {
	Calculate(arg float64) float64
	GetName() string
}

type Sin struct {
	Name string
}

type Cos struct {
	Name string
}

type Log struct {
	Name string
}

func (s Sin) GetName() string {
	return s.Name
}

func (c Cos) GetName() string {
	return c.Name
}

func (l Log) GetName() string {
	return l.Name
}

func (s Sin) Calculate(arg float64) float64 {
	return math.Sin(arg)
}

func (c Cos) Calculate(arg float64) float64 {
	return math.Cos(arg)
}

func (l Log) Calculate(arg float64) float64 {
	return math.Log(arg) //natural log e
}
