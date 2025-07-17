package vm

import (
	"github.com/unnamedxaer/interpreter-in-go/code"
	"github.com/unnamedxaer/interpreter-in-go/object"
)

type Frame struct {
	fn          *object.CompiledFunction
	ip          int // instruction pointer in THIS frame for this function;
	basePointer int // the stack pointer's value before we execute a function. It pointes to the bottom of the stack of the current call frame;
}

func NewFrame(fn *object.CompiledFunction, basePointer int) *Frame {
	return &Frame{
		fn:          fn,
		ip:          -1,
		basePointer: basePointer,
	}
}

func (f *Frame) Instructions() code.Instructions {
	return f.fn.Instructions
}
