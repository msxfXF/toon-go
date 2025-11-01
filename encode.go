package toon

// EncodeOptions configures the encoding behavior
type EncodeOptions struct {
	// Indent specifies the number of spaces per indentation level (default: 2)
	Indent int
	// Delimiter specifies the delimiter for array values: ',' (comma), '\t' (tab), '|' (pipe)
	Delimiter rune
	// LengthMarker adds an optional '#' prefix to array lengths (e.g., items[#3])
	LengthMarker bool
}

// DefaultEncodeOptions returns the default encoding options
func DefaultEncodeOptions() EncodeOptions {
	return EncodeOptions{
		Indent:       2,
		Delimiter:    ',',
		LengthMarker: false,
	}
}

// Encode converts a Go value to TOON format
//
// Parameters:
//   - value: Any Go value (map, slice, primitive, or nested structure)
//   - options: Optional encoding configuration (use nil for defaults)
//
// Returns:
//   - A TOON-formatted string with no trailing newline
//   - An error if encoding fails
func Encode(value interface{}, options *EncodeOptions) (string, error) {
	panic("toon.Encode: not yet implemented - this is a namespace reservation. See https://github.com/toon-format/toon-go")
}
