package main

import (
	"fmt"
)

func main() {
	var celsius float64
	fmt.Print("Enter temperature in Celsius: ")
	fmt.Scan(&celsius)

	fmt.Println("Convert Celsius to:")
	fmt.Println("1. Fahrenheit")
	fmt.Println("2. Kelvin")
	fmt.Println("3. Reaumur")

	var choice int
	fmt.Print("Choose an option (1-3): ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		fahrenheit := (celsius * 9 / 5) + 32
		fmt.Printf("%.2f Celsius = %.2f Fahrenheit\n", celsius, fahrenheit)
	case 2:
		kelvin := celsius + 273.15
		fmt.Printf("%.2f Celsius = %.2f Kelvin\n", celsius, kelvin)
	case 3:
		reaumur := celsius * 4 / 5
		fmt.Printf("%.2f Celsius = %.2f Reaumur\n", celsius, reaumur)
	default:
		fmt.Println("Invalid option.")
	}
}
