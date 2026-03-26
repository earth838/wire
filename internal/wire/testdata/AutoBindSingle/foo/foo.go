package main

import "fmt"

// S1 implements I
type S1 struct {
	Foo string
}

func (S1) Do() {}

// I is the interface we want to bind to
type I interface {
	Do()
}

func provideS1() S1 {
	return S1{Foo: "Hello Autowire"}
}

func main() {
	fmt.Println(injectedMessage())
}
