package evaluator

import (
	"bytes"
	"fmt"
	"os"

	"github.com/unnamedxaer/interpreter-in-go/object"
)

var builtins = map[string]*object.Builtin{
	"len": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			switch arg := args[0].(type) {

			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}

			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}

			default:
				return newError("argument to `len` not supported, got %s", args[0].Type())
			}

		},
	},

	"first": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `first` must be ARRAY, got %s", args[0].Type())
			}

			arr := args[0].(*object.Array)
			if len(arr.Elements) > 0 {
				return arr.Elements[0]
			}

			return NULL
		},
	},

	"last": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `last` must be ARRAY, got %s", args[0].Type())
			}

			arr := args[0].(*object.Array)
			if len(arr.Elements) > 0 {
				return arr.Elements[len(arr.Elements)-1]
			}

			return NULL
		},
	},

	"rest": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `rest` must be ARRAY, got %s", args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			if length > 0 {
				newElements := make([]object.Object, length-1, length-1)
				copy(newElements, arr.Elements[1:])
				return &object.Array{Elements: newElements}
			}

			return NULL
		},
	},

	"push": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 2 {
				return newError("wrong number of arguments. got=%d, want=>=2", len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `push` must be ARRAY, got %s", args[0].Type())
			}

			arr := args[0].(*object.Array)
			incomingElements := args[1:]

			length := len(arr.Elements)
			newLenght := length + len(incomingElements)

			updatedElements := make([]object.Object, newLenght, newLenght)
			copy(updatedElements, arr.Elements)
			copy(updatedElements[length:], incomingElements)
			return &object.Array{Elements: updatedElements}

		},
	},

	"puts": {
		Fn: func(args ...object.Object) object.Object {

			var out bytes.Buffer

			if len(args) == 0 {
				return NULL
			}

			out.WriteString(args[0].Inspect())

			if len(args) > 1 {
				for i := 1; i < len(args); i++ {
					out.WriteString(" ")
					out.WriteString(args[i].Inspect())
				}
			}

			fmt.Fprintln(os.Stdout, out.String())

			return NULL
		},
	},
}
