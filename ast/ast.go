package ast

import (
	"bytes"

	"github.com/ktny/monkey/token"
)

// Node ASTの最も基本的な単位。Token1つであり、TokenLiteralを持つ
type Node interface {
	TokenLiteral() string
	String() string
}

// Statement 文。Nodeの集まり
type Statement interface {
	Node
	statementNode()
}

// Expression 式。Nodeの集まり
type Expression interface {
	Node
	expressionNode()
}

// Identifier 識別子。Token1つからなり、Valueを持つ。変数
// Expression I/F
// 	expressionNode()
// Node I/F
// 	TokenLiteral()
// 	String()
type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string       { return i.Value }

// IntegerLiteral 数値リテラル。Token1つからなり、Valueを持つ
// Expression I/F
// 	expressionNode()
// Node I/F
// 	TokenLiteral()
// 	String()
type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode()      {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string       { return il.Token.Literal }

// LetStatement let文。
// Statement I/F
// 	statementNode()
// Node I/F
// 	TokenLiteral()
// 	String()
type LetStatement struct {
	Token token.Token // letトークン
	Name  *Identifier
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

// ReturnStatement return文。
// Statement I/F
// 	statementNode()
// Node I/F
// 	TokenLiteral()
// 	String()
type ReturnStatement struct {
	Token       token.Token // returnトークン
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

// ExpressionStatement 式文。
// Statement I/F
// 	statementNode()
// Node I/F
// 	TokenLiteral()
// 	String()
type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

// Program プログラム。複数の式からなる
// Node I/F
// 	TokenLiteral()
// 	String()
type Program struct {
	Statements []Statement
}

// TokenLiteral 先頭の文のTokenLiteralを返す
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// Programが持つ文すべてのStringを返す
func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}
