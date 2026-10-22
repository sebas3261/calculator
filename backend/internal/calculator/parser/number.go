package parser

import (
	"errors"
	"strconv"

	"github.com/sebas3261/calculator/backend/internal/calculator"
)

func (s *parserState) parseNumber() (calculator.Node, error) {
	start := s.pos
	hasDecimal := false

	for s.isDigit() || s.current() == '.' {
		if s.current() == '.' {
			if hasDecimal {
				return nil, errors.New("number cannot contain multiple decimal points")
			}

			hasDecimal = true
		}

		s.advance()
	}

	number := s.input[start:s.pos]
	value, err := strconv.ParseFloat(number, 64)
	if err != nil {
		return nil, err
	}

	return calculator.NumberNode{Value: value}, nil
}
