package main

import "fmt"

const (
	usdeur = 0.85
	usdrub = 78.17
)

func main() {
	var eurrub = usdrub / usdeur
	fmt.Println(eurrub)
}

func readinput() string {
	var sc string
	fmt.Scan(&sc)
	return sc
}

func cunt(int, string, string) int {
	return 0
}
