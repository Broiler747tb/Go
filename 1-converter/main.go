package main

import "fmt"

const (
	usdeur = 0.85
	usdrub = 78.17
)

var eurrub = usdrub / usdeur

func main() {
	for {
		gog := readcurinput()
		if gog == "Error: No such currency, try again." {
			fmt.Println("Error: No such currency, try again.")
			continue
		}
		ts :=
	}
}

func readcurinput() string {
	var sc string
	fmt.Scan(&sc)
	if sc == "eur" || sc == "rub" || sc == "usd" {
		return sc
	} else {
		return "Error: No such currency, try again."
	}
}

func readnumbinput() float64 {
	var numb float64
	fmt.Scan(&numb)
	
	return numb
}

func converter(numb float64, start string, finish string) float64 {
	if start == "usd" && finish == "eur" {
		return numb * usdeur
	}
	if start == "usd" && finish == "rub" {
		return numb * usdrub
	}
	if start == "eur" && finish == "rub" {
		return numb * eurrub
	}
	if finish == "usd" && start == "eur" {
		return numb / usdeur
	}
	if finish == "usd" && start == "rub" {
		return numb / usdrub
	}
	if finish == "eur" && start == "rub" {
		return numb / eurrub
	}
	return 0.0
}
