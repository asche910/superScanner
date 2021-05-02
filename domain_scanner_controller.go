package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func DomainResultController(context *gin.Context)  {
	domain := context.Query("domain")
	fmt.Println("param: ", domain)

	domains := DomainScan(domain)

	var data = map[string]interface{}{
		"domain": domain,
		"domains": domains,
	}

	context.HTML(200, "domain_result.html", data)
}