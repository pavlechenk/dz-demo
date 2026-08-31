package main

import (
	"fmt"
	"strconv"
	"strings"
)

type currency = map[string]float64


func main() {
	sourceCurrency := readCurrency("исходную")
	destinationCurrency := readCurrency("целевую")
	amount := readAmount(sourceCurrency)
	currencies := currency{
		"usdrub": 79.4637,
		"eurrub": 91.1925,
		"eurusd": 1.15325,
		"usdeur": 0.867115,
		"rubusd": 0.012584,
		"rubeur": 0.010966,
	}
	convertCurrency(amount, sourceCurrency, destinationCurrency, currencies)
}

func readCurrency(role string) string {
	var currency string
	for {
		fmt.Printf("Введите %s валюту USD/EUR/RUB: ", role)
		fmt.Scan(&currency)
		currency = strings.ToLower(strings.TrimSpace(currency))
		if checkCurrency(currency) {
			break
		} else {
			fmt.Println("Введенное значение не является валютой")
			continue
		}
	}

	return currency
}

func readAmount(baseCurrency string) float64 {
	var amount float64
	var input string
	var err error
	baseCurrency = strings.ToUpper(baseCurrency)
	for {
		fmt.Printf("Введите сумму в %s, которую хотите конвертировать: ", baseCurrency)
		fmt.Scan(&input)
		amount, err = strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Вы ввели неправильное число. Попробуйте снова")
			continue
		} else {
			break
		}
	}

	return amount
}

func checkCurrency(currency string) bool {
	return currency == "usd" || currency == "eur" || currency == "rub" 
}

func convertCurrency(amount float64, from, to string, currencies currency) {
	key := from + to
	fmt.Printf("Сумма %s в %s: %.2f\n", from, to, amount * currencies[key])
}