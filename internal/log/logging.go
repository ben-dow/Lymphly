package log

import (
	"context"
	"log/slog"
	"os"
	"runtime"
	"strconv"
)

var logger *slog.Logger

// creates global logger
func init() {
	logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
}

/// base logger 
func base(msg string, level slog.Level, args ...any) {
	_, file, line, _ := runtime.Caller(2)
	logger.Log(context.Background(), level, msg, append(args, "logger", file+":"+strconv.FormatInt(int64(line), 10))...)
}

// Debug creates a log entry at the DEBUG level
func Debug(msg string, args ...any) {
	base(msg, slog.LevelDebug, args...)
}

// Info creates a log enry at the INFO level
func Info(msg string, args ...any) {
	base(msg, slog.LevelInfo, args...)
}

// Warn creates a log enry at the WARN level
func Warn(msg string, args ...any) {
	base(msg, slog.LevelWarn, args...)
}

// Error creates a log enry at the ERROR level
func Error(msg string, err error, args ...any) {
	base(msg, slog.LevelError, append(args, "err", err.Error())...)
}
