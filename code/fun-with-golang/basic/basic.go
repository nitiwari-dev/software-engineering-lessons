package basic

import (
<<<<<<< HEAD
<<<<<<< HEAD
	"errors"
	"fun-with-golang/helper"
<<<<<<< HEAD
	"math/rand"
	"reflect"
	"unicode/utf8"
)

func Init() {
	helper.Println("Starting easy section...")
	helloWorld()
	variables()
	iterations()
	storageDataStructure()
	functions()
	closures()
	recursion()
	pointers()
	stringAndUTF8()
	helper.Println("Ending easy section...")
}

/**
Ref: https://gobyexample.com/strings-and-runes

In Go, the concept of a character is called a rune -
it’s an integer that represents a Unicode code point.
This Go blog post is a good introduction to the topic.
https://go.dev/blog/strings
*/
func stringAndUTF8() {
	helper.Println("Starting String with UTF-8")
	const name = "foo"
	helper.Println(utf8.RuneCountInString(name))
	for index, runeValue := range name {
		helper.Println(index, runeValue)
	}

	//
	symbolic := "日本語"
	helper.Println(len(symbolic))                    // We cannot rely on len as its not equal to no. of characters. Its equal to no of bytes altogether in the  UTF-8
	helper.Println(utf8.RuneCountInString(symbolic)) // Prints 3

	for index, character := range symbolic { //converting to string
		helper.Println(index, string(character))
	}

	runeArray := []rune(symbolic)             // using unicode point. its Like creating array of unicode
	for index, character := range runeArray { //decode each rune value
		helper.Println(index, string(character))
	}
}

func pointers() {
	helper.Println("Starting pointers")
	i := 10
	helper.Println("Address of i", &i)
	passByValue(i)
	helper.Println(i)

	passByPointer()
	helper.Println(i) // This will be printed as i = 0. Because we de-referenced it within the function
}

/**
* refere to value to which the pointer references eg.*i
& address to the variable

*/
func passByPointer() {
	var count = 4
	helper.Println("count=", count)
	var countAddr = &count //new variable to the address of count
	helper.Println("count", &count)
	helper.Println("countAddr", &countAddr)
	helper.Println("count pointer value=", *countAddr)
	*countAddr = 10                                // change the pointer value to 10
	helper.Println("count change pointer=", count) // value will be changed to 10
	helper.Println("count change pointer=", countAddr)

	//struct modification
	userStructModification()
}

func userStructModification() {
	helper.Println("Start Pointer with structs")
	user := User{name: "Foo", id: 10}
	helper.Println(user)
	modifyUserStruct(&user)
	helper.Println(user)

	//Create user with new

	newUser := *new(User)
	helper.Println(newUser)
	helper.Println(reflect.TypeOf(newUser)) // pointer
}

func modifyUserStruct(user *User) {
	user.name = "Bar"
	user.id = 20
}

type User struct {
	name string
	id   int64
}

func passByValue(i int) {
	i = 0 //changing the value of the variable
}

func recursion() {
	var findNext func(interval int) int // using closure with recursion
	findNext = func(interval int) int {
		if interval == 0 {
			return 0
		} else {
			return interval + findNext(interval-1)
		}
	}

	helper.Println(findNext(4))
}

func closures() {
	helper.Println("Starting closures")

	interval := 10
	nextInt := nextRandomInt(interval)
	helper.Println(nextInt())
	helper.Println(nextInt())
	helper.Println(nextInt())
	helper.Println(nextInt())
}
func nextRandomInt(interval int) func() int {
	return func() int {
		return rand.Intn(interval)
	}
}

func functions() {
	helper.Println("Starting functions")
	id := "124"
	httpCode, err := loginUser(id)
	if err != nil {
		helper.Println(httpCode, err.Error())
	} else {
		helper.Println(httpCode, "")
	}

	idsVariadic(1, 2, 3, 4)           // variadic function
	idsVariadic([]int{1, 2, 3, 4}...) // variadic with slice add ... to mark it as variadic
}

func idsVariadic(ids ...int) {
	for _, id := range ids {
		helper.Println(id)
	}
}

func loginUser(id string) (int, error) {
	if len(id) == 0 {
		return 400, errors.New("invalid input: `id` cannot be blank")
	}
	return 200, nil

}

