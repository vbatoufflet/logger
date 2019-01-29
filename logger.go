package logger

import (
	"fmt"
	"io"
	"os"
	"time"
)

// Logger is logging handler instance.
type Logger struct {
	ctx       Context
	formatter Formatter
	level     int
	output    io.WriteCloser
}

// New creates a new logging handler instance.
func New(ctx ...Context) *Logger {
	if ctx == nil {
		return root
	}
	return root.New(ctx...)
}

// Close closes the logging handler output.
func (l *Logger) Close() error {
	return l.output.Close()
}

// New creates a new logging handler instance given a context.
func (l *Logger) New(ctx ...Context) *Logger {
	return &Logger{
		ctx:       l.ctx.Union(ctx...),
		formatter: l.formatter,
		output:    l.output,
	}
}

// SetFormatter sets the logging message formatter.
func (l *Logger) SetFormatter(formatter Formatter) *Logger {
	l.formatter = formatter
	return l
}

// SetLevel sets the logging handler output severity maximum level.
func (l *Logger) SetLevel(level int) *Logger {
	l.level = level
	return l
}

// SetOutput sets the logging handler output.
func (l *Logger) SetOutput(output io.WriteCloser) *Logger {
	l.output = output
	return l
}

// Debug sends a debug message to the logging handler output.
func (l *Logger) Debug(format string, args ...interface{}) *Logger {
	l.write(l.ctx.Union(Context{"level": LevelDebug}), format, args...)
	return l
}

// Error sends an error message to the logging handler output.
func (l *Logger) Error(format string, args ...interface{}) *Logger {
	l.write(l.ctx.Union(Context{"level": LevelError}), format, args...)
	return l
}

// Info sends an info message to the logging handler output.
func (l *Logger) Info(format string, args ...interface{}) *Logger {
	l.write(l.ctx.Union(Context{"level": LevelInfo}), format, args...)
	return l
}

// Notice sends a notice message to the logging handler output.
func (l *Logger) Notice(format string, args ...interface{}) *Logger {
	l.write(l.ctx.Union(Context{"level": LevelNotice}), format, args...)
	return l
}

// Warning sends a warning message to the logging handler output.
func (l *Logger) Warning(format string, args ...interface{}) *Logger {
	l.write(l.ctx.Union(Context{"level": LevelWarning}), format, args...)
	return l
}

func (l *Logger) write(ctx Context, format string, args ...interface{}) {
	level := ctx.Get("level").(int)
	if l.level > 0 && level > l.level || l.level == 0 && level > root.level {
		return
	}

	b, err := l.formatter.Format(ctx.Union(Context{
		"level":   levelMap[level],
		"message": fmt.Sprintf(format, args...),
		"time":    time.Now(),
	}))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: failed to format message: %s", err)
		return
	}

	l.output.Write(b)
}
