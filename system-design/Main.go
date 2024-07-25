package main

import (
	"system_design/hash_func"
	"system_design/redirect_router"
)

func main() {
	println(hash_func.HashFunction("hey folks"))
	redirect_router.ListenAndServe()
}
