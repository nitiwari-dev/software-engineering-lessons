package basic

import (
	"errors"
	"fun-with-golang/helper"
	"math"
)

func Init() {
	helloWorld()
	variables()
	iterations()
	storageDataStructure()
	functions()
}

func functions() {
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

const PI = math.Pi

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
