package parse

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	SelectKeyword = "select"
	FromKeyword   = "from"
	WhereKeyword  = "where"
	AndKeyword    = "and"
	InsertKeyword = "insert"
	IntoKeyword   = "into"
	ValuesKeyword = "values"
	DeleteKeyword = "delete"
	UpdateKeyword = "update"
	SetKeyword    = "set"
	TableKeyword  = "table"
	CreateKeyword = "create"
	IntKeyword    = "int"
	TextKeyword   = "varchar"
	ViewKeyword   = "view"
	AsKeyword     = "as"
	IndexKeyword  = "index"
	OnKeyword     = "on"

	SemicolonSymbol  = ";"
	CommaSymbol      = ","
	LeftParenSymbol  = "("
	RightParenSymbol = ")"
	EqSymbol         = "="
	IllegalSymbol    = "ILLEGAL"

	Identifier     = "IDENT"
	IntConstant    = "INT"
	StringConstant = "STRING"
)

var keywords = map[string]TokenType{
	"select":  SelectKeyword,
	"from":    FromKeyword,
	"where":   WhereKeyword,
	"and":     AndKeyword,
	"insert":  InsertKeyword,
	"into":    IntoKeyword,
	"values":  ValuesKeyword,
	"delete":  DeleteKeyword,
	"update":  UpdateKeyword,
	"set":     SetKeyword,
	"table":   TableKeyword,
	"create":  CreateKeyword,
	"int":     IntKeyword,
	"varchar": TextKeyword,
	"view":    ViewKeyword,
	"as":      AsKeyword,
	"index":   IndexKeyword,
	"on":      OnKeyword,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return Identifier
}
