package main

type A interface {
	Foo() string
}

type B interface {
	Foo() string
	Bar() string
}

func takeA(a A) {
	a.Foo()
}

type C struct{}

func NewC() B {
	return &C{}
}

func (c *C) Foo() string {
	return ""
}

func (c *C) Bar() string {
	return ""
}

func main() {
	c := NewC()
	takeA(c)
}
