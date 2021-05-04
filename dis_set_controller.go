package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

var SlaveSet = map[string]int{
	"localhost:" + strconv.Itoa(ServerPort): 1,
}

func DisSetController(context *gin.Context){
	var slaveList []string
	for key, _ := range SlaveSet {
		slaveList = append(slaveList, key)
	}
	data := map[string]interface{}{
		"slaves": slaveList,
	}
	context.HTML(200, "dis_set.html", data)
}

func AddSlaveController(context *gin.Context){
	slaveUrl := context.PostForm("slaveUrl")
	fmt.Println("slave url:", slaveUrl)
	res := map[string]interface{}{
		"code": 0,
		"msg": "success",
	}
	if slaveUrl == "" {
		res["code"] = 1
		res["msg"] = "slaveUrl is null!"
		context.JSON(200, res)
	}else{
		SlaveSet[slaveUrl] = 1
		context.JSON(200, res)
	}
}

func DeleteSalveController(context *gin.Context){
	slaveUrl := context.Query("slaveUrl")
	_, ok := SlaveSet[slaveUrl]
	if ok {
		delete(SlaveSet, slaveUrl)
	}
	res := map[string]interface{}{
		"code": 0,
		"data": nil,
	}
	context.JSON(200, res)
}

func GetAllSlaveController(context *gin.Context){
	var slaveList []string
	for key, _ := range SlaveSet {
		slaveList = append(slaveList, key)
	}
	data := map[string]interface{}{
		"code": 0,
		"data": slaveList,
	}
	context.JSON(200, data)
}
