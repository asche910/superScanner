package main

import (
	"fmt"
)

func main() {
	fmt.Println("SuperScanner start.")

	//isOpen := CheckPort("google.com", 1)
	//fmt.Println(isOpen)

	flag := CheckMultiPort("asche.top", 70, 999)
	fmt.Println(flag)
}
