package logger

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// JSONFormatter is a logging JSON formatter.
type JSONFormatter struct{}

// Format satisfies the logger.Formatter interface.
func (f *JSONFormatter) Format(ctx Context) ([]byte, error) {
	b := bytes.NewBuffer(nil)

	data, err := json.Marshal(ctx)
	if err != nil {
		return nil, fmt.Errorf("cannot marshal JSON: %s", err)
	}

	b.Write(data)
	b.WriteString("\n")

	return b.Bytes(), nil
}
