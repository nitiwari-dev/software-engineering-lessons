package medium

import "fun-with-golang/helper"

func Init() {
	structs()
	methodsInStructs()
	interfaces()
}

func interfaces() {

}

type maths interface {
	calculate() int64
}

func (c calculation) calculate() {

}

func methodsInStructs() {
	c := calculation{input: 10}
	helper.Println(c.square())
	helper.Println(c.cube())
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

func structs() {
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
	id   int
	name string
}
