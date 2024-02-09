package main

import (
	"system-design/hash_func"
	"system-design/redirect_router"
)

func main() {
	println(hash_func.HashFunction("hey folks"))
	redirect_router.ListenAndServe()
}
