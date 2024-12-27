package lexer

type Token struct {
	Kind  TokenType
	Value string
}

type TokenType int

const (
	// Literals
	NUMBER TokenType = iota
	STRING
	IDENTIFIER

	LEFT_PAREN
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE

	EQUAL
	EQUAL_EQUAL
	NOT
	NOT_EQUAL

	SLASH
	DOUBLE_SLASH

	// Keywords
	LET
	PRINT
	FOR
	FN
	RETURN
	TRUE
	FALSE
)

func NewToken(kind TokenType, value string) Token {
	return Token{kind, value}
}
