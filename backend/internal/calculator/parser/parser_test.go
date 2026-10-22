package parser_test

import (
	"reflect"
	"testing"

	"github.com/sebas3261/calculator/backend/internal/calculator"
	"github.com/sebas3261/calculator/backend/internal/calculator/parser"
)

func num(v float64) calculator.Node { return calculator.NumberNode{Value: v} }

func bin(l calculator.Node, op string, r calculator.Node) calculator.Node {
	return calculator.BinaryNode{Left: l, Operator: op, Right: r}
}

func un(op string, operand calculator.Node) calculator.Node {
	return calculator.UnaryNode{Operator: op, Operand: operand}
}

func TestParse_ValidExpressions(t *testing.T) {
	cases := []struct {
		name string
		expr string
		want calculator.Node
	}{
		{"integer", "42", num(42)},
		{"decimal", "3.14", num(3.14)},
		{"addition", "1 + 2", bin(num(1), "+", num(2))},
		{"subtraction", "5 - 3", bin(num(5), "-", num(3))},
		{"multiplication_precedence_over_addition", "2 + 3 * 4", bin(num(2), "+", bin(num(3), "*", num(4)))},
		{"left_associative_division", "2 * 3 / 4", bin(bin(num(2), "*", num(3)), "/", num(4))},
		{"parentheses_override_precedence", "(1 + 2) * 3", bin(bin(num(1), "+", num(2)), "*", num(3))},
		{"unary_minus", "-5", un("-", num(5))},
		{"unary_plus", "+5", un("+", num(5))},
		{"unary_minus_binds_tighter_than_multiplication", "-2 * 3", bin(un("-", num(2)), "*", num(3))},
		{"power", "2 ^ 3", bin(num(2), "^", num(3))},
		{"power_is_right_associative", "2 ^ 3 ^ 2", bin(num(2), "^", bin(num(3), "^", num(2)))},
		{"power_binds_tighter_than_unary_minus", "-2 ^ 2", un("-", bin(num(2), "^", num(2)))},
		{"power_with_negative_exponent", "2 ^ -2", bin(num(2), "^", un("-", num(2)))},
		{"square_root", "√9", un("√", num(9))},
		{"square_root_binds_like_unary", "√9 + 1", bin(un("√", num(9)), "+", num(1))},
		{"square_root_of_parenthesized_expression", "√(4 + 5)", un("√", bin(num(4), "+", num(5)))},
		{"percent", "50%", un("%", num(50))},
		{"percent_binds_tighter_than_multiplication", "50% * 2", bin(un("%", num(50)), "*", num(2))},
		{"whitespace_is_ignored", "  1   +   2  ", bin(num(1), "+", num(2))},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := parser.NewExpressionParser().Parse(tc.expr)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Parse(%q) = %#v, want %#v", tc.expr, got, tc.want)
			}
		})
	}
}

func TestParse_Errors(t *testing.T) {
	cases := []struct {
		name string
		expr string
	}{
		{"empty_expression", ""},
		{"whitespace_only", "   "},
		{"trailing_garbage", "1 2"},
		{"unexpected_character", "@"},
		{"missing_closing_parenthesis", "(1 + 2"},
		{"multiple_decimal_points", "1.2.3"},
		{"dangling_binary_operator", "1 +"},
		{"operator_with_no_operand", "+"},
		{"power_with_no_exponent", "2^"},
		{"term_with_no_right_operand", "2*"},
		{"error_inside_parentheses", "(1+)"},
		{"lone_decimal_point", "."},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := parser.NewExpressionParser().Parse(tc.expr)
			if err == nil {
				t.Fatalf("Parse(%q): expected error, got nil", tc.expr)
			}
		})
	}
}

// TestParse_IndependentStatePerCall guards against regressing the
// concurrency-safety fix: Parse must not share mutable state across calls.
func TestParse_IndependentStatePerCall(t *testing.T) {
	p := parser.NewExpressionParser()
	results := make(chan error, 2)

	go func() {
		_, err := p.Parse("1 + 2 + 3 + 4 + 5")
		results <- err
	}()
	go func() {
		_, err := p.Parse("(10 - 2) * 3")
		results <- err
	}()

	for i := 0; i < 2; i++ {
		if err := <-results; err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	}
}
