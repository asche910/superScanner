package main

import (
	"fmt"
	"testing"
)

func TestCheckIp(t *testing.T) {
	ip := "asche.top"
	flag := CheckIp(ip)
	fmt.Println(flag)
}

func TestCheckIpRange(t *testing.T) {
	rawIp := "114.115.157.1-60"
	aliveIps := CheckIpRange(rawIp)
	fmt.Println(aliveIps)
}