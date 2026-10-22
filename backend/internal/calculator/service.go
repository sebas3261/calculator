package calculator

import (
    "errors"
    "strings"
)

type Service struct {
    parser    Parser
    evaluator Evaluator
}

func NewService(parser Parser, evaluator Evaluator) *Service {
    return &Service{
        parser:    parser,
        evaluator: evaluator,
    }
}

func (s *Service) Calculate(request CalculationRequest) (CalculationResponse, error) {
    if strings.TrimSpace(request.Expression) == "" {
        return CalculationResponse{}, errors.New("expression cannot be empty")
    }

    ast, err := s.parser.Parse(request.Expression)
    if err != nil {
        return CalculationResponse{}, err
    }

    result, err := s.evaluator.Evaluate(ast)
    if err != nil {
        return CalculationResponse{}, err
    }

    return CalculationResponse{Answer: result}, nil
}