package calculator

type Parser interface {
    Parse(expression string) (Node, error)
}

type Evaluator interface {
    Evaluate(node Node) (float64, error)
}