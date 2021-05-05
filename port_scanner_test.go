package main

import (
	"fmt"
	"testing"
)

func TestCheckMultiPort(t *testing.T) {
	ip := "asche.top"
	ports := CheckMultiPort(ip, 3000, 4000)
	fmt.Println(ports)
}