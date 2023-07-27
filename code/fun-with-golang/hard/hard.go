package hard

import (
	"encoding/json"
	"fmt"
	"fun-with-golang/helper"
	"os"
	"sync"
	"time"
)

func Init() {
	helper.Println("Starting Hard section...")
	jsonMarshalling()
	timeExample()
	goroutines()
	channels()
	selectChannels()
	timeouts()
	timer()
	workers()
	waitgroup()
	rateLimiting()
	mutex()
	helper.Println("Ending Hard section...")
}

type Container struct {
	mu      sync.Mutex
	counter map[string]int
}

func (c *Container) incr(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counter[name]++
}
func mutex() {

	c := Container{
		counter: map[string]int{"a": 0, "b": 0},
	}

	var wq sync.WaitGroup

	incrementFunction := func(name string, count int) {
		for i := 0; i < count; i++ {
			c.incr(name)
		}
		wq.Done()
	}

	wq.Add(3)
	go incrementFunction("a", 1)
	go incrementFunction("b", 1)
	go incrementFunction("b", 1)
	wq.Wait()
	fmt.Println(c.counter)
}

func rateLimiting() {
	const requestCount = 5
	request := make(chan int, requestCount)
	for i := 1; i <= requestCount; i++ {
		request <- i
	}

	close(request)

	limiter := time.Tick(500 * time.Millisecond)

	for req := range request {
		<-limiter // blocker and after every 500 ms its excecuted
		fmt.Println("request of ", req, "at", time.Now())
	}

	burstLimiter := make(chan time.Time, 3) // 3 event burst
	for i := 0; i < 3; i++ {
		burstLimiter <- time.Now()
	}

	//every 200 ms lets add to burstLimiter
	go func() {
		for t := range time.Tick(500 * time.Millisecond) {
			burstLimiter <- t
		}
	}()

	//similate 5 request
	limiterReq := make(chan int, requestCount)
	for i := 1; i <= requestCount; i++ {
		limiterReq <- i
	}
	close(limiterReq)

	for r := range limiterReq {
		<-burstLimiter
		fmt.Println("request of limiter", r, time.Now())
	}

}

func waitgroup() {
	//declare
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			unitWork(i)
		}()
	}

	wg.Wait()
}

func unitWork(input int) {
	fmt.Println("Unit work started", input)
	time.Sleep(time.Second)
	fmt.Println("Unit work completed", input)
}

func workers() {
	const jobsCount = 5
	jobs := make(chan int, jobsCount)
	result := make(chan int, jobsCount)

	for workerId := 1; workerId <= 5; workerId++ {
		go asyncWork(jobs, result, workerId)
	}

	//assign jobs to worker
	for i := 1; i <= jobsCount; i++ {
		jobs <- i
	}
	close(jobs)

	for j := 1; j <= jobsCount; j++ {
		fmt.Println("result", <-result)
	}
}

func asyncWork(job <-chan int, result chan<- int, workerId int) {
	for j := range job {
		fmt.Println("worker start", workerId, j)
		//time.Sleep(100 * time.Millisecond)
		fmt.Println("worker end ", workerId, j, "done")
		result <- j * 2
	}
}

func timer() {
	cronTimer := time.NewTimer(2 * time.Second)
	<-cronTimer.C // Block and then excecute
	fmt.Println("execute after 2 seconds")

	//using ticker
	ticker := time.NewTicker(500 * time.Millisecond) // ticker with 500mx
	status := make(chan bool, 1)                     // status to stop
	go func() {
		for {
			select {
			case done := <-status:
				fmt.Println("Timer status: ", done)
				return
			case t := <-ticker.C:
				fmt.Println("Timer ticking", t)
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	status <- true
	fmt.Println("Timer stoppped")
}

func timeouts() {
	in := make(chan string, 1)

	go func() {
		time.Sleep(3 * time.Second)
		in <- "incoming"
	}()

	//timeout select - timeout while accessing the <-in channel
	select {
	case message := <-in:
		helper.Println("incoming", message)
	case <-time.After(2 * time.Second):
		helper.Println("incoming timeout")
	}

	//NOTE: We can use default case to apply non-blocking channel operations
}

/**
* select multiple goroutines with blocking operation
 */
func selectChannels() {
	in := make(chan string, 1)
	out := make(chan string, 1)

	go func() {
		time.Sleep(1 * time.Second)
		in <- "in"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		out <- "out"
	}()

	for i := 0; i < 2; i++ {
		select {
		case inMessage := <-in:
			fmt.Println("Incoming message", inMessage)
		case outMessage := <-out:
			fmt.Println("Outgoing message", outMessage)
		}
	}
}

/*
program to connect concurrent goroutines
*/
func channels() {

	//channel
	subscribe := make(chan string)
	go func() { subscribe <- "subscribed" }()
	syncCall("sync")

	message := <-subscribe
	fmt.Println(message)

	//Channel buffering with only two values
	serverStatus := make(chan int, 2)
	serverStatus <- 0
	serverStatus <- 1
	fmt.Println(<-serverStatus) // will print 0
	fmt.Println(<-serverStatus) // will print 1

	//channel sync
	completed := make(chan bool, 1)
	go syncCallChannel(completed)

	<-completed // wait till the goroutines returns status as done

	//direction incoming and outgoing
	in := make(chan string, 1)
	out := make(chan string, 1)

	incoming("incoming and outgoing", in)
	outgoing(in, out)
	helper.Println(<-out)

	//closing channel
	close(in)
	close(out)

	//range over the channel
	statusRange := make(chan string, 3)
	statusRange <- "Todo"
	statusRange <- "In Progress"
	statusRange <- "Done"
	close(statusRange)

	for status := range statusRange {
		helper.Println(status)
	}

}

func outgoing(in chan string, out chan string) {
	msg := <-in
	out <- msg
}

func incoming(s string, in chan string) {
	in <- s
}

func goroutines() {
	syncCall("sync")
	go syncCall("goroutines")

	time.Sleep(time.Second)
}

func syncCall(from string) {
	for i := 0; i < 2; i++ {
		helper.Println(from, i)
	}
}

func syncCallChannel(done chan bool) {
	for i := 0; i < 2; i++ {
		helper.Println("channel ", i)
	}

	done <- true

}

func timeExample() {
	now := time.Now()
	helper.Println(now.Date())    // [2023 June 25]
	helper.Println(now.Day())     // 25
	helper.Println(now.Month())   // June
	helper.Println(now.Weekday()) // Sunday
}

func jsonMarshalling() {
	marshallString()
	marshallUnmarshallStrut()
}

func marshallUnmarshallStrut() {
	//marshalling user properties
	userByte := &user{Id: 10, Name: "FooBar", Address: []address{{PinCode: "400001"}, {PinCode: "400002"}}}
	userMarshal, _ := json.Marshal(userByte)
	helper.Println(string(userMarshal))

	//unmarshalling user properties
	userData := user{}
	err := json.Unmarshal(userMarshal, &userData)
	if err != nil {
		helper.Println(err)
	} else {
		helper.Println(userData.Address)
	}

	//use jsonEncoder to stdout
	encoder := json.NewEncoder(os.Stdout)
	err = encoder.Encode(userData)
}

func marshallString() {
	marshallString, err := json.Marshal("Hello marshall")
	if err != nil {
		helper.Println(err.Error())
	} else {
		helper.Println(string(marshallString)) // marshalling of string type
	}

}

type address struct {
	PinCode string `json:"pin_code"`
}

type user struct {
	Id      int64     `json:"id"`
	Name    string    `json:"name"`
	Address []address `json:"address"`
}
