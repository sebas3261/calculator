package calculator

type Node interface {
    node()
}

type NumberNode struct {
    Value float64
}

func (NumberNode) node() {}

type BinaryNode struct {
    Left     Node
    Operator string
    Right    Node
}

func (BinaryNode) node() {}

type UnaryNode struct {
    Operator string
    Operand  Node
}

func (UnaryNode) node() {}

type FunctionNode struct {
    Name      string
    Arguments []Node
}

func (FunctionNode) node() {}