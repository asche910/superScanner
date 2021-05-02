package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func PortResultController(context *gin.Context)  {
	ip := context.Query("ip")
	from := context.Query("from")
	to := context.Query("to")
	fmt.Println("param: ", ip, from, to)

	fromInt, _ := strconv.Atoi(from)
	toInt, _ := strconv.Atoi(to)
	ports := CheckMultiPort(ip, fromInt, toInt)

	data := map[string]interface{}{
		"ports": ports,
		"ip": ip,
		"from": from,
		"to": to,
	}
	context.HTML(200, "port_result.html", data)
}