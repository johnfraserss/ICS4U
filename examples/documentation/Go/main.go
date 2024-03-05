package main

import "fmt"

/*
InchesToCentimeters converts a quantity in inches to centimeters.

Parameter: inches, float32 -> the value of the measurement to convert

Returns: float32 -> the converted value
*/
func InchesToCentimeters(inches float32) float32 {
    return inches * 2.54
}

func main() {
    fmt.Println("Hello, world!")
    fmt.Printf("1in = %.2fcm", InchesToCentimeters(1))
}
