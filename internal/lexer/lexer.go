package main

import (
	"fmt"
	"io"
	"strings"
)

type sqlLexer struct {
	yyLexer

	input io.Reader
	state int

	// lexer-specific state
	pos    int
	buffer string
}

func newSQLLexer(input io.Reader) *sqlLexer {
	return &sqlLexer{
		input: input,
		state: 0,
		pos:   0,
	}
}

func (l *sqlLexer) Lex(lval *yySymType) int {
	for {
		switch l.state {
		case 0:
			// whitespace and comments
			for {
				c := l.read()
				if c == 0 {
					return 0
				}
				if !isWhitespace(c) {
					l.unread()
					break
				}
			}
			if l.peek() == '-' {
				l.read()
				if l.read() == '-' {
					l.state = 1
					break
				}
				l.unread()
				l.unread()
			}
			if l.peek() == '/' {
				l.read()
				if l.read() == '*' {
					l.state = 2
					break
				}
				l.unread()
				l.unread()
			}
			l.state = 3
		case 1:
			// single-line comment
			for {
				c := l.read()
				if c == 0 || c == '\n' {
					l.unread()
					break
				}
			}
			l.state = 0
		case 2:
			// multi-line comment
			for {
				c := l.read()
				if c == 0 {
					return 0
				}
				if c == '*' && l.peek() == '/' {
					l.read()
					l.state = 0
					break
				}
			}
		case 3:
			// keywords, identifiers, literals, and symbols
			c := l.read()
			if c == 0 {
				return 0
			}
			if isAlphaNumeric(c) {
				l.buffer += string(c)
				for {
					c = l.read()
					if !isAlphaNumeric(c) {
						l.unread()
						break
					}
					l.buffer += string(c)
				}
				if keyword, ok := keywords[strings.ToUpper(l.buffer)]; ok {
					return keyword
				}
				lval.sval = l.buffer
				return ID
			}
			switch c {
			case ',':
				return COMMA
			case '(':
				return LPAREN
			case ')':
				return RPAREN
			case '=':
				if l.read() == '=' {
					return EQ
				}
				l.unread()
				return '='
			case '<':
				if l.read() == '=' {
					return LTE
				}
				l.unread()
				return LT
			case '>':
				if l.read() == '=' {
					return GTE
				}
				l.unread()
				return GT
			case '!':
				if l.read() == '=' {
					return NEQ
				}
				l.unread()
				return '!'
			case '\'':
				for {
					c = l.read()
					if c == 0 {
						return 0
					}
					if c == '\'' {
						break
					}
					l.buffer += string(c)
				}
				lval.sval = l.buffer
				l.buffer = ""
				return STRING
			default:
				// ignore unrecognized characters
			}
		}
	}
}

func (l *sqlLexer) Error(s string) {
	fmt.Printf("Error at position %d: %s\n", l.pos, s)
}

func (l *sqlLexer) read() byte {
	buf := make([]byte, 1)
	n, err := l.input.Read(buf)
	if n == 0 || err != nil {
		return 0
	}
	l.pos++
	return buf[0]
}

func (l *sqlLexer) peek() byte {
	buf := make([]byte, 1)
	n, err := l.input.Read(buf)
	if n == 0 || err != nil {
		return 0
	}
	l.input.UnreadByte()
	return buf[0]
}

func (l *sqlLexer) unread() {
	l.input.UnreadByte()
	l.pos--
}

func isAlphaNumeric(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') || c == '_'
}

func isWhitespace(c byte) bool {
	return c == ' ' || c == '\t' || c == '\r' || c == '\n'
}

var keywords = map[string]int{
	"SELECT":   SELECT,
	"INSERT":   INSERT,
	"INTO":     INTO,
	"VALUES":   VALUES,
	"FROM":     FROM,
	"WHERE":    WHERE,
	"AND":      AND,
	"OR":       OR,
	"NOT":      NOT,
	"GROUP":    GROUP,
	"BY":       BY,
	"HAVING":   HAVING,
	"ORDER":    ORDER,
	"ASC":      ASC,
	"DESC":     DESC,
	"LIMIT":    LIMIT,
	"OFFSET":   OFFSET,
}	

