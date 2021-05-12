package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
	con, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer con.Close()

	buf := make([]byte, 1024)
	for {

		n, err := con.Read(buf)
		if err != nil {
			break
		}
		fmt.Print(string(buf[:n]))
	}

}
