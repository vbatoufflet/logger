package logger

// Context is a logging message context map.
type Context map[string]interface{}

// Get returns the value of a context field.
func (c Context) Get(key string) interface{} {
	v, _ := c[key]
	return v
}

// Pop returns the value of a context field removing it from the map.
func (c Context) Pop(key string) interface{} {
	v, ok := c[key]
	if ok {
		delete(c, key)
	}
	return v
}

// Union returns a new context map instance representing the union of both current and given ones.
func (c Context) Union(ctx ...Context) Context {
	out := make(Context)
	for k, v := range c {
		out[k] = v
	}
	for _, sc := range ctx {
		for k, v := range sc {
			out[k] = v
		}
	}
	return out
}
