package main

import (
	"fmt"
	"sync"
	"sync/atomic"
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

func demoAtomicCounter() {
	var ops uint64

	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			for j := 1; j <= 1000; j++ {
				atomic.AddUint64(&ops, 1)
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("ops:", ops)
}

func demoChannel11() {
	/*
		pipe := make(chan string)
		pipe <- "water"
	*/
	//fatal error: all goroutines are asleep - deadlock!

	/*
		pipe := make(chan string)
		pipe <- "water"
		receiver := <-pipe
		fmt.Println(receiver)
	*/
	//fatal error: all goroutines are asleep - deadlock!

	/*
		pipe := make(chan string)
		go func() {
			pipe <- "water"
		}()

		receiver := <-pipe
		fmt.Println(receiver)
	*/
	/*
		pipe := make(chan string, 1)
		pipe <- "water"
		receiver := <-pipe
		fmt.Println(receiver)
	*/
	/*
		pipe := make(chan string, 1)
		pipe <- "water"
		pipe <- "water"
		receiver := <-pipe
		fmt.Println(receiver)
	*/
	//fatal error: all goroutines are asleep - deadlock!

	/*
		pipe := make(chan string, 2)
		pipe <- "water"
		pipe <- "water"
		receiver := <-pipe
		fmt.Println(receiver)
		receiver = <-pipe
		fmt.Println(receiver)
	*/
	/*
		pipe := make(chan string, 2)
		pipe <- "water"
		pipe <- "water"
		for receiver := range pipe {
			fmt.Println(receiver)
		}
	*/
	/*
		pipe := make(chan string, 2)
		pipe <- "water"
		pipe <- "water"
		close(pipe)
		for receiver := range pipe {
			fmt.Println(receiver)
		}
	*/
	pipe := make(chan string, 2)
	go func() {
		pipe <- "Water"
		pipe <- "water"
		close(pipe)
	}()

	for receiver := range pipe {
		fmt.Println(receiver)
	}
}

func demoChannel12() {
	/*
		pipe := make(chan string)
		go func() { //1
			for receiver := range pipe { //2
				fmt.Println(receiver) //3
			}
		}()
		pipe <- "water 1" //4
		close(pipe)
	*/

	/*
		pipe := make(chan string)
		go func() {
			for receiver := range pipe {
				fmt.Println(receiver)
			}
		}()
		pipe <- "water 1"
		close(pipe)
		time.Sleep(time.Millisecond)
	*/
	pipe := make(chan string)
	done := make(chan bool)

	go func() {
		for {
			receiver, more := <-pipe
			fmt.Println(receiver)
			if !more {
				done <- true
				time.Sleep(time.Millisecond)
				fmt.Println("Print me out")
			}
		}
	}()
	pipe <- "water"
	close(pipe)
	<-done
}

/*
type Queuable interface {
	Handle() error
}

type Worker struct {
	Name       string
	WorkerPool chan chan Queuable
	JobChannel chan Queuable
	quit       chan bool
}

var JobQueue chan Queuable

type Dispatcher struct {
	maxWorkers int
	WorkerPool chan chan Queuable
	Workers    []Worker
}

func NewDispather(maxWorkers int) *Dispatcher {
	if JobQueue == nil {
		JobQueue = make(chan Queuable, 10)
	}

	pool := make(chan chan Queuable, maxWorkers)
	return &Dispatcher{WorkerPool: pool, maxWorkers: maxWorkers}
}

func (d *Dispatcher) Run() {
	for i := o; i < d.maxWorkers; i++ {
		worker := NewWorker(d.WorkerPool)
		worker.Start()
		d.Workers = append(d.Workers, worker)
	}
}

func (d *Dispatcher) dispatch() {
	for {
		select {
		case job := <-JobQueue:
			go func(job Queuable) {
				jobChannel := <-d.WorkerPool
				jobChannel <- job
			}(job)
		}
	}
}

func (W Worker) Start() {
	go func() {
		for {
			select {
			case job := <-W.JobChannel:
				if err := job.Handle(); err != nil {
					fmt.Printf("Error in job: %s\n", err.Error())
				}
			}
		}
	}()
}

type Email struct {
	To      string `json:"to"`
	From    string `json:"from"`
	Subject string `json:"subject"`
	Content string `json:"content"`
}

func (e Email) Handle() error {
	r := rand.Intn(200)
	time.Sleep(time.Duration(r) * time.Millisecond)
	return nil
}

type EmailService struct {
	Queue chan queue.Queuable
}

func NewEmailService(q chan queue.Queuable) *EmailService {
	service := &EmailService{
		Queue: q,
	}
	return service
}

func (s EmailService) Send(e Email) {
	s.Queue <- e
}
*/

func demoChannel13() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)
	limiter := time.Tick(200 * time.Millisecond)

	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(200 * time.Microsecond) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
