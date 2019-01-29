package logger

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
	"time"
)

var colors = map[string]int{
	"error":   31,
	"warning": 33,
	"notice":  35,
	"info":    34,
	"debug":   36,
}

// TextFormatter is a logging text formatter.
type TextFormatter struct {
	DisableColors bool
}

// Format satisfies the logger.Formatter interface.
func (f *TextFormatter) Format(ctx Context) ([]byte, error) {
	t := ctx.Pop("time").(time.Time)
	level := ctx.Pop("level").(string)
	message := ctx.Pop("message").(string)

	b := bytes.NewBuffer(nil)

	b.WriteString(t.Format(time.RFC3339))

	b.WriteString(" ")

	part := fmt.Sprintf("%7s", strings.ToUpper(level))
	if !f.DisableColors {
		code, _ := colors[level]
		b.WriteString(color(part, code))
	} else {
		b.WriteString(part + ":")
	}

	b.WriteString(" ")

	b.WriteString(message)

	if len(ctx) > 0 {
		var (
			keys  []string
			parts []string
		)

		for k := range ctx {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		for _, k := range keys {
			part = k + "="
			if _, ok := ctx[k].(string); ok {
				part += fmt.Sprintf("%q", ctx[k])
			} else {
				part += fmt.Sprintf("%v", ctx[k])
			}
			parts = append(parts, part)
		}

		if !f.DisableColors {
			b.WriteString(" " + color(strings.Join(parts, " "), 90))
		} else {
			b.WriteString(" [" + strings.Join(parts, " ") + "]")
		}
	}

	b.WriteString("\n")

	return b.Bytes(), nil
}

func color(text string, code int) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", code, text)
}
