package mux

// checkPairs returns the count of strings passed in, and an error if
import (
	"fmt"
	"net/http"
	"strings"
)

// the count is not an even number.
func isEvenPairs(pairs ...string) (int, error) {
	length := len(pairs)
	if length%2 != 0 {
		return length, fmt.Errorf("mux: number of parameters must be multiple of 2, got %v", pairs)
	}
	return length, nil
}

// convertStringsToMap converts variadic string parameters to a
// string to string map.
func convertStringsToMap(iep func(pairs ...string) (int, error), pairs ...string) (map[string]string, error) {
	length, err := iep(pairs...)
	if err != nil {
		return nil, err
	}
	m := make(map[string]string, length/2)
	for i := 0; i < length; i += 2 {
		m[pairs[i]] = pairs[i+1]
	}
	return m, nil
}

// matchMapWithString returns true if the given key/value pairs exist in a given map.
func matchMapWithString(toCheck map[string]string, toMatch map[string][]string, canonicalKey bool) bool {
	for k, v := range toCheck {
		// Check if key exists.
		if canonicalKey {
			k = http.CanonicalHeaderKey(k)
		}

		values := toMatch[k]

		if values == nil {
			return false
		}

		if v != "" {
			// If value was defined as an empty string we only check that the
			// key exists. Otherwise we also check for equality.
			valueExists := false
			for _, value := range values {
				if v == value {
					valueExists = true
					break
				}
			}
			if !valueExists {
				return false
			}
		}
	}

	return true
}

// containsRegexPath returns true if the path a regex path
func containsRegex(path string) bool {
	return strings.Contains(path, "#")
}

// containsRegexPath returns true if the path contains vars
func containsVars(path string) bool {
	return strings.Contains(path, ":")
}
