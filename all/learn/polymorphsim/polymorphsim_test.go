package polymorphism

import (
	"fmt"
	"testing"
)

type Code string
type Programmer interface {
	WriteHelloWorld() Code
}

type GoProgrammer struct {
}

func (p *GoProgrammer) WriteHelloWorld() Code {
	return "fmt.Println(\"Hello World!\")"
}

type JavaProgrammer struct {
}

func (p *JavaProgrammer) WriteHelloWorld() Code {
	return "System.out.Println(\"Hello World!\")"
}

// Programmer 是 Interface ,那么它对应的只能是一个指针
func writeFirstProgram(p Programmer) {
	fmt.Printf("%T %v\n", p, p.WriteHelloWorld())
}

func TestPolymorphism(t *testing.T) {
	goProg := &GoProgrammer{}
	javaProg := new(JavaProgrammer)
	writeFirstProgram(goProg)
	writeFirstProgram(javaProg)
}
