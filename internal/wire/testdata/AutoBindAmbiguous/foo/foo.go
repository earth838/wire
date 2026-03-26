package foo

// S1 implements I
type S1 struct {
	Foo string
}

func (S1) Do() {}

// S2 implements I
type S2 struct {
	Bar int
}

func (S2) Do() {}

// I is the interface we want to bind to
type I interface {
	Do()
}

func provideS1() S1 {
	return S1{Foo: "One"}
}

func provideS2() S2 {
	return S2{Bar: 2}
}
