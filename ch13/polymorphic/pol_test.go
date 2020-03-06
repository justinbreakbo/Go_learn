package polymorphic

import (
	"fmt"
	"testing"
)

// 多态的实现

type Programmer interface {
	writeHelloWorld() string
}

type GoProgrammer struct {
	name string
}

func (g GoProgrammer) writeHelloWorld() string {
	return fmt.Sprintf("hello world, go")
}

type JavaProgrammer struct {
}

func (j JavaProgrammer) writeHelloWorld() string {
	return fmt.Sprintf("hello world, java")
}

func writeFirstProgram(p Programmer) {
	fmt.Printf("%T %v \n", p, p.writeHelloWorld())
}
func TestPolymorphic(t *testing.T) {
	goProg := new(GoProgrammer)
	jaProg := new(JavaProgrammer)
	writeFirstProgram(goProg)
	writeFirstProgram(jaProg)
}
