package parser

import (
	"fmt"

	"github.com/ktny/monkey/ast"
	"github.com/ktny/monkey/lexer"
	"github.com/ktny/monkey/token"
)

// 演算子の優先順位
const (
	_ int = iota
	LOWEST
	EQUALS      // ==
	LESSGREATER // >, <
	SUM         // +
	PRODUCT     // *
	PREFIX      // -X, !X
	CALL        // myFunction(X)
)

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)

type Parser struct {
	l         *lexer.Lexer
	errors    []string
	curToken  token.Token
	peekToken token.Token

	// トークンに対応した前置、中置構文解析関数のマップ
	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
}

// New Parserインスタンスを返す
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}

	p.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	p.registerPrefix(token.IDENT, p.parseIdentifier)

	// 1回呼び出しただけではcurTokenに何も入らないため2回呼び出す
	p.nextToken()
	p.nextToken()
	return p
}

// Errors Parserインスタンスが持つエラーを返す
func (p *Parser) Errors() []string {
	return p.errors
}

// lexerで解析して取り出したTokenをcurToken, peekTokenに入れる
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// curTokenが指定のTokenTypeか否かを返す
func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

// peekTokenが指定のTokenTypeか否かを返す
func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

// peekTokenが指定のTokenでない場合はParserインスタンスのErrorに追加する
func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instaed", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

// peekTokenが期待するものか否かを返す。期待するものであればトークンをすすめる
func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

// 指定のTokenTypeに指定の前置構文解析関数を紐付ける
func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

// 指定のTokenTypeに指定の中置構文解析関数を紐付ける
func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

// ParseProgram lexerを通したTokenをStatementの集まりにして返す
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for !p.curTokenIs(token.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

// 構文解析して文を返す
func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
	}
}

// 構文解析してlet文を返す
func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// TODO: 式のパースに置き換える。今はセミコロンまで読み飛ばしている
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// 構文解析してreturn文を返す
func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}

	p.nextToken()

	// TODO: 式のパースに置き換える。今はセミコロンまで読み飛ばしている
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// 構文解析して式文を返す
func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.curToken}

	stmt.Expression = p.parseExpression(LOWEST)

	// セミコロンがあればトークンを進める。なくてもセミコロン省略可能になるだけなので問題ない
	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// 構文解析して式を返す。現在のTokenTypeに対応する前置構文解析関数を返す
func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParseFns[p.curToken.Type]
	if prefix == nil {
		return nil
	}
	leftExp := prefix()

	return leftExp
}

// 構文解析して式（識別子）を返す
func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
}
