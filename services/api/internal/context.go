package internal

import (
	"go.uber.org/zap"
)

// Context is the internal structure
// of this API
type Context struct {
	Logger *zap.Logger
	Env    map[string]string
}
