package main

import "fmt"

func main() {
	var high int
	var low int
	fmt.Print("Enter high then low numbers:")
	fmt.Scan(&high, &low)
	//fmt.Print("Enter high number:")
	//fmt.Scan(&high)
	fmt.Println(high % low)
}
