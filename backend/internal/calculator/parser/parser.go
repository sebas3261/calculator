package parser

import (
	"errors"
	"unicode/utf8"

	"github.com/sebas3261/calculator/backend/internal/calculator"
)

type ExpressionParser struct{}

type parserState struct {
	input string
	pos   int
}

func NewExpressionParser() *ExpressionParser {
	return &ExpressionParser{}
}

func (p *ExpressionParser) Parse(expression string) (calculator.Node, error) {
	state := &parserState{input: expression}

	state.skipWhitespace()

	if state.current() == 0 {
		return nil, errors.New("expression cannot be empty")
	}

	node, err := state.parseExpression()
	if err != nil {
		return nil, err
	}

	state.skipWhitespace()

	if state.current() != 0 {
		return nil, errors.New("unexpected character at end of expression")
	}

	return node, nil
}


func (s *parserState) current() rune {
	if s.pos >= len(s.input) {
		return 0
	}

	r, _ := utf8.DecodeRuneInString(s.input[s.pos:])
	return r
}

func (s *parserState) advance() {
	if s.pos >= len(s.input) {
		return
	}

	_, size := utf8.DecodeRuneInString(s.input[s.pos:])
	s.pos += size
}

func (s *parserState) skipWhitespace() {
	for s.current() == ' ' {
		s.advance()
	}
}

func (s *parserState) isDigit() bool {
	return s.current() >= '0' && s.current() <= '9'
}
