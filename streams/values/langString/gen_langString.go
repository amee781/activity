package langstring

import (
	"fmt"
	"sort"
)

// SerializeLangString converts a langString value to an interface representation
// suitable for marshalling into a text or binary format.
func SerializeLangString(this map[string]string) (interface{}, error) {
	return this, nil
}

// DeserializeLangString creates langString value from an interface representation
// that has been unmarshalled from a text or binary format.
func DeserializeLangString(this interface{}) (map[string]string, error) {
	if m, ok := this.(map[string]interface{}); ok {
		r := make(map[string]string)
		for k, v := range m {
			if s, ok := v.(string); ok {
				r[k] = s
			} else {
				return nil, fmt.Errorf("value %v cannot be interpreted as a string for rdf:langString", v)
			}
		}
		return r, nil
	} else {
		return nil, fmt.Errorf("%v cannot be interpreted as a map[string]interface{} for rdf:langString", this)
	}
}

// LessLangString returns true if the left langString value is less than the right
// value.
func LessLangString(lhs, rhs map[string]string) bool {
	var lk []string
	var rk []string
	for k := range lhs {
		lk = append(lk, k)
	}
	for k := range rhs {
		rk = append(rk, k)
	}
	sort.Sort(sort.StringSlice(lk))
	sort.Sort(sort.StringSlice(rk))
	for i := 0; i < len(lk) && i < len(rk); i++ {
		if lk[i] < rk[i] {
			return true
		} else if rk[i] < lk[i] {
			return false
		} else if lhs[lk[i]] < rhs[rk[i]] {
			return true
		} else if rhs[rk[i]] < lhs[lk[i]] {
			return false
		}
	}
	if len(lk) < len(rk) {
		return true
	} else {
		return false
	}
}
