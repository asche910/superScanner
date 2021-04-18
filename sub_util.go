package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func SubScan(domain string) {
	cmd := exec.Command("helper/subhelper", "-d",  "asche.top", "-o", "subdomain.txt")
	//cmd := exec.Command("helper\\subhelper")

	var out, stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	//output, err := cmd.Output()
	err := cmd.Run()
	if err != nil{
		log.Println(err)
	}
	fmt.Println("output: ", string(out.String()), stderr.String())

}
