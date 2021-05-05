package main

import (
	"fmt"
	"net"
	"sort"
	"strings"
	"sync"
	"time"
)

func Hello() {
	fmt.Println("hhhhhh")
}

func CheckPort(ip string, port int) bool {
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", ip, port),
		time.Second * TimeOut)
	if err != nil {
		//fmt.Println(ip, port, err)
		if strings.Contains(err.Error(), "too many open files") {
			fmt.Println("超过最大连接数！" + err.Error())
		}
		return false
	}
	_ = conn.Close()
	return true
}

func CheckMultiPort(ip string, from , to int) []int {
	fmt.Println(ip, from, to)
	var arr []int
	if from > to {
		return arr
	}
	if from < 1{
		from = 1
	}
	if to > 65535 {
		to = 65535
	}
	var wg = sync.WaitGroup{}
	for i := from; i <= to; i += 1 {
		wg.Add(1)
		idx := i
		go func() {
			defer wg.Done()
			time.Sleep(1 * time.Millisecond)
			if CheckPort(ip, idx) {
				arr = append(arr, idx)
			}
		}()
	}
	wg.Wait()
	sort.Ints(arr)
	return arr
}