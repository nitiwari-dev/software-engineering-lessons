package hard

import (
	"encoding/json"
	"fun-with-golang/helper"
	"os"
)

func Init() {
	helper.Println("Starting Hard section...")
	jsonMarshalling()
	helper.Println("Ending Hard section...")
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

type user struct {
	Id      int64     `json:"id"`
	Name    string    `json:"name"`
	Address []address `json:"address"`
}

type address struct {
	PinCode string `json:"pin_code"`
}
