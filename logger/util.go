package logger

import (
	"QCaller/types"
	"fmt"
)

// ContextMsg : util method for including request id in log
func ContextMsg(ctx types.Context, msg string) string {
	return fmt.Sprintf("[ %v ] - %v", ctx.GetReqID(), msg)
}
