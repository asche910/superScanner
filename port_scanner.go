package main

import (
	"fmt"
	"net"
	"os"
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
		fmt.Println(ip, port, err)
		if strings.Contains(err.Error(), "too many open files") {
			fmt.Println("连接数超出系统限制！" + err.Error())
			os.Exit(1)
		}
		return false
	}
	_ = conn.Close()
	return true
}

func CheckMultiPort(ip string, from , to int) []int {
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
			time.Sleep(10 * time.Millisecond)
			if CheckPort(ip, idx) {
				arr = append(arr, idx)
			}
		}()
	}
	wg.Wait()
	sort.Ints(arr)
	return arr
}