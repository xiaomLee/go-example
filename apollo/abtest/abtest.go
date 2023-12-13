package abtest

import (
	"apollo"
	"context"
)

// Allow is the key allowed to be through
// ctx is the context.Context with common key-pairs e.g.:
// phone: 123xxx
// name: superAdmin
// companyId: 0
func Allow(ctx context.Context, key string) bool {
	exprStr := apollo.GetKey("abtest", key).String()
	expr, err := ParseExpr(exprStr)
	if err != nil {
		return false
	}
	return expr.Execute(ctx)
}
