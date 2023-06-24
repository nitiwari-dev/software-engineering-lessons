package hard

import (
	"encoding/json"
	"fmt"
	"fun-with-golang/helper"
	"os"
	"time"
)

func Init() {
	helper.Println("Starting Hard section...")
	jsonMarshalling()
	timeExample()
	goroutines()
	channels()
	helper.Println("Ending Hard section...")
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
