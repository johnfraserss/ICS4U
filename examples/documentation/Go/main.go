package main

import "fmt"

// InchesToCentimeters returns inches in centimeters.
func InchesToCentimeters(inches float32) float32 {
	return inches * 2.54
}

func main() {
	fmt.Println("Hello, world!")
	fmt.Printf("1in = %.2fcm", InchesToCentimeters(1))
}
