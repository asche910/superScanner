package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"time"
)

func SubScan(domain string) {
	fileName := fmt.Sprintf("output/subdomain_%d.txt", time.Now().Unix())
	cmd := exec.Command("helper/subhelper", "-d",  domain, "-o", fileName)

	var out, stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil{
		log.Println(err)
	}
	//fmt.Println("output: ", out.String(), stderr.String())

	if CheckFileExists(fileName){
		ipList := FileToList(fileName)
		fmt.Println(ipList)
	}else{
		fmt.Println("scan failed!")
	}
}
