package main

import (
	"fmt"
	"math"
)

func main() {
	const IMTPower float64 = 2
	var userHeight float64
	var userKG float64
	fmt.Print("Height calculator \n")
	fmt.Print("Type your Height:")
	fmt.Scan(&userHeight)
	fmt.Print("Type your weight:")
	fmt.Scan(&userKG)
	IMT := userKG / math.Pow(userHeight, IMTPower)
	fmt.Print("Your body index:")
	fmt.Print(IMT)
}
