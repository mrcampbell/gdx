package gdx

import "fmt"

type TokenType string

type DataType interface {
	int | string | map[string]interface{}
}

const (
	// single character tokens
	LEFT_PAREN  = "("
	RIGHT_PAREN = ")"
	LEFT_BRACE  = "{"
	RIGHT_BRACE = "}"
	COMMA       = ","
	DOT         = "."
	MINUS       = "-"
	PLUS        = "+"
	SEMICOLON   = ";"
	SLASH       = "/"
	STAR        = "*"
	// one-two character tokens
	BANG          = "!"
	BANG_EQUAL    = "!="
	EQUAL         = "="
	EQUAL_EQUAL   = "=="
	GREATER       = ">"
	GREATER_EQUAL = ">="
	LESS          = "<"
	LESS_EQUAL    = "<="

	// literals
	IDENTIFIER = "IDENTIFIER"
	STRING     = "STRING"
	NUMBER     = "NUMBER"

	// keywords
	AND    = "AND"
	CLASS  = "CLASS"
	ELSE   = "ELSE"
	FALSE  = "FALSE"
	FUN    = "FUN"
	FOR    = "FOR"
	IF     = "IF"
	NIL    = "NIL"
	OR     = "OR"
	PRINT  = "PRINT"
	RETURN = "RETURN"
	SUPER  = "SUPER"
	THIS   = "THIS"
	TRUE   = "TRUE"
	VAR    = "VAR"
	WHILE  = "WHILE"

	EOF = "EOF"
)

type Token struct {
	tokenType TokenType
	lexeme    string
	literal   interface{}
	line      int
}

func NewToken(tokenType TokenType, lexeme string, literal interface{}, line int) Token {
	return Token{tokenType, lexeme, literal, line}
}

func (t *Token) ToString() string {
	return fmt.Sprintf("%s %s %v", t.tokenType, string(t.lexeme), t.literal)
}

func (t *Token) Lexeme() string {
	return t.lexeme
}
