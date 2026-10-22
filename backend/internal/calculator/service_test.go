package calculator_test

import (
	"errors"
	"math"
	"testing"

	"github.com/sebas3261/calculator/backend/internal/calculator"
	"github.com/sebas3261/calculator/backend/internal/calculator/evaluator"
	"github.com/sebas3261/calculator/backend/internal/calculator/parser"
)

// fakeParser and fakeEvaluator let us test the Service's error-propagation
// contract in isolation from the real parser/evaluator implementations.
type fakeParser struct {
	node calculator.Node
	err  error
}

func (f fakeParser) Parse(string) (calculator.Node, error) { return f.node, f.err }

type fakeEvaluator struct {
	value float64
	err   error
}

func (f fakeEvaluator) Evaluate(calculator.Node) (float64, error) { return f.value, f.err }

func newRealService() *calculator.Service {
	return calculator.NewService(parser.NewExpressionParser(), evaluator.NewEvaluator())
}

func TestService_Calculate_EndToEnd(t *testing.T) {
	cases := []struct {
		name string
		expr string
		want float64
	}{
		{"operator_precedence", "2 + 3 * 4", 14},
		{"parentheses_override_precedence", "(2 + 3) * 4", 20},
		{"power_and_root", "√9 + 2 ^ 3", 11},
		{"percent", "50% * 10", 5},
		{"unary_minus", "-5 + 10", 5},
	}

	service := newRealService()

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := service.Calculate(calculator.CalculationRequest{Expression: tc.expr})
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if math.Abs(got.Answer-tc.want) > 1e-9 {
				t.Errorf("Calculate(%q) = %v, want %v", tc.expr, got.Answer, tc.want)
			}
		})
	}
}

func TestService_Calculate_EmptyExpression(t *testing.T) {
	service := newRealService()

	cases := []string{"", "   "}

	for _, expr := range cases {
		_, err := service.Calculate(calculator.CalculationRequest{Expression: expr})
		if err == nil {
			t.Errorf("Calculate(%q): expected error, got nil", expr)
		}
	}
}

func TestService_Calculate_PropagatesParserError(t *testing.T) {
	wantErr := errors.New("boom from parser")
	service := calculator.NewService(fakeParser{err: wantErr}, fakeEvaluator{})

	_, err := service.Calculate(calculator.CalculationRequest{Expression: "1 + 1"})
	if !errors.Is(err, wantErr) {
		t.Fatalf("got error %v, want %v", err, wantErr)
	}
}

func TestService_Calculate_PropagatesEvaluatorError(t *testing.T) {
	wantErr := errors.New("boom from evaluator")
	service := calculator.NewService(fakeParser{node: calculator.NumberNode{Value: 1}}, fakeEvaluator{err: wantErr})

	_, err := service.Calculate(calculator.CalculationRequest{Expression: "1"})
	if !errors.Is(err, wantErr) {
		t.Fatalf("got error %v, want %v", err, wantErr)
	}
}

func TestService_Calculate_DivisionByZero(t *testing.T) {
	service := newRealService()

	_, err := service.Calculate(calculator.CalculationRequest{Expression: "1 / 0"})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestService_Calculate_SquareRootOfNegative(t *testing.T) {
	service := newRealService()

	_, err := service.Calculate(calculator.CalculationRequest{Expression: "√(0 - 9)"})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
