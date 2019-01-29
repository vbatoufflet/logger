package logger

import (
	"os"
)

var root = &Logger{
	ctx:       make(Context),
	formatter: new(TextFormatter),
	level:     LevelInfo,
	output:    os.Stdout,
}
