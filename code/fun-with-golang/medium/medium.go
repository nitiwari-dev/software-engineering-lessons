package medium

import (
	"errors"
	"fun-with-golang/helper"
	"math"
	"time"
)

func Init() {
	helper.Println("Starting medium section...")
	structs()
	deferFunction()
	methodsInStructs()
	interfaces()
	generics()
	time.Now().Day()
	helper.Println("Ending medium section...")

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

	u := user{
		id:      1,
		name:    "foobar",
		usrType: "free",
		address: address{
			pincode: "",
		},
		alternateNumbers: alternateNumbers{
			[]string{"1234567890", "9876543210"},
		},
	}

	helper.Println(u)
	helper.Println(u.pincode)
	helper.Println(u.validateAddress()) // access the embeded function
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
	user1 := user{id: 10, name: "Foo"}
	helper.Println(user1)

	user2 := newUser(11, "Bar")
	helper.Println(user2)

	user2.id = 12
	helper.Println(user2.id)
	newId := &user2.id
	*newId = 15
	helper.Println(*newId)
}

func newUser(id int, name string) user { // we can also return *. Its better to keep it immutable
	return user{id: id, name: name}
}

type user struct {
	id      int
	name    string
	usrType string
	address
	alternateNumbers alternateNumbers
}

type alternateNumbers struct {
	numbers []string
}

type address struct {
	pincode string
}

func (address address) validateAddress() bool {
	if len(address.pincode) <= 0 {
		return false
	}
	return true
}

func (numbers alternateNumbers) validateAlternateNumbers() (bool, error) {
	if len(numbers.numbers) == 0 {
		return false, errors.New("alternate number cannot be zero")
	}
	return true, nil
}
