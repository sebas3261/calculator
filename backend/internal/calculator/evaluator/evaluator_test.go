package evaluator_test

import (
	"math"
	"testing"

	"github.com/sebas3261/calculator/backend/internal/calculator"
	"github.com/sebas3261/calculator/backend/internal/calculator/evaluator"
)

func num(v float64) calculator.Node { return calculator.NumberNode{Value: v} }

func bin(l calculator.Node, op string, r calculator.Node) calculator.Node {
	return calculator.BinaryNode{Left: l, Operator: op, Right: r}
}

func un(op string, operand calculator.Node) calculator.Node {
	return calculator.UnaryNode{Operator: op, Operand: operand}
}

const epsilon = 1e-9

func TestEvaluate_ValidExpressions(t *testing.T) {
	cases := []struct {
		name string
		node calculator.Node
		want float64
	}{
		{"number", num(42), 42},
		{"unary_plus", un("+", num(5)), 5},
		{"unary_minus", un("-", num(5)), -5},
		{"square_root", un("√", num(9)), 3},
		{"square_root_of_zero", un("√", num(0)), 0},
		{"percent", un("%", num(50)), 0.5},
		{"addition", bin(num(2), "+", num(3)), 5},
		{"subtraction", bin(num(5), "-", num(3)), 2},
		{"multiplication", bin(num(4), "*", num(3)), 12},
		{"division", bin(num(10), "/", num(4)), 2.5},
		{"power", bin(num(2), "^", num(10)), 1024},
		{"power_fractional_exponent", bin(num(9), "^", num(0.5)), 3},
		{"power_negative_exponent", bin(num(2), "^", un("-", num(1))), 0.5},
		{"nested_precedence", bin(num(2), "+", bin(num(3), "*", num(4))), 14},
		{"nested_parentheses_equivalent", bin(bin(num(1), "+", num(2)), "*", num(3)), 9},
		{"double_negative", un("-", un("-", num(5))), 5},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := evaluator.NewEvaluator().Evaluate(tc.node)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if math.Abs(got-tc.want) > epsilon {
				t.Errorf("Evaluate() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestEvaluate_Errors(t *testing.T) {
	cases := []struct {
		name string
		node calculator.Node
	}{
		{"division_by_zero", bin(num(1), "/", num(0))},
		{"square_root_of_negative_number", un("√", num(-9))},
		{"overflow_to_infinity", bin(num(10), "^", num(400))},
		{"negative_base_fractional_exponent_is_not_real", bin(un("-", num(8)), "^", num(0.5))},
		{"zero_to_negative_power_is_infinity", bin(num(0), "^", un("-", num(1)))},
		{"unsupported_binary_operator", bin(num(1), "mod", num(2))},
		{"unsupported_unary_operator", un("!", num(1))},
		{"unsupported_node_type", calculator.FunctionNode{Name: "sqrt", Arguments: []calculator.Node{num(9)}}},
		{"error_in_left_operand_propagates_through_addition", bin(bin(num(1), "/", num(0)), "+", num(2))},
		{"error_in_right_operand_propagates_through_addition", bin(num(2), "+", bin(num(1), "/", num(0)))},
		{"error_in_operand_propagates_through_unary", un("-", bin(num(1), "/", num(0)))},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := evaluator.NewEvaluator().Evaluate(tc.node)
			if err == nil {
				t.Fatalf("expected error, got nil")
			}
		})
	}
}
