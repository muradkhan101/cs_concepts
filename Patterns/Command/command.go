package main

// Command and corresponding arguments
type Command struct {
	fn   func(args []string)
	args []string
}

// Execute command function
func (c *Command) Execute() {
	if c.fn != nil {
		c.fn(c.args)
	}
}

// SetArguments - read name
func (c *Command) SetArguments(newArgs []string) {
	c.args = newArgs
}
