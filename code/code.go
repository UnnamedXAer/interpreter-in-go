package code

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type Instructions []byte

func (ins Instructions) String() string {

	var out bytes.Buffer

	i := 0
	for i < len(ins) {
		def, err := Lookup(ins[i])
		if err != nil {
			fmt.Fprintf(&out, "ERROR: %s\n", err)
			continue
		}

		operands, read := ReadOperands(def, ins[i+1:])

		fmt.Fprintf(&out, "%04d %s\n", i, ins.fmtInstruction(def, operands))

		i += 1 + read
	}

	return out.String()
}

func (ins Instructions) fmtInstruction(def *Definition, operands []int) any {
	operandCount := len(def.OperandWidths)

	if len(operands) != operandCount {
		return fmt.Sprintf("ERROR: operand len %d does not match defined %d\n", len(operands), operandCount)
	}

	switch operandCount {
	case 0:
		return def.Name
	case 1:
		return fmt.Sprintf("%s %d", def.Name, operands[0])
	case 2:
		return fmt.Sprintf("%s %d %d", def.Name, operands[0], operands[1])

	default:
		return fmt.Sprintf("ERROR: unhandled operandCount for %s\n", def.Name)
	}
}

type Opcode byte

const (
	OpConstant       Opcode = iota // accepts a constant (e.g. 2)
	OpAdd                          // a+b
	OpPop                          //
	OpSub                          // a-b
	OpMul                          // a*b
	OpDiv                          // a/b
	OpTrue                         //
	OpFalse                        //
	OpEqual                        // a==b
	OpNotEqual                     // a!=b
	OpGreaterThan                  // a>b
	OpMinus                        // -a
	OpBang                         // !true
	OpJumpNotTruthy                // this instruction will tell the VM to only jump if the value on top of the stack is not 'truthy' (i.e. not `false` nor `null`), accepts offset
	OpJump                         // accepts offset
	OpNull                         // 'null'
	OpGetGlobal                    // retrieves value of a global variable
	OpSetGlobal                    // sets value of a global variable
	OpArray                        // accepts size of the array
	OpHash                         // accept number of keys + number of values
	OpIndex                        // [1,2][0], {a:b}[a]
	OpCall                         // calls function myFunc()
	OpReturnValue                  //
	OpReturn                       // just return (go back), no value
	OpGetLocal                     // retrieves value of a local variable
	OpSetLocal                     // sets value of a local variable
	OpGetBuiltin                   // get built-in function: len, push...
	OpClosure                      // tell the vm to wrap the specified "compiled function" int an "closure"
	OpGetFree                      // get "free variable", accepts index of closure's `Free` field
	OpCurrentClosure               //
)

type Definition struct {
	Name          string
	OperandWidths []int
}

var definitions = map[Opcode]*Definition{
	OpConstant:       {"OpConstant", []int{2}},
	OpAdd:            {"OpAdd", []int{}},
	OpPop:            {"OpPop", []int{}},
	OpSub:            {"OpSub", []int{}},
	OpMul:            {"OpMul", []int{}},
	OpDiv:            {"OpDiv", []int{}},
	OpTrue:           {"OpTrue", []int{}},
	OpFalse:          {"OpFalse", []int{}},
	OpEqual:          {"OpEqual", []int{}},
	OpNotEqual:       {"OpNotEqual", []int{}},
	OpGreaterThan:    {"OpGreaterThan", []int{}},
	OpMinus:          {"OpMinus", []int{}},
	OpBang:           {"OpBang", []int{}},
	OpJump:           {"OpJump", []int{2}},
	OpJumpNotTruthy:  {"OpJumpNotTruthy", []int{2}},
	OpNull:           {"OpNull", []int{}},
	OpGetGlobal:      {"OpGetGlobal", []int{2}}, // up to 65535 global variables (two bytes of data)
	OpSetGlobal:      {"OpSetGlobal", []int{2}},
	OpArray:          {"OpArray", []int{2}},
	OpHash:           {"OpHash", []int{2}},
	OpIndex:          {"OpIndex", []int{}},
	OpCall:           {"OpIndex", []int{1}}, // accepts number of arguments (up to 256)
	OpReturnValue:    {"OpReturnValue", []int{}},
	OpReturn:         {"OpReturn", []int{}},
	OpGetLocal:       {"OpGetLocal", []int{1}}, // up to 256 local variables
	OpSetLocal:       {"OpSetLocal", []int{1}},
	OpGetBuiltin:     {"OpGetBuiltin", []int{1}}, // index of built-in, up to 256 built-in functions;
	OpClosure:        {"OpClosure", []int{2, 1}}, // accepts: 1. `constant index` - specifies where is the constant pool we can find the `compiled function` to be converted into a closure. 2. how many `free variables` sit on the stack and need to be transferred to the about-to-be-created closure.
	OpGetFree:        {"OpGetFree", []int{1}},
	OpCurrentClosure: {"OpCurrentClosure", []int{}},
}

func Lookup(op byte) (*Definition, error) {
	def, ok := definitions[Opcode(op)]
	if !ok {
		return nil, fmt.Errorf("opcode %d undefined", op)
	}

	return def, nil
}

func Make(op Opcode, operands ...int) []byte {
	def, ok := definitions[op]
	if !ok {
		return []byte{}
	}

	instructionLen := 1

	for _, w := range def.OperandWidths {
		instructionLen += w
	}

	instruction := make([]byte, instructionLen)
	instruction[0] = byte(op)

	offset := 1
	for i, o := range operands {
		width := def.OperandWidths[i]
		switch width {
		case 2:
			binary.BigEndian.PutUint16(instruction[offset:], uint16(o))

		case 1:
			instruction[offset] = byte(o)
		}

		offset += width
	}

	return instruction
}

func ReadOperands(def *Definition, ins Instructions) ([]int, int) {
	operands := make([]int, len(def.OperandWidths))
	offset := 0

	for i, width := range def.OperandWidths {
		switch width {
		case 2:
			operands[i] = int(ReadUint16(ins[offset:]))

		case 1:
			operands[i] = int(ReadUint8(ins[offset:]))
		}

		offset += width

	}

	return operands, offset
}

func ReadUint8(ins Instructions) uint8 {
	return uint8(ins[0])
}

func ReadUint16(ins Instructions) uint16 {
	return binary.BigEndian.Uint16(ins)
}
