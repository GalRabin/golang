// godoc -http=:8080
// go test
// go test -bench .
// don’t forget the “.” in the line above
// go test -cover
// go test -coverprofile c.out
// go tool cover -html=c.out

package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	type test struct {
		input int
		expected int
	}

	tests := []test{
		{
			input: 1,
			expected: 7,
		},
		{
			input: 2,
			expected: 14,
		},
	}

	for _, test := range tests {
		output := Years(test.input)
		if output != test.expected {
			t.Errorf("Expected value %v, But got %v", test.expected, output)
		}
	}
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(8)
	}
}

func ExampleYears() {
	fmt.Println(Years(8))
	// output:
	// 56
}


func TestYearsTwo(t *testing.T) {
	type test struct {
		input int
		expected int
	}

	tests := []test{
		{
			input: 1,
			expected: 7,
		},
		{
			input: 2,
			expected: 14,
		},
	}

	for _, test := range tests {
		output := YearsTwo(test.input)
		if output != test.expected {
			t.Errorf("Expected value %v, But got %v", test.expected, output)
		}
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(8)
	}
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(8))
	// output:
	// 56
}