package main

import "fmt"

func main() {
	const usdEuro float64 = 0.87
	const usdRub float64 = 79.26

	const euroRub float64 = usdRub / usdEuro

	fmt.Println(euroRub) 
}
