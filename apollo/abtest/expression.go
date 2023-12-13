package abtest

import "context"

type Expression struct{}

func ParseExpr(expr string) (*Expression, error) {
	return &Expression{}, nil
}

func (e *Expression) Execute(ctx context.Context) bool {
	return false
}
