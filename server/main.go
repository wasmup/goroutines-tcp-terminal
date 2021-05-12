package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"sync"
	"time"
)

func main() {
	fmt.Println("Now Run 2 client terminals...")
	ln, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
	w1, err := ln.Accept()
	if err != nil {
		log.Fatal(err)
	}
	defer w1.Close()
	fmt.Println("Terminal 1 connected.")
	w2, err := ln.Accept()
	if err != nil {
		log.Fatal(err)
	}
	defer w2.Close()
	fmt.Println("Terminal 2 connected.")

	// your code:
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	wg := &sync.WaitGroup{}

	wg.Add(2)

	go func() {
		for i := 0; i < 100; i++ {
			time.Sleep(time.Microsecond * time.Duration(r1.Intn(100)))
			fmt.Fprintln(w1, "T1 : ", i)
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 100; i++ {
			time.Sleep(time.Microsecond * time.Duration(r1.Intn(100)))
			fmt.Fprintln(w2, "T2 : ", i)
		}
		wg.Done()
	}()

	wg.Wait()
}
