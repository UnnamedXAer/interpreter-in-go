package vm

import (
	"github.com/unnamedxaer/interpreter-in-go/code"
	"github.com/unnamedxaer/interpreter-in-go/object"
)

type Frame struct {
	fn *object.CompiledFunction
	ip int //instruction pointer in THIS frame for this function;
}

func NewFrame(fn *object.CompiledFunction) *Frame {
	return &Frame{fn: fn, ip: -1}
}

func (f *Frame) Instructions() code.Instructions {
	return f.fn.Instructions
}
