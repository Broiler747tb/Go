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
	choisesPrinter("")
	var start string

	for {
		start = readCurInput("")
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
		num = readNumbInput()
		if num == 0.0 {
			fmt.Printf("Ошибка, введено неправильное количество %v , повторите выбор\n", start)
		} else {
			break
		}
	}

	fmt.Printf("Выберите валюту в которую будет происходить конвертация %v %v:\n", num, start)
	choisesPrinter(start)
	var finish string

	for {
		finish = readCurInput(start)
		if finish == "Ошибка: Валюта введена неверно, повторите выбор:" {
			fmt.Println("Ошибка: Валюта введена неверно, повторите выбор:")
			continue
		} else {
			break
		}
	}

	result := converter(num, start, finish, &convertmap)
	fmt.Printf("%v %v равняются %.2f %v\n", num, start, result, finish)
}

func readCurInput(ignore string) string {
	var sc string
	fmt.Scan(&sc)
	if (sc == "eur" || sc == "rub" || sc == "usd" || sc == "quit") && sc != ignore {
		return sc
	} else {
		return "Ошибка: Валюта введена неверно, повторите выбор:"
	}
}

func readNumbInput() float64 {
	var numb float64
	fmt.Scan(&numb)

	return numb
}

func choisesPrinter(except string) {
	choises := []string{"eur", "rub", "usd"}
	for _, value := range choises {
		if value != except {
			fmt.Println(value)
		}
	}
}

func converter(numb float64, start string, finish string, convertmap *map[string]float64) float64 {
	coefficient := *convertmap
	return numb * coefficient[start+finish]
}

var convertmap = map[string]float64{
	"usdeur": usdeur,
	"usdrub": usdrub,
	"eurrub": eurrub,
	"eurusd": 1 / usdeur,
	"rubusd": 1 / usdrub,
	"rubeur": 1 / eurrub,
}
