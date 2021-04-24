package main

import (
	"fmt"
	"testing"
)

func TestCheckMultiPort(t *testing.T) {
	ip := "asche.top"
	ports := CheckMultiPort(ip, 1, 1000)
	fmt.Println(ports)
}