package gdx

import (
	"errors"
	"fmt"
)

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
	s.tokens = append(s.tokens, Token{tokenType: tokenType, lexeme: text, literal: literal, line: s.line})
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
			for !s.isAtEnd() && s.peek() != '\n' {
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
	default:
		return errors.New(fmt.Sprintf("Scanning Error [line %d] Unexpected Character: %v", s.line, c))
	}
	return nil
}

// this doesn't advance, it's purely a lookahead
func (s *Scanner) peek() byte {
	if s.isAtEnd() {
		return '\n'
	}

	return s.source[s.current]
}

// only advance if the character is what we expect
func (s *Scanner) match(expected byte) bool {
	if s.isAtEnd() {
		return false
	}
	if s.source[s.current] != expected {
		return false
	}
	s.current = s.current + 1
	return true
}

// advance and return the current character
func (s *Scanner) advance() byte {
	c := s.source[s.current]
	s.current = s.current + 1
	return c
}
