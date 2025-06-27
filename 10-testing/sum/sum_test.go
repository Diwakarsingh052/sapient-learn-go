package sum

import (
	"github.com/stretchr/testify/require"
	"testing"
)

type args struct {
	input []int
}

// go test ./... // run all the tests for the project

// function names must start with word Test to signal it is a test
// file name must have _test.go to register a file as a test
func TestSumInt(t *testing.T) {
	// Figure out two things
	// What are inputs (parameters)
	// Expected output

	input := []int{1, 2, 3, 4, 5}
	want := 15
	got := SumInt(input)
	if got != want {
		// test would continue on if test case fail
		t.Errorf("sum of 1 to 5 should be %v; got %v", want, got)

		// Uncomment next line to stop the test if it fails at this point.
		//t.Fatalf("sum of 1 to 5 should be %v; got %v", want, got)
	}

	want = 0
	got = SumInt(nil)

	if got != want {
		t.Errorf("sum of nil should be %v; got %v", want, got)
	}
}

func TestSumIntV2(t *testing.T) {
	//table tests

	//columns , rows

	tt := []struct {
		name  string
		input []int
		want  int
	}{
		// each index in slice would represent one test case
		{
			name:  "one to five numbers",
			input: []int{1, 2, 3, 4, 5},
			want:  16,
		},
		{
			name:  "empty slice",
			input: nil,
			want:  0,
		},
	}

	for _, tc := range tt {

		// t.Run creates a subtest, you can run each index of slice as a subtest
		t.Run(tc.name, func(t *testing.T) {
			got := SumInt(tc.input)
			require.Equal(t, tc.want, got)
		})

	}

}
