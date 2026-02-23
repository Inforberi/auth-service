package logger

import (
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/inforberi/auth-service/internal/config"
)

func NewLogger(cfg *config.Config) *slog.Logger {
	level := parseLevel(cfg.Logger.Level)

	opts := &slog.HandlerOptions{
		Level:     level,
		AddSource: cfg.AppEnv != "prod",
	}

	if cfg.AppEnv != "prod" {
		opts.ReplaceAttr = devReplaceAttr
	}

	w := os.Stdout
	h := chooseLogHandler(cfg.Logger.Format, w, opts, cfg)

	log := slog.New(h).With("env", cfg.AppEnv)

	return log
}

func parseLevel(level string) slog.Level {
	lvl := strings.ToLower(strings.TrimSpace(level))

	switch lvl {
	case "debug":
		return slog.LevelDebug
	case "warn", "warning":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	case "info":
		return slog.LevelInfo
	default:
		return slog.LevelInfo
	}
}

func chooseLogHandler(format string, w io.Writer, opts *slog.HandlerOptions, cfg *config.Config) slog.Handler {
	fmt := strings.ToLower(strings.TrimSpace(format))

	switch fmt {
	case "json":
		return slog.NewJSONHandler(w, opts)
	case "text":
		return slog.NewTextHandler(w, opts)
	default:
		if cfg.AppEnv == "prod" {
			return slog.NewJSONHandler(w, opts)
		}
		return slog.NewTextHandler(w, opts)
	}
}

func devReplaceAttr(_ []string, a slog.Attr) slog.Attr {
	// Время: HH:MM:SS
	if a.Key == slog.TimeKey {
		t := a.Value.Time()
		return slog.String(slog.TimeKey, t.Format("15:04:05"))
	}

	// Source: только имя файла + строка
	if a.Key == slog.SourceKey {
		src, ok := a.Value.Any().(*slog.Source)
		if ok && src != nil {
			return slog.String(slog.SourceKey, filepath.Base(src.File)+":"+strconv.Itoa(src.Line))
		}
	}

	return a
}
