package helpers

import (
	"log/slog"
	"os"
)

// FatalError wraps call to slog.Error and issues os.Exit(1)
func FatalError(msg string, args ...any) {
	slog.Error(msg, args...)
	os.Exit(1)
}
