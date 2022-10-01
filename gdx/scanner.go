package gdx

import (
	"errors"
	"fmt"
	"strconv"
)

const TERMINATOR = '\x00'

var keywordTokenMap map[string]TokenType = map[string]TokenType{
	"and":    AND,
	"class":  CLASS,
	"else":   ELSE,
	"false":  FALSE,
	"for":    FOR,
	"fun":    FUN,
	"if":     IF,
	"nil":    NIL,
	"or":     OR,
	"print":  PRINT,
	"return": RETURN,
	"super":  SUPER,
	"this":   THIS,
	"true":   TRUE,
	"var":    VAR,
	"while":  WHILE,
}

type Scanner struct {
	source  string
	tokens  []Token
	start   int
	current int
	line    int
}

func NewScanner(source string) *Scanner {
	return &Scanner{source: source, tokens: make([]Token, 0), start: 0, current: 0, line: 1}
}

func (s *Scanner) ScanTokens() ([]Token, []error) {
	// printGreen("SCAN TOKENS", nil)
	errors := []error{}

	for !s.isAtEnd() {
		s.start = s.current
		err := s.scanToken()
		if err != nil {
			errors = append(errors, err)
		}
	}

	s.tokens = append(s.tokens, Token{tokenType: EOF})
	return s.tokens, errors
}

func (s *Scanner) addToken(tokenType TokenType, literal interface{}) {
	text := s.source[s.start:s.current]
	t := Token{tokenType: tokenType, lexeme: text, literal: literal, line: s.line}
	printYellow("ADD TOKEN %+v", t)
	s.tokens = append(s.tokens, t)
}

func (s *Scanner) isAtEnd() bool {
	return s.current == len(s.source)-1
}

func (s *Scanner) scanToken() error {
	c := s.advance()

	switch c {
	case '(':
		s.addToken(LEFT_PAREN, nil)
		break
	case ')':
		s.addToken(RIGHT_PAREN, nil)
		break
	case '{':
		s.addToken(LEFT_BRACE, nil)
		break
	case '}':
		s.addToken(RIGHT_BRACE, nil)
		break
	case ',':
		s.addToken(COMMA, nil)
		break
	case '.':
		s.addToken(DOT, nil)
		break
	case '-':
		s.addToken(MINUS, nil)
		break
	case '+':
		s.addToken(PLUS, nil)
		break
	case ';':
		s.addToken(SEMICOLON, nil)
		break
	case '*':
		s.addToken(STAR, nil)
		break
	case '!':
		if s.match('=') {
			s.addToken(BANG_EQUAL, nil)
		} else {
			s.addToken(BANG, nil)
		}
		break
	case '=':
		if s.match('=') {
			s.addToken(EQUAL_EQUAL, nil)
		} else {
			s.addToken(EQUAL, nil)
		}
		break
	case '<':
		if s.match('=') {
			s.addToken(LESS_EQUAL, nil)
		} else {
			s.addToken(LESS, nil)
		}

		break
	case '>':
		if s.match('=') {
			s.addToken(GREATER_EQUAL, nil)
		} else {
			s.addToken(GREATER, nil)
		}
	case '/':
		if s.match('/') {
			for !s.isAtEnd() && s.peek() != TERMINATOR {
				s.advance()
			}
		} else {
			s.addToken(SLASH, nil)
		}
	case ' ':
	case '\r':
	case '\t':
		// ignore whitespace
		break
	case '\n':
		s.line = s.line + 1
		break

	case '"':
		err := s.string()
		if err != nil {
			return err
		}
		break
	default:
		if s.isDigit(c) {
			err := s.number()
			if err != nil {
				return err
			}
			break
		}
		if s.isAlpha(c) {
			s.identifier()
			break
		}

		return errors.New(fmt.Sprintf("Scanning Error [line %d] Unexpected Character: %v", s.line, c))
	}
	return nil
}

func (s *Scanner) isDigit(c byte) bool {
	if c == TERMINATOR {
		printGreen("IS DIGIT: TERMINATED, RETURNING FALSE")
		return false
	}
	var result bool = (c >= '0' && c <= '9') && true
	printGreen("IS DIGIT: %t", result)
	return result
}

func (s *Scanner) isAlpha(c byte) bool {
	if c == TERMINATOR {
		printGreen("IS ALPHA: TERMINATED, RETURNING FALSE")
		return false
	}
	result := (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		c == '_'
	printGreen("IS ALPHA: ", result)

	return result
}

func (s *Scanner) isAlphaNumeric(c byte) bool {
	printGreen("IS ALPHANUMERIC", string(c))
	return s.isAlpha(c) || s.isDigit(c)
}

func (s *Scanner) identifier() {
	printGreen("IDENTIFIER")
	for s.isAlphaNumeric(s.peek()) {
		s.advance()
	}
	text := s.source[s.start:s.current]
	if token, ok := keywordTokenMap[text]; ok {
		s.addToken(token, nil)
		return
	}

	s.addToken(IDENTIFIER, nil)
	return
}

func (s *Scanner) number() error {
	for s.isDigit(s.peek()) {
		s.advance()
	}
	printRed("HERE")
	if s.peek() == '.' && s.isDigit(s.peekNext()) {
		fmt.Println("There's a number after the dot")
		s.advance() // consume the '.'

		// advance along the rest of the decimal
		for s.isDigit(s.peek()) {
			s.advance()
		}
	} else if s.peek() == '.' {
		return errors.New("There is no number after the dot")
	}
	printRed("END HERE")

	if f, err := strconv.ParseFloat(s.source[s.start:s.current], 32); err == nil {
		s.addToken(NUMBER, f)
	}

	return nil
}

func (s *Scanner) string() error {
	for s.peek() != '"' && !s.isAtEnd() {
		if s.peek() == '\n' {
			s.line = s.line + 1
		}
		s.advance()
	}

	if s.isAtEnd() {
		return errors.New(fmt.Sprintf("Unterminated string at line! %d", 0))
	}

	s.advance() // the cloing '"'

	value := s.source[s.start+1 : s.current-1]
	s.addToken(STRING, value)
	fmt.Println("Got string value: ", value)
	return nil
}

func (s *Scanner) peekNext() byte {
	if s.current+1 > len(s.source) {
		Debug("PEEK NEXT: AT END OF FILE, RETURNING TERMINATOR")
		return TERMINATOR
	}
	Debug("PEEK NEXT: ", s.source[s.current+1])
	return s.source[s.current+1]
}

// this doesn't advance, it's purely a lookahead
func (s *Scanner) peek() byte {
	if s.isAtEnd() {
		Debug("PEEK: AT END OF FILE, RETURNING TERMINATOR")
		return TERMINATOR
	}

	Debug("PEEK: ", string(s.source[s.current]))
	return s.source[s.current]
}

// only advance if the character is what we expect
func (s *Scanner) match(expected byte) bool {
	if s.isAtEnd() {
		Debug("MATCH, AT END OF FILE, Returning `false`")
		return false
	}
	if s.source[s.current] != expected {
		Debug("MATCH: NO MATCH", "current", s.current, "expected", expected)
		return false
	}
	s.current = s.current + 1
	return true
}

// advance and return the current character
func (s *Scanner) advance() byte {
	c := s.source[s.current]
	s.current = s.current + 1
	Debug("ADVANCE: ", string(c))
	return c
}
