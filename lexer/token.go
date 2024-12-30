package lexer

import "fmt"

type Token struct {
	Kind  TokenKind
	Value string
}

type TokenKind int

const (
	// Literals
	EOF TokenKind = iota
	NUMBER
	STRING
	IDENTIFIER

	// Blocks
	OPEN_BRACKET
	CLOSE_BRACKET
	OPEN_PAREN
	CLOSE_PAREN
	OPEN_BRACE
	CLOSE_BRACE

	// Comment Operator
	DOUBLE_SLASH

	// Operators
	EQUAL
	EQUAL_EQUAL
	NOT
	NOT_EQUAL
	PLUS
	PLUS_PLUS
	PLUS_EQUAL
	MINUS
	MINUS_MINUS
	MINUS_EQUAL
	STAR
	STAR_EQUAL
	SLASH
	SLASH_EQUAL

	// Special Symbols
	DOT

	// Keywords
	LET
	PRINT
	IF
	ELSE
	FOR
	BREAK
	CONTINUE
	FN
	RETURN
	TRUE
	FALSE
)

func NewToken(kind TokenKind, value string) Token {
	return Token{
		kind, value,
	}
}

func (token Token) isOneOf(kinds ...TokenKind) bool {
	for _, kind := range kinds {
		if token.Kind == kind {
			return true
		}
	}

	return false
}

func (token Token) Debug() {
	if token.isOneOf(NUMBER, STRING, IDENTIFIER) {
		fmt.Printf("%s (%s)\n", TokenKindString(token.Kind), token.Value)
	} else {
		fmt.Printf("%s ()\n", TokenKindString(token.Kind))
	}
}

func TokenKindString(kind TokenKind) string {
	switch kind {
	case EOF:
		return "eof"
	case NUMBER:
		return "number"
	case STRING:
		return "string"
	case IDENTIFIER:
		return "identifier"
	case OPEN_BRACKET:
		return "open_bracket"
	case CLOSE_BRACKET:
		return "close_bracket"
	case OPEN_PAREN:
		return "open_paren"
	case CLOSE_PAREN:
		return "close_paren"
	case OPEN_BRACE:
		return "open_brace"
	case CLOSE_BRACE:
		return "close_brace"
	case DOUBLE_SLASH:
		return "double_slash"
	case EQUAL:
		return "equal"
	case EQUAL_EQUAL:
		return "equal_equal"
	case NOT:
		return "not"
	case NOT_EQUAL:
		return "not_equal"
	case PLUS:
		return "plus"
	case PLUS_PLUS:
		return "plus_plus"
	case PLUS_EQUAL:
		return "plus_equal"
	case MINUS:
		return "minus"
	case MINUS_MINUS:
		return "minus_minus"
	case MINUS_EQUAL:
		return "minus_equal"
	case STAR:
		return "star"
	case STAR_EQUAL:
		return "star_equal"
	case SLASH:
		return "slash"
	case SLASH_EQUAL:
		return "slash_equal"
	case DOT:
		return "dot"
	case LET:
		return "let"
	case PRINT:
		return "print"
	case IF:
		return "if"
	case ELSE:
		return "else"
	case FOR:
		return "for"
	case BREAK:
		return "break"
	case CONTINUE:
		return "continue"
	case FN:
		return "fn"
	case RETURN:
		return "return"
	case TRUE:
		return "true"
	case FALSE:
		return "false"
	default:
		return ""
	}
}
