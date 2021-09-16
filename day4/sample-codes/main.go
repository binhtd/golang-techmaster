package main

func main() {
	//demoGoroutines1()
	//demoChannel1()
	//demoChannel2()

	//demoChannel3()

	worker1 := make(chan bool, 1)
	worker2 := make(chan bool, 1)

	go demoWorker1(worker1)
	go demoWorker2(worker1, worker2)
	<-worker2
}
