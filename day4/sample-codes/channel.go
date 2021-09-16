package main

import (
	"fmt"
	"sync"
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

func demoChannel4() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	func(pings chan<- string, msg string) {
		pings <- msg
	}(pings, "passed message")

	func(pings <-chan string, pongs chan<- string) {
		msg := <-pings
		pongs <- msg
	}(pings, pongs)
	fmt.Println(<-pongs)
}

func demoChannel5() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 1; i <= 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}

func demoChannel6() {
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second):
		fmt.Println("c1 timeout")
	}

	go func() {
		time.Sleep(1 * time.Second)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("c2 timeout")
	}
}

func demoChannel7() {
	ticker := time.NewTicker(500 * time.Millisecond)

	go func() {
		for {
			fmt.Println("Tick at ", <-ticker.C)
		}
	}()
	select {}
}

func demoChannel8() {
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case res := <-ticker.C:
				fmt.Println("ticker at ", res)
			}
		}
	}()
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("ticker stopped")
}

func worker(id int, jobs <-chan int, result chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job\n", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "stopped job\n", j)
		result <- j * 2
	}
}

func demoChannel9() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)
	for a := 1; a <= numJobs; a++ {
		<-results
	}
}

func worker3(id int) {
	fmt.Printf("Worker %d starting id\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func demoChannel10() {
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			worker3(i)
		}(i)
	}
	wg.Wait()
}
