package main

import (
	"fmt"
	temperature "maino/temp-abstraction/temperature"
)

func main() {

	fmt.Println("1 for Kelvin 2 for celsius")

	var input int
	fmt.Scanln(&input)

	switch input {
	case 1:
		name, temp := getNameAndTemp()
		newTempK := temperature.Kelvin{Name: name, Temp: temp}
		fmt.Printf("New Kelvin's name is: %s \n", newTempK.GetName())
		fmt.Printf("New Kelvin's temp is %0.2f \n", newTempK.GetTemperature())
		fmt.Printf("New Kelvin's temp converted to C is %0.2f\n", newTempK.Convert())
	case 2:
		name, temp := getNameAndTemp()
		newTempC := temperature.Celsius{Name: name, Temp: temp}
		fmt.Printf("New Celcius' name is: %s \n", newTempC.GetName())
		fmt.Printf("New Celcius' temp is %0.2f \n", newTempC.GetTemperature())
		fmt.Printf("New Celcius' temp converted to C is %0.2f\n", newTempC.Convert())

	default:
		fmt.Println("Invalid input")

	}
}

func getNameAndTemp() (string, float64) {
	var name string
	var temp float64
	fmt.Println("Please enter a name:")
	fmt.Scanln(&name)
	fmt.Println("Please enter a temp: ")
	fmt.Scanln(&temp)

	return name, temp
}
