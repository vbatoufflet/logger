package logger

import (
	"os"
)

var root = &Logger{
	ctx:       make(Context),
	formatter: &TextFormatter{},
	level:     LevelInfo,
	output:    os.Stdout,
}