func storageDataStructure() {
	//Arrays
	var array [5]int
	helper.Println(array)

	arrayValues := [5]int{1, 2, 3}
	helper.Println(arrayValues)
	helper.Println(arrayValues[2])
	arrayValues[2] = arrayValues[2] * 2
	helper.Println(arrayValues[2])
	helper.Println(len(arrayValues))

	//Slice ~ to Advanced form of Array
	countryList := make([]string, 1)
	countryList[0] = "in"
	helper.Println(countryList)
	countryList = append(countryList, "eg", "ph", "sa", "za")
	helper.Println(countryList)
	helper.Println(countryList[2:]) // from 2 to end
	helper.Println(countryList[:2]) // from 0 to 1 (end - 1)

	//Map - KV pairg
	mapCountryByCode := make(map[string]string) // declare empty map
	mapCountryByCode["eg"] = "Egypt"
	mapCountryByCode["za"] = "South Africa"
	helper.Println(mapCountryByCode) // Print the map

	helper.Println(mapCountryByCode["eg"])              // Access by key
	helper.Println(mapCountryByCode["in"])              // Invalid key access return empty
	_, isINCountryCodePresent := mapCountryByCode["in"] // Invalid key access return boolean viz false
	helper.Println(isINCountryCodePresent)

	helper.Println(len(mapCountryByCode)) // Number of key value pair
	delete(mapCountryByCode, "eg")        // Delete a key
	helper.Println(mapCountryByCode)

	//Range - perform operation on arrays and slices
	ages := []int{10, 25, 30, 21} // over the array
	sum := 0
	for _, age := range ages {
		sum += age
	}
	helper.Println(sum / len(ages))

	for countryCode, countryName := range mapCountryByCode { // over the map
		helper.Println(countryCode, countryName)
	}

	for index, character := range "Hello" { // over string
		helper.Println(index, character)
	}

=======
=======
	"errors"
>>>>>>> 86c14fc (Lession go | variadic function)
	"fun-with-golang/helper"
	"math"
=======
>>>>>>> 97690bd (Lession | Golang | Refactor)
	"math/rand"
	"reflect"
	"unicode/utf8"
)

func Init() {
	helloWorld()
	variables()
	iterations()
	storageDataStructure()
	functions()
	closures()
	recursion()
	pointers()
	stringAndUTF8()
}

/**
Ref: https://gobyexample.com/strings-and-runes

In Go, the concept of a character is called a rune -
it’s an integer that represents a Unicode code point.
This Go blog post is a good introduction to the topic.
https://go.dev/blog/strings
*/
func stringAndUTF8() {
	helper.Println("Starting String with UTF-8")
	const name = "foo"
	helper.Println(utf8.RuneCountInString(name))
	for index, runeValue := range name {
		helper.Println(index, runeValue)
	}
}

func pointers() {
	helper.Println("Starting pointers")
	i := 10
	helper.Println("Address of i", &i)
	passByValue(i)
	helper.Println(i)

	passByPointer()
	helper.Println(i) // This will be printed as i = 0. Because we de-referenced it within the function
}

/**
* refere to value to which the pointer references eg.*i
& address to the variable

*/
func passByPointer() {
	var count = 4
	helper.Println("count=", count)
	var countAddr = &count //new variable to the address of count
	helper.Println("count", &count)
	helper.Println("countAddr", &countAddr)
	helper.Println("count pointer value=", *countAddr)
	*countAddr = 10                                // change the pointer value to 10
	helper.Println("count change pointer=", count) // value will be changed to 10
	helper.Println("count change pointer=", countAddr)

	//struct modification
	userStructModification()
}

func userStructModification() {
	helper.Println("Start Pointer with structs")
	user := User{name: "Foo", id: 10}
	helper.Println(user)
	modifyUserStruct(&user)
	helper.Println(user)

	//Create user with new

	newUser := *new(User)
	helper.Println(newUser)
	helper.Println(reflect.TypeOf(newUser)) // pointer
}

func modifyUserStruct(user *User) {
	user.name = "Bar"
	user.id = 20
}

type User struct {
	name string
	id   int64
}

func passByValue(i int) {
	i = 0 //changing the value of the variable
}

