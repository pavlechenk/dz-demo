package main

import "fmt"

// const usdEuro float64 = 0.87
// const usdRub float64 = 79.26

func main() {
	usdEuro, usdRub := getUserInput()
	var euroRub float64 = usdRub / usdEuro

	fmt.Println(euroRub)
}

func getUserInput() (usdEuro, usdRub float64) {
	fmt.Print("Введите курс доллара к евро: ")
	fmt.Scan(&usdEuro)
	fmt.Print("Введите курс доллара к рублю: ")
	fmt.Scan(&usdRub)
	return
}

func convertCurrency(amount float64, from, to string) {

}