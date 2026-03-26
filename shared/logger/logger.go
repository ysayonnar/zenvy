package logger

import (
	"io"
	"log/slog"
)

// New returns *slog.Logger with the specified write destination
func New(dest io.Writer, isDebug bool) *slog.Logger {
	if isDebug {
		h := &colorHandler{out: dest, level: slog.LevelDebug, addSource: true}
		return slog.New(h)
	}

	opts := slog.HandlerOptions{Level: slog.LevelInfo, AddSource: true}
	return slog.New(slog.NewJSONHandler(dest, &opts))
}

// Err returns attribute error with provided error message
func Err(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}
