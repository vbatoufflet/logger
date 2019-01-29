package logger

const (
	_ = iota
	// LevelError is the error logging level.
	LevelError
	// LevelWarning is the warning logging level.
	LevelWarning
	// LevelNotice is the notice logging level.
	LevelNotice
	// LevelInfo is the info logging level.
	LevelInfo
	// LevelDebug is the debug logging level.
	LevelDebug
)

var levelMap = map[int]string{
	LevelError:   "error",
	LevelWarning: "warning",
	LevelNotice:  "notice",
	LevelInfo:    "info",
	LevelDebug:   "debug",
}

// Level returns a logging level given its textual key.
func Level(key string) int {
	for k, v := range levelMap {
		if v == key {
			return k
		}
	}

	// Return default level
	return LevelInfo
}
