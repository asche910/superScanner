package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"sync"
	"time"
)

func ResultController(context *gin.Context) {
	domain := context.Query("domain")
	fmt.Println("param: ", domain)

	var wg sync.WaitGroup
	var ip string
	var ports []int
	var domains []string
	var https bool

	go func(ip *string, wg *sync.WaitGroup) {
		defer (*wg).Done()
		(*wg).Add(1)

		*ip = DomainToIP(domain)
	}(&ip, &wg)

	go func(ports *[]int, wg *sync.WaitGroup) {
		defer (*wg).Done()
		(*wg).Add(1)

		*ports = CheckMultiPort(domain, 1, 1024)
	}(&ports, &wg)

	go func(domains *[]string, wg *sync.WaitGroup) {
		defer (*wg).Done()
		(*wg).Add(1)

		*domains = DomainScan(domain)
	}(&domains, &wg)

	go func(https *bool, wg *sync.WaitGroup) {
		defer (*wg).Done()
		(*wg).Add(1)

		*https = CheckHttps(domain)
	}(&https, &wg)

	time.Sleep(5 * time.Millisecond)
	wg.Wait()

	data := map[string]interface{}{
		"domain":  domain,
		"ip":      ip,
		"https":   https,
		"ports":   ports,
		"domains": domains,
	}

	context.HTML(200, "result.html", data)
}
