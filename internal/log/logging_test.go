package log

import (
	"errors"
	"io"
	"log/slog"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoggers(t *testing.T) {
	t.Run("Debug", func(t *testing.T) {
		r, w, _ := os.Pipe()
		logger = slog.New(slog.NewJSONHandler(w, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))
		Debug("TestDebug")
		w.Close()
		out, _ := io.ReadAll(r)
		assert.Contains(t, string(out), "TestDebug")
	})

	t.Run("Info", func(t *testing.T) {
		r, w, _ := os.Pipe()
		logger = slog.New(slog.NewJSONHandler(w, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}))
		Info("TestInfo")
		w.Close()
		out, _ := io.ReadAll(r)
		assert.Contains(t, string(out), "TestInfo")
	})

	t.Run("Warn", func(t *testing.T) {
		r, w, _ := os.Pipe()
		logger = slog.New(slog.NewJSONHandler(w, &slog.HandlerOptions{
			Level: slog.LevelWarn,
		}))
		Warn("TestWarn")
		w.Close()
		out, _ := io.ReadAll(r)
		assert.Contains(t, string(out), "TestWarn")
	})

	t.Run("Error", func(t *testing.T) {
		r, w, _ := os.Pipe()
		logger = slog.New(slog.NewJSONHandler(w, &slog.HandlerOptions{
			Level: slog.LevelError,
		}))
		Error("TestError", errors.New("err"))
		w.Close()
		out, _ := io.ReadAll(r)
		assert.Contains(t, string(out), "TestError")
	})
}
