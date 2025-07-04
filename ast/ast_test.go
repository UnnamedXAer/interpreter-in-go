package ast

import (
	"testing"

	"github.com/unnamedxaer/interpreter-in-go/token"
)

func TestString(t testing.T) {

	// let myVar = anotherVar;

	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET},
				Name: &Idendifier{
					Token:   token.IDENT,
					Literal: "let",
				}},
		},
	}
}
