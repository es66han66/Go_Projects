package main

import "example.com/eshan/singleton"

func main() {
	s1 := singleton.GetInstance()
	s2 := singleton.GetInstance()
	println("s1.GetName()", s1.GetName())
	println("s2.GetName()", s2.GetName())
	println("s1 == s2", s1 == s2)
}
