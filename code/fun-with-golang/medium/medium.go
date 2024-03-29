package medium

import (
	"errors"
	"fmt"
	"fun-with-golang/helper"
	"math"
	"sort"
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
	sorting()
	recoverExample()
	helper.Println("Ending medium section...")
}

/*
*
1. Recovers from the panic
2. Must be called within the defer function
*/
func recoverExample() {
	defer recoverResponse()
	panicExample()
}

func recoverResponse() {
	r := recover()
	if r != nil {
		fmt.Println("error ", r)
	}
}

func triggerPanicToRecover() {
	panic("triggr panic to be recovered")
}

type Response struct {
	status int
}

func triggerApi(request Request) (Response, error) {

	if len(request.name) == 0 {
		return Response{}, errors.New("name cannot be blank")
	}

	return Response{status: 200}, nil

}

type Request struct {
	name string
}

func panicExample() {
	_, err := triggerApi(Request{name: ""})
	if err != nil {
		panic(err)
	}

	result, err := triggerApi(Request{name: "Mum"})
	if err != nil {
		panic(err)
	} else {
		fmt.Println(result)
	}
}

type byLength []string

func (s byLength) Len() int {
	return len(s)
}

func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func sorting() {
	ints := []int{5, 2, 3}
	sort.Ints(ints)
	fmt.Println(ints)

	//status of sorting
	fmt.Println(sort.IntsAreSorted(ints))

	//sort strings using slice - ascending
	font := "123arial"
	fontRune := []rune(font)
	sort.Slice(fontRune, func(i, j int) bool {
		return font[i] < font[j] // ascending sort
	})
	fmt.Println(string(fontRune))

	//sort strings using slice - ascending
	sort.Slice(fontRune, func(i, j int) bool {
		return fontRune[i] > fontRune[j]
	})
	fmt.Println(string(fontRune))

	//custom sorting using only length for string array
	unsortedArray := []string{"bangkok", "mumbai", "pune"}
	sort.Sort(byLength(unsortedArray))
	fmt.Println(unsortedArray)

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
	//helper.Println(i.(int64)) // trigger panicExample as it's not of correct type
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
