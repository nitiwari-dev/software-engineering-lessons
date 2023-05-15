package helper

import "fmt"

func Println(a ...interface{}) {
	fmt.Println(a)
}

func Printf(format string, a ...interface{}) string {
	return fmt.Sprintf(format, a)
}
