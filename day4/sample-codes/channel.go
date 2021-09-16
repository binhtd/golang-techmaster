package main

import (
	"fmt"
	"time"
)

func demoChannel1() {
	messages := make(chan string)

	go func() {
		messages <- "Hello World"
	}()
	fmt.Println("channel is the way to in/out data from goroutine", <-messages)
}

func demoChannel2() {
	messages := make(chan string, 2)
	messages <- "buffered"
	messages <- "channel"
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}

func demoWorker1(worker1 chan bool) {
	for {

		fmt.Println("worker 1 is working")
		time.Sleep(time.Second)
		fmt.Println("worker 1 has done")
		worker1 <- true
		if <-worker1 {
			break
		}
	}
}

func demoWorker2(worker1 chan bool, worker2 chan bool) {
	for {
		if <-worker1 {
			fmt.Println("worker 2 is working")
			//time.Sleep(time.Second)
			fmt.Println("worker 2 has done")
			worker2 <- true
			break
		}
	}
}

func demoChannel3() {

	worker1 := make(chan bool, 1)
	worker2 := make(chan bool, 1)

	fmt.Println("you are calling me")
	go func() {
		for {
			fmt.Println("worker 1 is working")
			//time.Sleep(time.Second)
			fmt.Println("worker 1 has done")
			worker1 <- true
			if <-worker1 {
				break
			}
		}
	}()

	go func() {
		for {
			if <-worker1 {
				fmt.Println("worker 2 is working")
				//time.Sleep(time.Second)
				fmt.Println("worker 2 has done")
				worker2 <- true
				break
			}
		}

	}()

	<-worker2
	fmt.Println("both worker 1, worker 2 have done the work")

}
