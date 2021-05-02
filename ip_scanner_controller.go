package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func IPResultController(context *gin.Context)  {
	rawIp := context.Query("ips")
	fmt.Println("param: ", rawIp)

	ips := CheckIpRange(rawIp)
	fmt.Println(ips)

	var data = map[string]interface{}{
		"rawIp": rawIp,
		"ips": ips,
	}

	context.HTML(200, "ip_result.html", data)
}