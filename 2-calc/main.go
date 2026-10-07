package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	operation := getOperation()
	numbers := getNumbers()
	newNumbers, err := parseToInt(numbers)
	if err != nil {
		return
	}
	num := calculate(newNumbers, operation)
	if num == 0 && operation != "sum" && operation != "avg" && operation != "med" {
		return
	}

	fmt.Printf("%s равна: %d\n", strings.ToUpper(operation), num)
}

func parseToInt(numbers []string) (newNumbers []int, err error) {
	newNumbers = make([]int, 0, len(numbers))
	for _, value := range numbers {
		cleanValue := strings.TrimSpace(value)
		if cleanValue == "" {
			continue
		}

		var num int
		num, err = strconv.Atoi(cleanValue)
		if err != nil {
			fmt.Println("Вы ввели не целые числа:", err)
			return nil, err
		}

		newNumbers = append(newNumbers, num)
	}

	return newNumbers, nil
}

func getOperation() string {
	var operation string
	fmt.Print("Введите одну из операции AVG/SUM/MED: ")
	fmt.Scanln(&operation)

	return strings.ToLower(strings.TrimSpace(operation))
}

func getNumbers() []string {
	fmt.Print("Введите необходимую последовательность чисел через запятую: ")
	
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	numbers := scanner.Text()

	return strings.Split(numbers, ",")
}

func calculate(numbers []int, operation string) (num int) {
	if len(numbers) == 0 {
		return 0
	}

	switch operation {
	case "sum":
		for _, value := range numbers {
			num += value
		}
	case "avg":
		for _, value := range numbers {
			num += value
		}
		num = num / len(numbers)
	case "med":
		slices.Sort(numbers)
		n := len(numbers)
		if n%2 != 0 {
			num = numbers[n/2]
		} else {
			mid1 := numbers[n/2-1]
			mid2 := numbers[n/2]
			num = (mid1 + mid2) / 2
		}
	default:
		fmt.Println("Вы ввели неправильную операцию")
	}

	return num
}


