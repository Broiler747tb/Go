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
