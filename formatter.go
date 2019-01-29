package logger

// Formatter is a logging formatter.
type Formatter interface {
	Format(Context) ([]byte, error)
}
