package lexer

import (
	"fmt"
	"regexp"
)

type regexHandler func(lex *lexer, regex *regexp.Regexp)

type regexPattern struct {
	regex   *regexp.Regexp
	handler regexHandler
}

type lexer struct {
	pattrens []regexPattern
	tokens   []Token
	source   string
	pos      int
}

func (lex *lexer) Tokenize() []Token {
	for !lex.atEOF() {
		var matched bool

		for _, pattern := range lex.pattrens {
			loc := pattern.regex.FindStringIndex(lex.remainder())

			if loc != nil && loc[0] == 0 {
				pattern.handler(lex, pattern.regex)
				matched = true
				break
			}
		}

		if !matched {
			panic(fmt.Sprintf("LEXER::ERR : unrecognized token at %d", lex.pos))
		}
	}

	lex.push(NewToken(EOF, "EOF"))
	return lex.tokens
}

func (lex *lexer) advanceN(n int) {
	lex.pos += n
}

func (lex *lexer) push(token Token) {
	lex.tokens = append(lex.tokens, token)
}

func (lex *lexer) at() byte {
	return lex.source[lex.pos]
}

func (lex *lexer) remainder() string {
	return lex.source[lex.pos:]
}

func (lex *lexer) atEOF() bool {
	return lex.pos >= len(lex.source)
}

func defaultHandler(kind TokenKind, value string) regexHandler {
	return func(lex *lexer, regex *regexp.Regexp) {
		lex.advanceN(len(value))
		lex.push(NewToken(kind, value))
	}
}

func CreateLexer(source string) *lexer {
	return &lexer{
		source: source,
		pos:    0,
		tokens: make([]Token, 0),
		pattrens: []regexPattern{
			{regexp.MustCompile(`[0-9]+(\.[0-9]+)?`), numberHandler},
			{regexp.MustCompile(`\[`), defaultHandler(OPEN_BRACKET, "[")},
			{regexp.MustCompile(`\]`), defaultHandler(CLOSE_BRACKET, "]")},
			{regexp.MustCompile(`\(`), defaultHandler(OPEN_PAREN, "(")},
			{regexp.MustCompile(`\)`), defaultHandler(CLOSE_PAREN, ")")},
			{regexp.MustCompile(`\{`), defaultHandler(OPEN_BRACE, "{")},
			{regexp.MustCompile(`\}`), defaultHandler(CLOSE_BRACE, "}")},
			{regexp.MustCompile(`\+`), defaultHandler(PLUS, "+")},
			{regexp.MustCompile(`-`), defaultHandler(MINUS, "-")},
			{regexp.MustCompile(`\*`), defaultHandler(STAR, "*")},
			{regexp.MustCompile(`\/`), defaultHandler(SLASH, "/")},
			// {regexp.MustCompile(`\[`), defaultHandler(OPEN_BRACKET, "[")},
			// {regexp.MustCompile(`\[`), defaultHandler(OPEN_BRACKET, "[")},
			// {regexp.MustCompile(`\[`), defaultHandler(OPEN_BRACKET, "[")},
			// {regexp.MustCompile(`\[`), defaultHandler(OPEN_BRACKET, "[")},
			// {regexp.MustCompile(`\[`), defaultHandler(OPEN_BRACKET, "[")},
			// {regexp.MustCompile(`\[`), defaultHandler(OPEN_BRACKET, "[")},
			// {regexp.MustCompile(`\[`), defaultHandler(OPEN_BRACKET, "[")},
			// {regexp.MustCompile(`\[`), defaultHandler(OPEN_BRACKET, "[")},
			// {regexp.MustCompile(`\[`), defaultHandler(OPEN_BRACKET, "[")},
			// {regexp.MustCompile(`\[`), defaultHandler(OPEN_BRACKET, "[")},
			// {regexp.MustCompile(`\[`), defaultHandler(OPEN_BRACKET, "[")},
			// {regexp.MustCompile(`\[`), defaultHandler(OPEN_BRACKET, "[")},
			// {regexp.MustCompile(`\[`), defaultHandler(OPEN_BRACKET, "[")},
			// {regexp.MustCompile(`\[`), defaultHandler(OPEN_BRACKET, "[")},
		},
	}
}

func numberHandler(lex *lexer, regex *regexp.Regexp) {
	match := regex.FindString(lex.remainder())
	lex.push(NewToken(NUMBER, match))
	lex.advanceN(len(match))
}
