package main

import "fmt"

func main() {
	var usdTOEuro, usdToRub = usdTo()
	var euroToRub = usdToRub / usdTOEuro

	fmt.Printf("1 EUR = %.2f RUB\n", euroToRub)
}

func usdTo() (float64, float64) {
	var euro float64
	var rub float64
	fmt.Println("Калькулятор валют")
	fmt.Print("Введите курс USD к EUR: ")
	fmt.Scan(&euro)

	fmt.Print("Введите курс USD к RUB: ")
	fmt.Scan(&rub)
	return euro, rub
}

func convert(amount int, from float64, to float64) {

}
