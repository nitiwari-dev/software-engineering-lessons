package intro

import (
	"fun-with-golang/intro/basic"
	"fun-with-golang/intro/hard"
	"fun-with-golang/intro/medium"
)

// Intro Here Basic, Medium, Hard donates the order in which the material to be consumed, it doesn't donate the level of your expertise/**
func Intro() {
	basic.Init()
	medium.Init()
	hard.Init()
}
