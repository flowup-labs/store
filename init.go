package store

var (
	// C is the global instance of client
	c Client
)

// C returns the global instance of the client
func C() Client {
	return c
}

// Initialize sets the global access client
// Warning: this will be effective only during init!
func Initialize(x Client) {
	c = x
}
