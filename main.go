package main

import (
	"fmt"
	"math"
	"net"
	"os"
	"strconv"
	"time"
)

func scanports(start, end int, host string) {
	for ; start < end; start++ {
		address := net.JoinHostPort(host, strconv.Itoa(start))

		conn, err := net.DialTimeout("tcp", address, 3*time.Second)

		if err == nil && conn != nil {
			fmt.Println(host, start, "is open")
			conn.Close()
		}
	}
}

func main() {
	ports := math.Pow(2, 16)
	threads := 100
	var s, e int
	host := os.Args[1]

	distance := (int(ports) / threads)
	for i := 0; i < threads-1; i++ {
		s = i * distance
		e = s + distance
		go scanports(s, e, host)
	}

	scanports(int(ports)-(int(ports)/threads), int(ports), host)
}
