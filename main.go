package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("SuperScanner start.")

	//isOpen := CheckPort("google.com", 1)
	//fmt.Println(isOpen)

	//flag := CheckMultiPort("asche.top", 20, 3333)
	//fmt.Println(flag)

	//flag := CheckIp("114.115.157.53")
	//fmt.Println(flag)

	CheckIpRange("114-119.115.157-158.1-20")


	//time.Sleep(5 * time.Second)
}

func test()  {
	fmt.Println("test")
	go test2()
	fmt.Println("test done")
}

func test2()  {
	fmt.Println("start")
	time.Sleep(3 * time.Second)
	fmt.Println("done")
}