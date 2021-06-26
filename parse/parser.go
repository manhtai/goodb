package parse

import (
	"fmt"
	"goodb/query"
	"goodb/record"
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
	p := &Parser{
		lexer: lexer,
	}
	p.nextToken()
	p.nextToken()
	return p
}

func (parser *Parser) ParseStatement() Statement {
	switch parser.curToken.Type {
	case SelectKeyword:
		return parser.parseSelectStatement()
	case CreateKeyword:
		parser.nextToken()
		switch parser.curToken.Type {
		case TableKeyword:
			return parser.parseCreateTableStatement()
		case IndexKeyword:
			return parser.parseCreateIndexStatement()
		}
	case InsertKeyword:
		return parser.parseInsertStatement()
	case DeleteKeyword:
		return parser.parseDeleteStatement()
	case UpdateKeyword:
		return parser.parseUpdateStatement()
	}
	panic("invalid statement")
}

func (parser *Parser) parseSelectStatement() Statement {
	var fields []string
	for !parser.curTokenIs(FromKeyword) {
		if parser.curTokenIs(Identifier) {
			fields = append(fields, parser.curToken.Literal)
		}
		parser.nextToken()
	}

	var tables []string
	for !parser.curTokenIs(EOF) {
		if parser.curTokenIs(Identifier) {
			tables = append(tables, parser.curToken.Literal)
		}
		parser.nextToken()

		if parser.curTokenIs(CommaSymbol) {
			parser.nextToken()
		}

		if parser.curTokenIs(WhereKeyword) {
			break
		}
	}

	pred := parser.parsePredicate()

	stmt := SelectStatement{
		Fields:    fields,
		Tables:    tables,
		Predicate: pred,
	}

	return Statement{
		SelectStatement: stmt,
		Kind:            SelectKind,
	}
}

func (parser *Parser) parseTerm() query.Term {
	parser.nextToken() // left
	left := parser.parseExpression()
	parser.nextToken() // =
	parser.nextToken() // right
	right := parser.parseExpression()
	parser.nextToken()

	return query.NewTerm(left, right)
}

func (parser *Parser) parseExpression() query.Expression {
	if parser.curTokenIs(Identifier) {
		return query.NewFieldExpression(parser.curToken.Literal)
	}
	return query.NewConstantExpression(parseConstant(parser.curToken))
}

func parseConstant(token Token) query.Constant {
	if token.Type == StringConstant {
		return query.NewStrConstant(token.Literal)
	}
	i, _ := strconv.Atoi(token.Literal)
	return query.NewIntConstant(i)
}

func (parser *Parser) parseCreateTableStatement() Statement {
	parser.nextToken()
	tableName := parser.curToken.Literal
	parser.nextToken() // (
	schema := parser.parseFieldDefs()
	stmt := CreateTableStatement{
		TableName: tableName,
		Schema:    schema,
	}
	return Statement{
		CreateTableStatement: stmt,
		Kind:                 CreateTableKind,
	}
}

func (parser *Parser) parseFieldDefs() record.Schema {
	parser.nextToken()
	schema := parser.parseFieldDef()
	for !parser.curTokenIs(RightParenSymbol) {
		if parser.curTokenIs(CommaSymbol) {
			parser.nextToken()
			continue
		}
		fSchema := parser.parseFieldDef()
		schema.Add(fSchema)

	}
	return schema
}

func (parser *Parser) parseFieldDef() record.Schema {
	fieldName := parser.curToken.Literal
	schema := record.NewSchema()

	if parser.peekTokenIs(IntKeyword) {
		parser.nextToken()
		schema.AddIntField(fieldName)
	} else {
		parser.nextToken() // varchar
		parser.nextToken() // (
		parser.nextToken() // length
		fieldLength, _ := strconv.Atoi(parser.curToken.Literal)
		parser.nextToken() // )
		schema.AddStringField(fieldName, fieldLength)
	}
	parser.nextToken()
	return *schema
}

func (parser *Parser) parseCreateIndexStatement() Statement {
	panic("implement me")
}

func (parser *Parser) parseInsertStatement() Statement {
	parser.nextToken() // into
	parser.nextToken() // table
	tableName := parser.curToken.Literal

	parser.nextToken() // (

	var fields []string
	for !parser.curTokenIs(RightParenSymbol) {
		if parser.curTokenIs(Identifier) {
			fields = append(fields, parser.curToken.Literal)
		}
		parser.nextToken()
	}

	parser.nextToken() // Values
	parser.nextToken() // (

	parser.nextToken()
	var values []query.Constant
	for !parser.curTokenIs(RightParenSymbol) {
		if parser.curTokenIs(StringConstant) || parser.curTokenIs(IntConstant) {
			values = append(values, parseConstant(parser.curToken))
		}
		parser.nextToken()
	}

	stmt := InsertStatement{
		TableName: tableName,
		Fields:    fields,
		Values:    values,
	}
	return Statement{
		InsertStatement: stmt,
		Kind:            InsertKind,
	}
}

func (parser *Parser) parseDeleteStatement() Statement {
	parser.nextToken() // from

	parser.nextToken() // table
	tableName := parser.curToken.Literal

	parser.nextToken() // where
	pred := parser.parsePredicate()
	stmt := DeleteStatement{
		TableName: tableName,
		Predicate: pred,
	}

	return Statement{
		DeleteStatement: stmt,
		Kind:            DeleteKind,
	}
}

func (parser *Parser) parseUpdateStatement() Statement {
	parser.nextToken() // table
	tableName := parser.curToken.Literal

	parser.nextToken() // set

	parser.nextToken() // field
	fieldName := parser.curToken.Literal

	parser.nextToken() // =

	parser.nextToken()
	exp := parser.parseExpression()

	parser.nextToken() // where
	pred := parser.parsePredicate()

	stmt := UpdateStatement{
		TableName:  tableName,
		FieldName:  fieldName,
		Expression: exp,
		Predicate:  pred,
	}

	return Statement{
		UpdateStatement: stmt,
		Kind:            UpdateKind,
	}
}

func (parser *Parser) parseExpressionStatement() query.Expression {
	var exp query.Expression
	if parser.curTokenIs(Identifier) {
		exp = query.NewFieldExpression(parser.curToken.Literal)
	} else {
		exp = query.NewConstantExpression(parseConstant(parser.curToken))
	}
	return exp
}

func (parser *Parser) parsePredicate() query.Predicate {
	terms := make([]query.Term, 0)
	for !parser.curTokenIs(EOF) {
		term := parser.parseTerm()
		terms = append(terms, term)

		if parser.peekTokenIs(AndKeyword) {
			parser.nextToken()
		}
	}

	return query.NewPredicateFromTerms(terms)
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
