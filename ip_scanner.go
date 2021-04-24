package main

import (
	"fmt"
	"github.com/tatsushid/go-fastping"
	"log"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"
)



// CheckIpRange
// "114.115.150-160.1-60"
// "114.115.157.1-60"
func CheckIpRange(rawIp string) []string {
	var arr [4][]int


	var aliveIp []string
	var mutex sync.Mutex
	var wg = sync.WaitGroup{}


	subs := strings.Split(rawIp, ".")
	for i := 0; i < 4; i += 1 {
		idx := strings.Index(subs[i], "-")
		if idx == -1 {
			num, err := strconv.Atoi(subs[i])
			if err != nil {
				log.Println("ip illegal!")
				return nil
			}
			arr[i] = append(arr[i], num)
		} else {
			from, err1 := strconv.Atoi(subs[i][:idx])
			to, err2 := strconv.Atoi(subs[i][idx+1:])
			if err1 != nil || err2 != nil {
				log.Println("ip illegal!")
				return nil
			}
			if from > to {
				log.Println("ip illegal!")
				return nil
			}
			arr[i] = append(arr[i], from, to)
		}
	}

	fmt.Println(arr)

	var temp []int
	ipDFS(arr, 0, temp, &aliveIp, &mutex, &wg)
	wg.Wait()
	//fmt.Println(aliveIp)
	return aliveIp
}

func ipDFS(arr [4][]int, idx int, subIp []int, aliveIp *[]string, mutex *sync.Mutex, wg *sync.WaitGroup) {
	if len(subIp) == 4 {
		ipStr := fmt.Sprintf("%d.%d.%d.%d", subIp[0], subIp[1], subIp[2], subIp[3])
		fmt.Println(ipStr)
		wg.Add(1)
		go checkAndAdd(ipStr, aliveIp, mutex, wg)
		return
	}

	if len(arr[idx]) == 1 {
		subIp = append(subIp, arr[idx][0])
		ipDFS(arr, idx+1, subIp, aliveIp, mutex, wg)
	} else {
		for i := arr[idx][0]; i <= arr[idx][1]; i++ {
			subIp = append(subIp, i)
			ipDFS(arr, idx+1, subIp, aliveIp, mutex, wg)
			subIp = subIp[0 : len(subIp)-1]
		}
	}
}

func checkAndAdd(ip string, aliveIp *[]string, mutex *sync.Mutex, wg *sync.WaitGroup)  {
	defer wg.Done()
	if CheckIp(ip){
		fmt.Println(ip)
		mutex.Lock()
		*aliveIp = append(*aliveIp, ip)
		mutex.Unlock()
	}
}

// CheckIp check ip or domain available
func CheckIp(ip string) bool {
	isAlive := false
	p := fastping.NewPinger()
	ra, err := net.ResolveIPAddr("ip4:icmp", ip)
	if err != nil {
		fmt.Println(err)
		return false
	}
	p.AddIPAddr(ra)
	p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
		fmt.Printf("IP Addr: %s receive, RTT: %v\n", addr.String(), rtt)
		isAlive = true
	}
	//p.OnIdle = func() {
	//	fmt.Println("finish")
	//}
	err = p.Run()
	if err != nil {
		fmt.Println(err)
	}
	return isAlive
}
