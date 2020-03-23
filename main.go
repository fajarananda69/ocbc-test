package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

func main() {
	startTime := time.Now()
	n := 2
	runtime.GOMAXPROCS(n)
	dataset := []int64{1, 2, 4, 2, 3, 5, 2, 3, 1, 3}
	data := make(chan string, n)

	// handle termination signal
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("wait until finish")
		// os.Exit(1)
	}()

	hello := func(hello string, sleeps int64) {
		datas := fmt.Sprintf("hello %s %d", hello, sleeps)
		time.Sleep(time.Duration(sleeps) * time.Second)
		data <- datas
	}

	for _, i := range dataset {
		fmt.Println("send Data ", i)
		go hello("ocbc", i)
	}

	for i := 0; i < len(dataset); i++ {
		select {
		case message := <-data:
			fmt.Println("receive", message)
		}
	}
	fmt.Println("TIME ", int64(time.Since(startTime)/time.Millisecond))

}
