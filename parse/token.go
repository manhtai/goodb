package parse

type Keyword string

const (
	SelectKeyword Keyword = "select"
	FromKeyword   Keyword = "from"
	WhereKeyword  Keyword = "where"
	AndKeyword    Keyword = "and"
	InsertKeyword Keyword = "insert"
	IntoKeyword   Keyword = "into"
	ValuesKeyword Keyword = "values"
	DeleteKeyword Keyword = "delete"
	UpdateKeyword Keyword = "update"
	SetKeyword    Keyword = "set"
	TableKeyword  Keyword = "table"
	CreateKeyword Keyword = "create"
	IntKeyword    Keyword = "int"
	TextKeyword   Keyword = "varchar"
	ViewKeyword   Keyword = "view"
	AsKeyword     Keyword = "as"
	IndexKeyword  Keyword = "index"
	OnKeyword     Keyword = "on"
)

type Symbol string

const (
	CommaSymbol      Symbol = ","
	LeftParenSymbol  Symbol = "("
	RightParenSymbol Symbol = ")"
	EqSymbol         Symbol = "="
)

type TokenKind uint

const (
	KeywordKind TokenKind = iota
	SymbolKind
	IdentifierKind
	ConstantKind
)

type Token struct {
	Value string
	Kind  TokenKind
}
