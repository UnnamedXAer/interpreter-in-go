package vm

import (
	"github.com/unnamedxaer/interpreter-in-go/code"
	"github.com/unnamedxaer/interpreter-in-go/object"
)

type Frame struct {
	cl          *object.Closure
	ip          int // instruction pointer in THIS frame for this function;
	basePointer int // the stack pointer's value before we execute a function. It pointes to the bottom of the stack of the current call frame;
}

func NewFrame(closure *object.Closure, basePointer int) *Frame {
	return &Frame{
		cl:          closure,
		ip:          -1,
		basePointer: basePointer,
	}
}

func (f *Frame) Instructions() code.Instructions {
	return f.cl.Fn.Instructions
}
