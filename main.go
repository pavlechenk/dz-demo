package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	UsdToRub = 79.4637
	EurToRub = 91.1925
	EurToUsd = 1.15325

	UsdToEur = 0.867115
	RubToUsd = 0.012584
	RubToEur = 0.010966
)


func main() {
	sourceCurrency := readCurrency()
	destinationCurrency := readCurrency()
	amount := readAmount(sourceCurrency)
	convertCurrency(amount, sourceCurrency, destinationCurrency)
}

func readCurrency() string {
	var currency string
	for {
		fmt.Print("Введите предложенную валюту USD/EUR/RUB: ")
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

func convertCurrency(amount float64, from, to string) {
	switch {
		case from == "usd" && to == "rub":
			fmt.Printf("Сумма USD в RUB: %f\n", amount * UsdToRub)
		case from == "eur" && to == "rub":
			fmt.Printf("Сумма EUR в RUB: %f\n", amount * EurToRub)
		case from == "eur" && to == "usd":
			fmt.Printf("Сумма EUR в USD: %f\n", amount * EurToUsd)
		case from == "usd" && to == "eur":
			fmt.Printf("Сумма USD в EUR: %f\n", amount * UsdToEur)
		case from == "rub" && to == "usd":
			fmt.Printf("Сумма RUB в USD: %f\n", amount * RubToUsd)
		case from == "rub" && to == "eur":
			fmt.Printf("Сумма RUB в EUR: %f\n", amount * RubToEur)
		default: {
			fmt.Printf("Переданы неравильные валюты")
		}
	}
}