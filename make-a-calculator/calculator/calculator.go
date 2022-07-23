package main

import (
	"errors"
	"fmt"
	mathFunction "maino/make-a-calculator/math"
	"strings"
)

func main() {
	startCalculator()
}

var appLogger = newLogger()

type Calculator struct {
	name      string
	functions []mathFunction.MathFunction
}

func (c *Calculator) doCalculation(name string, arg float64) (float64, error) {
	if !mathFunction.CheckValidFunction(name) {
		appLogger.Warning().Printf("doCalculation: User input is invalid!: %s\n", name)
		err := errors.New(invalidFunctionName(name))
		return -1, err
	}

	var result float64
	for _, f := range c.functions {
		if name == f.GetName() {
			result = f.Calculate(arg)
			appLogger.Info().Printf("doCalculation: Result of the calculation is: : %f\n", result)
		}
	}
	return result, nil
}

var myCalculator Calculator = Calculator{name: "Test"}

func startCalculator() {

	appLogger.Info().Printf("Calculator created:  %v \n", myCalculator)

	var userMenuInput string
	var menuFlag bool = true

	for menuFlag {
		fmt.Println("Enter 1 to add a new function to calculator")
		fmt.Println("Enter 2 to use your calculator.")
		fmt.Println("Enter 3 to see your functions available.")
		fmt.Scanln(&userMenuInput)

		appLogger.Info().Printf("User input for menu is:  %s \n", userMenuInput)

		switch userMenuInput {
		case "1":
			addFunctionMenu()
		case "2":
			doCalculationMenu()
		case "3":
			fmt.Printf("%v \n", myCalculator.viewFunctions())

		default:
			fmt.Println("Please input either 1 or 2!")
		}
	}

}

func doCalculationMenu() {
	appLogger.Info().Printf("Opening calculator menu...\n")
	fmt.Printf("Please select an function, current available functions: %v\n Use the index of your operation.\n", myCalculator.functions)
	var calculationMenuInput int
	var funcValue float64
	fmt.Scanln(&calculationMenuInput)
	fmt.Println("Please enter the value for your function: ")
	fmt.Scanln(&funcValue)
	result, _ := myCalculator.doCalculation(myCalculator.functions[calculationMenuInput].GetName(), funcValue)
	if result != -1 {
		fmt.Printf("Result is: %f \n", result)
	}
}

func addFunctionMenu() {
	appLogger.Info().Printf("Opening adding a new function menu...\n")
	fmt.Println("Please enter the name of the function you'd like to add")
	appLogger.Info().Printf("Printing valid functions... %s \n", mathFunction.GetValidFunctions())
	fmt.Printf("Valid functions are: %s: \n", mathFunction.GetValidFunctions())
	var addFunctionInput string
	fmt.Scanln(&addFunctionInput)
	appLogger.Info().Printf("User input for adding a function is %s...\n", addFunctionInput)

	appLogger.Info().Printf("User called add function method with parameter %s", addFunctionInput)
	returnMessage, _ := myCalculator.addFunction(addFunctionInput)
	fmt.Println(*returnMessage)
}

func (c *Calculator) viewFunctions() []string {
	appLogger.Info().Printf("Opening view menu...\n")
	functionNames := make([]string, 0)
	for i := 0; i < len(c.functions); i++ {
		functionNames = append(functionNames, c.functions[i].GetName())
		appLogger.Info().Printf("Appending : %s \n", c.functions[i].GetName())
	}
	return functionNames

}

func (c *Calculator) addFunction(name string) (*string, error) {
	name = strings.ToLower(name)
	if !mathFunction.CheckValidFunction(name) {
		appLogger.Warning().Printf("addFunction: User input is invalid!: %s\n", name)
		err := errors.New(invalidFunctionName(name))
		return nil, err
	}
	appLogger.Info().Printf("User input in valid: %s\n", name)
	successMessage := fmt.Sprintf("Successfully added %s to the calculator %v \n", name, c.name)

	switch name {
	case "sin":
		c.functions = append(c.functions, mathFunction.Sin{Name: "sin"})
		appLogger.Info().Printf("User matches with sin: %s ... adding a new function to calculator %v \n", name, c)
		appLogger.Info().Printf("Calculator %v's current functions are: %v \n", c, c.functions)
	case "cos":
		c.functions = append(c.functions, mathFunction.Cos{Name: "cos"})
		appLogger.Info().Printf("User matches with cos: %s ... adding a new function to calculator %v \n", name, c)
		appLogger.Info().Printf("Calculator %v's current functions are: %v \n", c, c.functions)
	case "log":
		c.functions = append(c.functions, mathFunction.Log{Name: "log"})
		appLogger.Info().Printf("User matches with log: %s ... adding a new function to calculator %v \n", name, c)
		appLogger.Info().Printf("Calculator %v's current functions are: %v \n", c, c.functions)
	default:
		appLogger.Fatal().Printf("UNKNOWN ERROR HAPPENED!")
	}
	return &successMessage, nil

}
