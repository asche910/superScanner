package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func ResultController(context *gin.Context)  {
	domain := context.Query("domain")
	fmt.Println("param: ", domain)

	ip := DomainToIP(domain)
	ports := CheckMultiPort(domain, 1, 1024)
	domains := DomainScan(domain)
	https := CheckHttps(domain)


	data := map[string]interface{}{
		"domain": domain,
		"ip": ip,
		"https": https,
		"ports": ports,
		"domains": domains,
	}

	context.HTML(200, "result.html", data)
}