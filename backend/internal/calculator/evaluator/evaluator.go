package evaluator

import (
	"errors"
	"fmt"
	"math"

	"github.com/sebas3261/calculator/backend/internal/calculator"
)

type Evaluator struct{}

func NewEvaluator() *Evaluator {
	return &Evaluator{}
}

func (e *Evaluator) Evaluate(node calculator.Node) (float64, error) {
	switch n := node.(type) {
	case calculator.NumberNode:
		return n.Value, nil
	case calculator.UnaryNode:
		return e.evaluateUnary(n)
	case calculator.BinaryNode:
		return e.evaluateBinary(n)
	default:
		return 0, fmt.Errorf("unsupported node type %T", node)
	}
}

func (e *Evaluator) evaluateUnary(n calculator.UnaryNode) (float64, error) {
	operand, err := e.Evaluate(n.Operand)
	if err != nil {
		return 0, err
	}

	var result float64

	switch n.Operator {
	case "+":
		result = operand
	case "-":
		result = -operand
	case "√":
		if operand < 0 {
			return 0, errors.New("cannot take the square root of a negative number")
		}
		result = math.Sqrt(operand)
	case "%":
		result = operand / 100
	default:
		return 0, fmt.Errorf("unsupported unary operator %q", n.Operator)
	}

	return finite(result)
}

func (e *Evaluator) evaluateBinary(n calculator.BinaryNode) (float64, error) {
	left, err := e.Evaluate(n.Left)
	if err != nil {
		return 0, err
	}

	right, err := e.Evaluate(n.Right)
	if err != nil {
		return 0, err
	}

	var result float64

	switch n.Operator {
	case "+":
		result = left + right
	case "-":
		result = left - right
	case "*":
		result = left * right
	case "/":
		if right == 0 {
			return 0, errors.New("division by zero")
		}
		result = left / right
	case "^":
		result = math.Pow(left, right)
	default:
		return 0, fmt.Errorf("unsupported binary operator %q", n.Operator)
	}

	return finite(result)
}

func finite(value float64) (float64, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return 0, errors.New("result is not a finite number")
	}

	return value, nil
}
