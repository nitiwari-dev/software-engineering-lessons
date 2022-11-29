package basic

import (
	"fun-with-golang/helper"
	"math"
)

func Init() {
	helloWorld()
	variables()
	iterations()

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
