package rfc5988

import "fmt"

// SerializeRfc5988 converts a rfc5988 value to an interface representation
// suitable for marshalling into a text or binary format.
func SerializeRfc5988(this string) (interface{}, error) {
	return this, nil
}

// DeserializeRfc5988 creates rfc5988 value from an interface representation that
// has been unmarshalled from a text or binary format.
func DeserializeRfc5988(this interface{}) (string, error) {
	if s, ok := this.(string); ok {
		return s, nil
	} else {
		return "", fmt.Errorf("%v cannot be interpreted as a string for rel", this)
	}
}

// LessRfc5988 returns true if the left rfc5988 value is less than the right value.
func LessRfc5988(lhs, rhs string) bool {
	return lhs < rhs
}
