package calculator

type CalculationRequest struct {
    Expression string `json:"expression"`
}

type CalculationResponse struct {
    Answer float64 `json:"answer"`
}