package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/http"
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

func DomainToIP(domain string) string {
	ra, _ := net.ResolveIPAddr("ip4:icmp", domain)
	return ra.IP.String()
}

func CheckHttps(domain string) bool {
	domain = "https://" + domain
	_, err := http.Get(domain)
	if err != nil {
		return false
	}
	return true
}

func RemoveDuplication(rawArr []string) []string {
	var res []string
	var mp = map[string]int{}

	for _, item := range rawArr {
		if _, ok := mp[item]; !ok {
			mp[item] = 1
			res = append(res, item)
		}
	}
	return res
}

func TempTest()  {
	domain := "asche.top"
	ra, _ := net.ResolveIPAddr("ip4:icmp", domain)

	fmt.Println(ra.IP.String())
}

func funA(strs *[]string)  {
	fmt.Println(strs)
	*strs = append(*strs, "222")
}