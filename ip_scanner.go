package main

import (
	"fmt"
	"github.com/tatsushid/go-fastping"
	"net"
	"time"
)

func CheckIp(ip string) bool {
	isAlive := false
	p := fastping.NewPinger()
	ra, err := net.ResolveIPAddr("ip4:icmp", ip)
	if err != nil {
		fmt.Println(err)
		return false
	}
	p.AddIPAddr(ra)
	p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
		fmt.Printf("IP Addr: %s receive, RTT: %v\n", addr.String(), rtt)
		isAlive = true
	}
	//p.OnIdle = func() {
	//	fmt.Println("finish")
	//}
	err = p.Run()
	if err != nil {
		fmt.Println(err)
	}
	return isAlive
}