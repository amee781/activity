package rfc2045

import "fmt"

// SerializeRfc2045 converts a rfc2045 value to an interface representation
// suitable for marshalling into a text or binary format.
func SerializeRfc2045(this string) (interface{}, error) {
	return this, nil
}

// DeserializeRfc2045 creates rfc2045 value from an interface representation that
// has been unmarshalled from a text or binary format.
func DeserializeRfc2045(this interface{}) (string, error) {
	if s, ok := this.(string); ok {
		return s, nil
	} else {
		return "", fmt.Errorf("%v cannot be interpreted as a string for MIME media type", this)
	}
}

// LessRfc2045 returns true if the left rfc2045 value is less than the right value.
func LessRfc2045(lhs, rhs string) bool {
	return lhs < rhs
}
