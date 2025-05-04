package context

import (
	"context"
	"time"
)

func TimeoutCtx() (context.Context, context.CancelFunc){
	return context.WithTimeout(context.Background(), 5*time.Second)
}