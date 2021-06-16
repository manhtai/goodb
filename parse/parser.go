package parse

import (
	"fmt"
	"goodb/query"
	"strconv"
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
	var fields []string
	for !parser.curTokenIs(FromKeyword) {
		fields = append(fields, parser.curToken.Literal)
		parser.nextToken()
	}

	parser.nextToken()

	var tables []string
	for !parser.curTokenIs(SemicolonSymbol) {
		tables = append(fields, parser.curToken.Literal)
		parser.nextToken()
	}

	predicate := &query.Predicate{}
	if parser.curTokenIs(WhereKeyword) {
		predicate = parser.parsePredicate()
	}

	stmt := &SelectStatement{
		fields: fields,
		tables: tables,
		predicate: predicate,
	}

	return &Statement{
		SelectStatement: stmt,
		Kind: SelectKind,
	}
}

func (parser *Parser) parsePredicate() *query.Predicate {
	term := parser.parseTerm()
	return query.NewPredicateFromTerm(term)
}

func (parser *Parser) parseTerm() *query.Term {
	left := parser.parseExpression()
	parser.nextToken()
	right := parser.parseExpression()
	return query.NewTerm(left, right)
}

func (parser *Parser) parseExpression() *query.Expression {
	if parser.curTokenIs(Identifier) {
		parser.nextToken()
		return query.NewFieldExpression(parser.curToken.Literal)
	}
	parser.nextToken()
	return query.NewConstantExpression(parseConstant(parser.curToken))
}

func parseConstant(token Token) *query.Constant {
	if token.Type == StringConstant {
		return query.NewStrConstant(token.Literal)
	}
	i, _ := strconv.Atoi(token.Literal)
	return query.NewIntConstant(i)
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
