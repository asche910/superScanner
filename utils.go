package main

import (
	"bufio"
	"log"
	"os"
)

type Pair struct {
	a, b interface{}
}

func CheckFileExists(name string) bool {
	if _, err := os.Stat(name); err != nil{
		if os.IsNotExist(err){
			return false
		}
	}
	return true
}

func FileToList(name string) []string{
	file, err := os.Open(name)
	if err != nil{
		log.Println(err)
		return nil
	}
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var res []string
	for scanner.Scan(){
		res = append(res, scanner.Text())
	}
	return res
}