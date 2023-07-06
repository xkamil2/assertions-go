package assertions

import (
	"fmt"
	"testing"
)

type aePerson struct {
	Name      string
	Age       int
	Address   aeAddress
	Languages []string
	Skills    map[string]int
}

type aeAddress struct {
	City string
}

type aePrivatePerson struct {
	name       string
	age        int
	HairColour string
}

func aeNewPrivatePerson(name string, age int) *aePrivatePerson {
	return &aePrivatePerson{
		name: name,
		age:  age,
	}
}

func TestAssertEquals(t *testing.T) {
	privatePerson1 := aeNewPrivatePerson("John", 2)
	privatePerson1.HairColour = "red"

	privatePerson2 := aeNewPrivatePerson("Jack", 66)
	privatePerson2.HairColour = "red"

	privatePerson3 := aeNewPrivatePerson("Jay", 56)
	privatePerson3.HairColour = "yellow"

	testCases := []struct {
		id       int
		expected interface{}
		actual   interface{}
		fails    bool
	}{
		{id: 1, expected: 42, actual: 42, fails: false},
		{id: 2, expected: "Hello", actual: "Hello", fails: false},
		{id: 3, expected: true, actual: true, fails: false},
		{id: 4, expected: []int{}, actual: []int{}, fails: false},
		{id: 5, expected: []int{1, 2, 3}, actual: []int{1, 2, 3}, fails: false},
		{id: 6, expected: map[string]int{}, actual: map[string]int{}, fails: false},
		{id: 7, expected: map[string]int{"a": 1, "b": 2}, actual: map[string]int{"a": 1, "b": 2}, fails: false},
		{id: 8, expected: &aePerson{Name: "Alice", Age: 25}, actual: &aePerson{Name: "Alice", Age: 25}, fails: false},

		{id: 9, expected: 42, actual: 43, fails: true},
		{id: 10, expected: "Hello", actual: "World", fails: true},
		{id: 11, expected: true, actual: false, fails: true},
		{id: 12, expected: []int{1, 2, 3}, actual: []int{3, 2, 1}, fails: true},
		{id: 13, expected: []int{1, 2, 3}, actual: []int{1, 2}, fails: true},
		{id: 14, expected: map[string]int{"a": 1, "b": 2}, actual: map[string]int{"a": 2, "b": 1}, fails: true},
		{id: 15, expected: &aePerson{Name: "Alice", Age: 25}, actual: &aePerson{Name: "Bob", Age: 30}, fails: true},

		{id: 16, expected: nil, actual: nil, fails: false},
		{id: 17, expected: nil, actual: 42, fails: true},
		{id: 18, expected: 42, actual: nil, fails: true},

		{id: 19, expected: aePerson{Name: "Alice", Age: 25}, actual: aePerson{Name: "Alice", Age: 25}, fails: false},
		{id: 20, expected: aePerson{Name: "Alice", Age: 25}, actual: aePerson{Name: "Alice", Age: 30}, fails: true},
		{id: 21, expected: aePerson{Name: "Alice", Age: 25, Address: aeAddress{City: "New York"}}, actual: aePerson{Name: "Alice", Age: 25, Address: aeAddress{City: "New York"}}, fails: false},
		{id: 22, expected: aePerson{Name: "Alice", Age: 25, Address: aeAddress{City: "New York"}}, actual: aePerson{Name: "Alice", Age: 25, Address: aeAddress{City: "Los Angeles"}}, fails: true},

		{id: 23, expected: aePerson{Name: "Alice", Age: 25, Languages: []string{"English", "Spanish"}}, actual: aePerson{Name: "Alice", Age: 25, Languages: []string{"English", "Spanish"}}, fails: false},
		{id: 24, expected: aePerson{Name: "Alice", Age: 25, Languages: []string{"English", "Spanish"}}, actual: aePerson{Name: "Alice", Age: 25, Languages: []string{"Spanish", "English"}}, fails: true},
		{id: 25, expected: aePerson{Name: "Alice", Age: 25, Skills: map[string]int{"Go": 5, "Java": 3}}, actual: aePerson{Name: "Alice", Age: 25, Skills: map[string]int{"Go": 5, "Java": 3}}, fails: false},
		{id: 26, expected: aePerson{Name: "Alice", Age: 25, Skills: map[string]int{"Go": 5, "Java": 3}}, actual: aePerson{Name: "Alice", Age: 25, Skills: map[string]int{"Java": 3, "Go": 5}}, fails: false},

		{id: 27, expected: aeNewPrivatePerson("Bob", 22), actual: aeNewPrivatePerson("Bob", 22), fails: false},
		{id: 28, expected: aeNewPrivatePerson("Kate", 53), actual: aeNewPrivatePerson("Bob", 25), fails: true},
		{id: 29, expected: privatePerson1, actual: privatePerson2, fails: true},

		{id: 30, expected: 42, actual: 43, fails: true},
		{id: 31, expected: "Hello", actual: "World", fails: true},
		{id: 32, expected: true, actual: false, fails: true},
		{id: 33, expected: []int{1, 2, 3}, actual: []int{3, 2, 1}, fails: true},
		{id: 34, expected: []int{1, 2, 3}, actual: []int{1, 2}, fails: true},
		{id: 35, expected: map[string]int{"a": 1, "b": 2}, actual: map[string]int{"a": 2, "b": 1}, fails: true},
		{id: 36, expected: &aePerson{Name: "Alice", Age: 25}, actual: &aePerson{Name: "Bob", Age: 30}, fails: true},
		{id: 37, expected: privatePerson2, actual: privatePerson3, fails: true},
	}

	for _, tc := range testCases {
		var message string

		if tc.fails {
			message = fmt.Sprintf("%d-Test case expected to fail", tc.id)
		} else {
			message = fmt.Sprintf("%d", tc.id)
		}

		t.Run(message, func(t *testing.T) {
			AssertEquals(t, tc.expected, tc.actual)
		})
	}
}
