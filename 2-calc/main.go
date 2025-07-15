package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	choiceHandler()
}

func inputHandler() []float64 {
	fmt.Println("Введите набор чисел через запятую:")
	var input string
	fmt.Scan(&input)
	stringslice := strings.Split(input, ",")
	finalslice := []float64{}
	for _, value := range stringslice {
		numb, _ := strconv.ParseFloat(value, 64)
		finalslice = append(finalslice, numb)
	}
	return finalslice
}

func avgPrint(slice []float64) {
	var sum float64
	var count int
	for index, value := range slice {
		count = index + 1
		sum = sum + value
	}
	fmt.Println(sum / float64(count))
}

func sumPrint(slice []float64) {
	var sum float64
	for _, value := range slice {
		sum = sum + value
	}
	fmt.Println(sum)
}

func medPrint(slice []float64) {
	sort.Float64s(slice)
	n := len(slice)
	if n == 0 {
		fmt.Println(0)
	}
	if n%2 == 1 {
		fmt.Println(slice[n/2])
	} else {
		fmt.Println((slice[n/2-1] + slice[n/2]) / 2)
	}

}

func choiceHandler() {
	fmt.Printf(`~~~~~~~ Калькулятор V0.1 ~~~~~~~

Выберете нужную операцию:
   "AVG" - среднее арифметическое чисел
   "SUM" - сумма чисел
   "MED" - медиана чисел
`)
	var choice string
	fmt.Scan(&choice)
	switch choice {
	case "AVG":
		avgPrint(inputHandler())
	case "SUM":
		sumPrint(inputHandler())
	case "MED":
		medPrint(inputHandler())
	}
}
