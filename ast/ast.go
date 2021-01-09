package ast

import (
	"bytes"
	"strings"

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

type Boolean struct {
	Token token.Token
	Value bool
}

func (b *Boolean) expressionNode()      {}
func (b *Boolean) TokenLiteral() string { return b.Token.Literal }
func (b *Boolean) String() string       { return b.Token.Literal }

// PrefixExpression 前置演算子式
// Expression I/F
// 	expressionNode()
// Node I/F
// 	TokenLiteral()
// 	String()
type PrefixExpression struct {
	Token    token.Token // 演算子トークン。ex. -
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) expressionNode()      {}
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")
	return out.String()
}

// InfixExpression 中置演算子式
// Expression I/F
// 	expressionNode()
// Node I/F
// 	TokenLiteral()
// 	String()
type InfixExpression struct {
	Token    token.Token // 演算子トークン。ex. +
	Left     Expression
	Operator string
	Right    Expression
}

func (in *InfixExpression) expressionNode()      {}
func (in *InfixExpression) TokenLiteral() string { return in.Token.Literal }
func (in *InfixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(in.Left.String())
	out.WriteString(" " + in.Operator + " ")
	out.WriteString(in.Right.String())
	out.WriteString(")")
	return out.String()
}

// IfExpression if式
// Expression I/F
// 	expressionNode()
// Node I/F
// 	TokenLiteral()
// 	String()
type IfExpression struct {
	Token       token.Token // ifトークン
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (ie *IfExpression) expressionNode()      {}
func (ie *IfExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IfExpression) String() string {
	var out bytes.Buffer
	out.WriteString("if")
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())

	if ie.Alternative != nil {
		out.WriteString("else ")
		out.WriteString(ie.Alternative.String())
	}

	return out.String()
}

// FunctionLiteral 関数リテラル
// Expression I/F
// 	expressionNode()
// Node I/F
// 	TokenLiteral()
// 	String()
type FunctionLiteral struct {
	Token      token.Token // fnトークン
	Parameters []*Identifier
	Body       *BlockStatement
}

func (fl *FunctionLiteral) expressionNode()      {}
func (fl *FunctionLiteral) TokenLiteral() string { return fl.Token.Literal }
func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(fl.Body.String())
	return out.String()
}

// CallExpression 呼び出し式
// Expression I/F
// 	expressionNode()
// Node I/F
// 	TokenLiteral()
// 	String()
type CallExpression struct {
	Token     token.Token // ( トークン
	Function  Expression  // Identifier または FunctionLiteral
	Arguments []Expression
}

func (ce *CallExpression) expressionNode()      {}
func (ce *CallExpression) TokenLiteral() string { return ce.Token.Literal }
func (ce *CallExpression) String() string {
	var out bytes.Buffer

	args := []string{}
	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}

	out.WriteString(ce.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")
	return out.String()
}

// BlockStatement { から始まるブロック文
// Statement I/F
// 	statementNode()
// Node I/F
// 	TokenLiteral()
// 	String()
type BlockStatement struct {
	Token      token.Token // { トークン
	Statements []Statement
}

func (bs *BlockStatement) statementNode()       {}
func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }
func (bs *BlockStatement) String() string {
	var out bytes.Buffer
	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

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

type StringLiteral struct {
	Token token.Token
	Value string
}

func (sl *StringLiteral) expressionNode()      {}
func (sl *StringLiteral) TokenLiteral() string { return sl.Token.Literal }
func (sl *StringLiteral) String() string       { return sl.Token.Literal }

type ArrayLiteral struct {
	Token    token.Token
	Elements []Expression
}

func (al *ArrayLiteral) expressionNode()      {}
func (al *ArrayLiteral) TokenLiteral() string { return al.Token.Literal }
func (al *ArrayLiteral) String() string {
	var out bytes.Buffer

	elements := []string{}
	for _, el := range al.Elements {
		elements = append(elements, el.String())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")
	return out.String()
}

type IndexExpression struct {
	Token token.Token // '[' トークン
	Left  Expression
	Index Expression
}

func (ie *IndexExpression) expressionNode()      {}
func (ie *IndexExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IndexExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString("[")
	out.WriteString(ie.Index.String())
	out.WriteString("])")
	return out.String()
}
