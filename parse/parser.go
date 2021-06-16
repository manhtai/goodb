package parse

import (
	"fmt"
)

type Parser struct {
	lexer     *Lexer
	curToken  Token
	peekToken Token
	errors    []string
}

func NewParser(input string) *Parser {
	lexer := NewLexer(input)
	return &Parser{
		lexer: lexer,
	}
}

func (parser *Parser) parse() *Ast {
	ast := &Ast{}
	ast.Statements = []*Statement{}
	for !parser.curTokenIs(EOF) {
		stmt := parser.parseStatement()
		if stmt != nil {
			ast.Statements = append(ast.Statements, stmt)
		}
		parser.nextToken()
	}
	return ast
}

func (parser *Parser) parseStatement() *Statement {
	switch parser.curToken.Type {
	case SelectKeyword:
		return parser.parseSelectStatement()
	case CreateKeyword:
		switch parser.peekToken.Type {
		case TableKeyword:
			return parser.parseCreateTableStatement()
		case ViewKeyword:
			return parser.parseCreateViewStatement()
		case IndexKeyword:
			return parser.parseCreateIndexStatement()
		default:
			panic("invalid create statement")
		}
	case InsertKeyword:
		return parser.parseInsertStatement()
	case DeleteKeyword:
		return parser.parseDeleteStatement()
	case UpdateKeyword:
		return parser.parseUpdateStatement()
	default:
		return parser.parseExpressionStatement()
	}
}

func (parser *Parser) parseSelectStatement() *Statement {
	panic("implement me")
}

func (parser *Parser) parseCreateTableStatement() *Statement {
	panic("implement me")
}

func (parser *Parser) parseCreateViewStatement() *Statement {
	panic("implement me")
}

func (parser *Parser) parseCreateIndexStatement() *Statement {
	panic("implement me")
}

func (parser *Parser) parseInsertStatement() *Statement {
	panic("implement me")
}

func (parser *Parser) parseDeleteStatement() *Statement {
	panic("implement me")
}

func (parser *Parser) parseUpdateStatement() *Statement {
	panic("implement me")
}

func (parser *Parser) parseExpressionStatement() *Statement {
	panic("implement me")
}

func (parser *Parser) nextToken() {
	parser.curToken = parser.peekToken
	parser.peekToken = parser.lexer.NextToken()
}

func (parser *Parser) curTokenIs(t TokenType) bool {
	return parser.curToken.Type == t
}

func (parser *Parser) peekTokenIs(t TokenType) bool {
	return parser.peekToken.Type == t
}

func (parser *Parser) expectPeek(t TokenType) bool {
	if parser.peekTokenIs(t) {
		parser.nextToken()
		return true
	}
	parser.peekError(t)
	return false
}

func (parser *Parser) Errors() []string {
	return parser.errors
}

func (parser *Parser) peekError(t TokenType) {
	msg := fmt.Sprintf(
		"expected next token to be %s, got %s instead",
		t,
		parser.peekToken.Type,
	)
	parser.errors = append(parser.errors, msg)
}
