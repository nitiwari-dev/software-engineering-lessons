package medium

import (
	"errors"
	"fmt"
	"fun-with-golang/helper"
	"math"
	"strconv"
)

func Init() {
	helper.Println("Starting medium section...")
	structs()
	deferFunction()
	methodsInStructs()
	interfaces()
	generics()
	typAssertions()
	typeStringerExample()
	helper.Println("Ending medium section...")
}

/*
*
Relate to toString() in java or kotlin
*/
func typeStringerExample() {
	user := User{name: "NT", id: 123}
	helper.Println(user)

	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

type IPAddr [4]byte

func (ip IPAddr) String() string {
	var output = ""
	var isFirst = true
	for _, value := range ip {
		if isFirst {
			output += strconv.Itoa(int(value))
			isFirst = false
		} else {
			output += "." + strconv.Itoa(int(value))
		}
	}
	return output
}

func (user User) String() string {
	return fmt.Sprintf("custom implemention for user name= %v and id=%v", user.name, user.id)
}

func typAssertions() {
	helper.Println("Type Assertions")
	var i interface{} = "Hello"
	helper.Println(i.(string))
	//helper.Println(i.(int64)) // trigger panic as it's not of correct type
	a, ok := i.(int64)
	helper.Println(a, ok)

	typeAssertionSwitch("123")
	typeAssertionSwitch(1)
}

func typeAssertionSwitch(i interface{}) {
	switch v := i.(type) {
	case string:
		helper.Println(v, "is string type length is", len(v))
	case int:
		helper.Println(v, "is int type")
	}
}

/*
*
defer
1. its defer the call and will not excecute till the calls after that is execute
2. defer stores the call in stack and pop in LIFO manner
3. Notice when the execution are defered it goes in stack
*/
func deferFunction() {

	helper.Println("starting defer example")
	for i := 0; i < 10; i++ {
		defer helper.Println("Defer", i)
	}

	helper.Println("completed defer example")

}

func generics() { // also called as type parameter

}

func interfaces() {
	i := arithmetic{input: -11}
	callInterface(i)
}

func callInterface(eval evaluate) {
	helper.Println(eval.abs())
}

type evaluate interface {
	abs() float64
}

type arithmetic struct {
	input float64
}

func (c arithmetic) abs() float64 {
	return math.Abs(c.input)
}

func methodsInStructs() {
	c := calculation{input: 10}
	helper.Println(c.square())
	helper.Println(c.cube())
	helper.Println(c.squareRoot())

	u := User{
		id:      1,
		name:    "foobar",
		usrType: "free",
		address: []Address{
			{pincode: "400000"},
			{pincode: "400001"},
		},
		alternateNumbers: AlternateNumbers{
			[]string{"1234567890", "9876543210"},
		},
	}

	helper.Println(u)
	helper.Println(u.address)
	helper.Println(len(u.address) > 0)
	helper.Println(u.address[0].validateAddress()) // access the embeded function
	helper.Println(u.alternateNumbers.validateAlternateNumbers())
}

type calculation struct {
	input int64
}

func (c calculation) square() int64 {
	return c.input * c.input
}

func (c calculation) cube() int64 {
	return c.input * c.input * c.input
}

func (c calculation) squareRoot() float64 {
	return math.Sqrt(float64(c.input))
}

func structs() {
	helper.Println("  Staring structs...")
	user1 := User{id: 10, name: "Foo"}
	helper.Println(user1)

	user2 := newUser(11, "Bar")
	helper.Println(user2)

	user2.id = 12
	helper.Println(user2.id)
	newId := &user2.id
	*newId = 15
	helper.Println(*newId)
}

func newUser(id int, name string) User { // we can also return *. Its better to keep it immutable
	return User{id: id, name: name}
}

type User struct {
	id               int
	name             string
	usrType          string
	address          []Address
	alternateNumbers AlternateNumbers
}

type AlternateNumbers struct {
	numbers []string
}

type Address struct {
	pincode string
}

func (address Address) validateAddress() bool {
	if len(address.pincode) <= 0 {
		return false
	}
	return true
}

func (numbers AlternateNumbers) validateAlternateNumbers() (bool, error) {
	if len(numbers.numbers) == 0 {
		return false, errors.New("alternate number cannot be zero")
	}
	return true, nil
}
