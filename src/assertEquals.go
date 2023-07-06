package assertions

import (
	"reflect"
	"testing"
)

// AssertEquals checks if the expected and actual values are equal using deep equality comparison.
// If the values are not equal, it reports an error.
//
// Parameters:
//   - t: *testing.T - The testing.T instance for reporting the error.
//   - expected: interface{} - The expected value.
//   - actual: interface{} - The actual value.
func AssertEquals(t *testing.T, expected, actual interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected and actual are not equal. Expected: %v (type: %T), Actual: %v (type: %T)",
			expected, expected, actual, actual)
	}
}
