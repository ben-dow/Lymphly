package log

import (
	"context"
	"log/slog"
	"os"
	"runtime"
	"strconv"
)

var logger *slog.Logger

func init() {
	logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
}

func base(msg string, level slog.Level, args ...any) {
	_, file, line, _ := runtime.Caller(2)
	logger.Log(context.Background(), level, msg, append(args, "logger", file+":"+strconv.FormatInt(int64(line), 10))...)
}

func Debug(msg string, args ...any) {
	base(msg, slog.LevelDebug, args...)
}

func Info(msg string, args ...any) {
	base(msg, slog.LevelInfo, args...)
}

func Warn(msg string, args ...any) {
	base(msg, slog.LevelWarn, args...)
}

func Error(msg string, err error, args ...any) {
	base(msg, slog.LevelError, append(args, "err", err.Error())...)
}
