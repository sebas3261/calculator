package parser

import (
	"errors"

	"github.com/sebas3261/calculator/backend/internal/calculator"
)

func (s *parserState) parsePrimary() (calculator.Node, error) {
	s.skipWhitespace()

	var node calculator.Node
	var err error

	switch {
	case s.isDigit() || s.current() == '.':
		node, err = s.parseNumber()
	case s.current() == '(':
		s.advance()

		node, err = s.parseExpression()
		if err != nil {
			return nil, err
		}

		s.skipWhitespace()

		if s.current() != ')' {
			return nil, errors.New("missing closing parenthesis")
		}

		s.advance()
	default:
		return nil, errors.New("unexpected character")
	}

	if err != nil {
		return nil, err
	}

	return s.parsePercent(node)
}

// parsePercent applies the postfix % operator (e.g. "50%"), which binds
// tighter than any other operator.
func (s *parserState) parsePercent(node calculator.Node) (calculator.Node, error) {
	s.skipWhitespace()

	if s.current() != '%' {
		return node, nil
	}

	s.advance()

	return calculator.UnaryNode{Operator: "%", Operand: node}, nil
}

// parsePower handles the right-associative exponentiation operator (^),
// which binds tighter than unary +/-/√ (so "-2^2" is "-(2^2)").
func (s *parserState) parsePower() (calculator.Node, error) {
	left, err := s.parsePrimary()
	if err != nil {
		return nil, err
	}

	s.skipWhitespace()

	if s.current() != '^' {
		return left, nil
	}

	s.advance()

	right, err := s.parseUnary()
	if err != nil {
		return nil, err
	}

	return calculator.BinaryNode{
		Left:     left,
		Operator: "^",
		Right:    right,
	}, nil
}

func (s *parserState) parseUnary() (calculator.Node, error) {
	s.skipWhitespace()

	if s.current() == '+' || s.current() == '-' || s.current() == '√' {
		operator := string(s.current())
		s.advance()

		operand, err := s.parseUnary()
		if err != nil {
			return nil, err
		}

		return calculator.UnaryNode{
			Operator: operator,
			Operand:  operand,
		}, nil
	}

	return s.parsePower()
}

func (s *parserState) parseTerm() (calculator.Node, error) {
	left, err := s.parseUnary()
	if err != nil {
		return nil, err
	}

	for {
		s.skipWhitespace()

		if s.current() != '*' && s.current() != '/' {
			break
		}

		operator := string(s.current())
		s.advance()

		right, err := s.parseUnary()
		if err != nil {
			return nil, err
		}

		left = calculator.BinaryNode{
			Left:     left,
			Operator: operator,
			Right:    right,
		}
	}

	return left, nil
}

func (s *parserState) parseExpression() (calculator.Node, error) {
	left, err := s.parseTerm()
	if err != nil {
		return nil, err
	}

	for {
		s.skipWhitespace()

		if s.current() != '+' && s.current() != '-' {
			break
		}

		operator := string(s.current())
		s.advance()

		right, err := s.parseTerm()
		if err != nil {
			return nil, err
		}

		left = calculator.BinaryNode{
			Left:     left,
			Operator: operator,
			Right:    right,
		}
	}

	return left, nil
}
