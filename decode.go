package toon

// DecodeOptions configures the decoding behavior
type DecodeOptions struct {
	// Indent specifies the expected number of spaces per indentation level (default: 2)
	Indent int
	// Strict enables strict validation (default: true)
	Strict bool
}

// DefaultDecodeOptions returns the default decoding options
func DefaultDecodeOptions() DecodeOptions {
	return DecodeOptions{
		Indent: 2,
		Strict: true,
	}
}

// Decode converts a TOON-formatted string to a Go value
//
// Parameters:
//   - input: A TOON-formatted string to parse
//   - options: Optional decoding configuration (use nil for defaults)
//
// Returns:
//   - A Go value (map, slice, or primitive) representing the parsed TOON data
//   - An error if parsing fails
func Decode(input string, options *DecodeOptions) (interface{}, error) {
	panic("toon.Decode: not yet implemented - this is a namespace reservation. See https://github.com/toon-format/toon-go")
}
