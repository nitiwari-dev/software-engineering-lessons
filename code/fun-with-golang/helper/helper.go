package helper

import "fmt"

func Println(a ...interface{}) {
	fmt.Println(a)
}

func Printf(format string, a ...interface{}) (n int, err error) {
	return fmt.Printf(format, a)
}