func recursion() {
	var findNext func(interval int) int // using closure with recursion
	findNext = func(interval int) int {
		if interval == 0 {
			return 0
		} else {
			return interval + findNext(interval-1)
		}
	}

	helper.Println(findNext(4))
}

func closures() {
	helper.Println("Starting closures")

	interval := 10
	nextInt := nextRandomInt(interval)
	helper.Println(nextInt())
	helper.Println(nextInt())
	helper.Println(nextInt())
	helper.Println(nextInt())
}
func nextRandomInt(interval int) func() int {
	return func() int {
		return rand.Intn(interval)
	}
}

func functions() {
	helper.Println("Starting functions")
	id := "124"
	httpCode, err := loginUser(id)
	if err != nil {
		helper.Println(httpCode, err.Error())
	} else {
		helper.Println(httpCode, "")
	}

	idsVariadic(1, 2, 3, 4)           // variadic function
	idsVariadic([]int{1, 2, 3, 4}...) // variadic with slice add ... to mark it as variadic
}

func idsVariadic(ids ...int) {
	for _, id := range ids {
		helper.Println(id)
	}
}

func loginUser(id string) (int, error) {
	if len(id) == 0 {
		return 400, errors.New("invalid input: `id` cannot be blank")
	}
	return 200, nil

}

func storageDataStructure() {
	//Arrays
	var array [5]int
	helper.Println(array)

	arrayValues := [5]int{1, 2, 3}
	helper.Println(arrayValues)
	helper.Println(arrayValues[2])
	arrayValues[2] = arrayValues[2] * 2
	helper.Println(arrayValues[2])
	helper.Println(len(arrayValues))

	//Slice ~ to Advanced form of Array
	countryList := make([]string, 1)
	countryList[0] = "in"
	helper.Println(countryList)
	countryList = append(countryList, "eg", "ph", "sa", "za")
	helper.Println(countryList)
	helper.Println(countryList[2:]) // from 2 to end
	helper.Println(countryList[:2]) // from 0 to 1 (end - 1)

	//Map - KV pair
	mapCountryByCode := make(map[string]string) // declare empty map
	mapCountryByCode["eg"] = "Egypt"
	mapCountryByCode["za"] = "South Africa"
	helper.Println(mapCountryByCode) // Print the map

	helper.Println(mapCountryByCode["eg"])              // Access by key
	helper.Println(mapCountryByCode["in"])              // Invalid key access return empty
	_, isINCountryCodePresent := mapCountryByCode["in"] // Invalid key access return boolean viz false
	helper.Println(isINCountryCodePresent)

	helper.Println(len(mapCountryByCode)) // Number of key value pair
	delete(mapCountryByCode, "eg")        // Delete a key
	helper.Println(mapCountryByCode)

	//Range - perform operation on arrays and slices
	ages := []int{10, 25, 30, 21} // over the array
	sum := 0
	for _, age := range ages {
		sum += age
	}
	helper.Println(sum / len(ages))

	for countryCode, countryName := range mapCountryByCode { // over the map
		helper.Println(countryCode, countryName)
	}

	for index, character := range "Hello" { // over string
		helper.Println(index, character)
	}

>>>>>>> 80df7ad (fun with go)
}

func iterations() {

	//Loop
	for i := 0; i <= 5; i++ {
		if factor := 2; i%factor == 0 { // If else with declaration and init
			helper.Println(i, " is even")
		} else {
			helper.Println(i, " is Odd")
		}
	}

	//switch
	alphabetCount := 'z' - 'a'
	helper.Println(alphabetCount)
	switch alphabetCount {
	case 25:
		helper.Println("Its z")
	default:
		helper.Println("unknown")
	}

}

<<<<<<< HEAD
<<<<<<< HEAD
=======
const PI = math.Pi

>>>>>>> 80df7ad (fun with go)
=======
>>>>>>> 97690bd (Lession | Golang | Refactor)
func variables() {

	const pi = 3.1415
	var isTaxable = true // Declare + Init

	age := 30 // Infer without declaration

	var taxRate float32 // Declare
	taxRate = 35.5      // Init

	city := "New York"

	helper.Println(isTaxable) // Boolean
	helper.Println(age)       // Integer
	helper.Println(taxRate)   // Floats
	helper.Println(city)      // String
	helper.Println(pi)        // Constant local scope is of higher precedence
}
func helloWorld() {
	helper.Println("Hello world")
}
