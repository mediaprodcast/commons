package executor

import (
	"fmt"
	"strings"
)

// ArgsBuilder struct to build command-line arguments.
type ArgsBuilder struct {
	args []string
}

// NewArgsBuilder creates a new ArgsBuilder instance.
func NewArgsBuilder() *ArgsBuilder {
	return &ArgsBuilder{
		args: make([]string, 0),
	}
}

// Withf adds a formatted argument using fmt.Sprintf.
func (b *ArgsBuilder) Withf(format string, a ...interface{}) *ArgsBuilder {
	b.args = append(b.args, fmt.Sprintf(format, a...))
	return b
}

// With appends multiple arguments.
func (b *ArgsBuilder) With(args ...string) *ArgsBuilder {
	b.args = append(b.args, args...)
	return b
}

// WithIf adds an argument only if the condition is true.
func (b *ArgsBuilder) WithIf(condition bool, arg string) *ArgsBuilder {
	if condition {
		b.args = append(b.args, arg)
	}
	return b
}

// AddMap adds arguments based on a map.  Keys become flags (prefixed with "-"), values become the argument.
// If value is empty string, only the key is added as a flag.
func (b *ArgsBuilder) WithMap(m map[string]string) *ArgsBuilder {
	for k, v := range m {
		b.args = append(b.args, "-"+k)
		if v != "" {
			b.args = append(b.args, v)
		}
	}
	return b
}

// AddSlice adds arguments from a string slice.
func (b *ArgsBuilder) AddSlice(s []string) *ArgsBuilder {
	b.args = append(b.args, s...)
	return b
}

// String returns the arguments as a space-separated string.  Useful for printing or debugging.
func (b *ArgsBuilder) String() string {
	return strings.Join(b.args, " ")
}

// Build returns the arguments as a string slice.  This is what you'd typically pass to exec.Command.
func (b *ArgsBuilder) Build() []string {
	return b.args
}
