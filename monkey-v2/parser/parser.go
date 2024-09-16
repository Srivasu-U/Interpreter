package parser

import (
	"Learning-Go/monkey-v2/ast"
	"Learning-Go/monkey-v2/lexer"
	"Learning-Go/monkey-v2/token"
)

/*
The Parser essentially takes in the lexer as an input and repeatedly
calls l.NextToken() to read through the tokens. We keep track of both
current and next token to see what is to be done based on both. For example,
if we end a node at '5;' then curToken is just a token.INT and peekToken
indicates the token.SEMICOLON to indicate end of line.
*/
type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// Read two tokens to set curToken and peekToken.
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
