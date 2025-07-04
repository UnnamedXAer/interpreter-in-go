package ast

import (
	"bytes"

	"github.com/unnamedxaer/interpreter-in-go/token"
)

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

func (p *Program) TokenLiteral() string {
	if (len(p.Statements)) > 0 {
		return p.Statements[0].TokenLiteral()
	}

	return ""
}

type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Idendifier
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

type Idendifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (i *Idendifier) expressionNode()      {}
func (i *Idendifier) TokenLiteral() string { return i.Token.Literal }

func (i *Idendifier) String() string {
	return i.Value
}

type ReturnStatement struct {
	Token       token.Token // the token.RETURN token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}

type ExpressonStatament struct {
	Token      token.Token // the first token of the expression
	Expression Expression
}

func (es *ExpressonStatament) statementNode()       {}
func (es *ExpressonStatament) TokenLiteral() string { return es.Token.Literal }

func (es *ExpressonStatament) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}

	return ""
}
