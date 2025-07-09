package main

import (
	"fmt"
)

const (
	usdeur = 0.85
	usdrub = 78.17
)

var eurrub = usdrub / usdeur

func main() {

	fmt.Println(`Выберите нужную валюту из которой происходит конвертация или введите "quit" для выхода`)
	choises_printer("")
	var start string

	for {
		start = readcurinput("")
		if start == "quit" {
			break
		}
		if start == "Ошибка: Валюта введена неверно, повторите выбор:" {
			fmt.Println("Ошибка: Валюта введена неверно, повторите выбор:")
			continue
		} else {
			break
		}
	}

	fmt.Printf("Выберите количество конвертируемых %v:\n", start)
	var num float64
	for {
		num = readnumbinput()
		if num == 0.0 {
			fmt.Printf("Ошибка, введено неправильное количество %v , повторите выбор\n", start)
		} else {
			break
		}
	}

	fmt.Printf("Выберите валюту в которую будет происходить конвертация %v %v:\n", num, start)
	choises_printer(start)
	var finish string

	for {
		finish = readcurinput(start)
		if finish == "Ошибка: Валюта введена неверно, повторите выбор:" {
			fmt.Println("Ошибка: Валюта введена неверно, повторите выбор:")
			continue
		} else {
			break
		}
	}

	result := converter(num, start, finish)
	fmt.Printf("%v %v равняются %.2f %v\n", num, start, result, finish)
}

func readcurinput(ignore string) string {
	var sc string
	fmt.Scan(&sc)
	if (sc == "eur" || sc == "rub" || sc == "usd" || sc == "quit") && sc != ignore {
		return sc
	} else {
		return "Ошибка: Валюта введена неверно, повторите выбор:"
	}
}

func readnumbinput() float64 {
	var numb float64
	fmt.Scan(&numb)

	return numb
}

func choises_printer(except string) {
	choises := []string{"eur", "rub", "usd"}
	for _, value := range choises {
		if value != except {
			fmt.Println(value)
		}
	}
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
