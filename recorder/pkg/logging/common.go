package logging

import "context"

// Retrieves a string from the provided context with the given key or the default def provided if not present
func RetrieveStringFromCtx(ctx context.Context, key any, def string) string {
	if val, ok := ctx.Value(key).(string); ok {
		return val
	}

	return def
}
