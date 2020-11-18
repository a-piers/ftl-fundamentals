package calculator_test

import (
	"calculator"
	"fmt"
	"math"
	"math/rand"
	"testing"
)

func closeEnough(want, got, tolerance float64) bool {
	return math.Abs(want-got) <= tolerance
}

func TestExpression(t *testing.T) {
	t.Parallel()
	type testCase struct {
		name       string
		expression string
		want       float64
	}

	testCases := []testCase{
		{name: "Multiply two numbers", expression: "8*8", want: 64},
		{name: "Add to numbers", expression: "2 + 2", want: 4},
		{name: "Divide two numbers", expression: "10.0	/	2", want: 5},
		{name: "Subtract numbers", expression: "10.5 - 5.5", want: 5.0},
	}

	for _, tc := range testCases {
		got, err := calculator.Expression(tc.expression)
		if err != nil {
			t.Errorf("Error occured in expression calculation: %s", err)
		}
		if !closeEnough(tc.want, got, 0.00001) {
			t.Errorf("%s - Expression(%s): want %f, got %f", tc.name, tc.expression, tc.want, got)
		}
	}
}

func ExampleAdd() {
	fmt.Println(calculator.Add(2, 3))
	// Output: 5
}

// func TestAdd(t *testing.T) {
// 	t.Parallel()
// 	var want float64 = 4
// 	got := calculator.Add(2, 2)
// 	if want != got {
// 		t.Errorf("want %f, got %f", want, got)
// 	}
// }

func TestAdd(t *testing.T) {
	t.Parallel()
	type testCase struct {
		name string
		a, b float64
		want float64
	}

	testCases := []testCase{
		{name: "Add two possitive numbers", a: 2, b: 2, want: 4},
		{name: "Add two possitive numbers", a: 1, b: 1, want: 2},
		{name: "Add two possitive numbers", a: 5, b: 0, want: 5},
		{name: "Add two possitive numbers", a: 0, b: 0, want: 0},
		{name: "Add two possitive floats", a: 1.1, b: 1.2, want: 2.3},
	}

	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("%s - Add(%f, %f): want %f, got %f", tc.name, tc.a, tc.b, tc.want, got)
		}
	}
}

func TestAddRandom(t *testing.T) {
	t.Parallel()
	type testCase struct {
		name string
		a, b float64
		want float64
	}
	testCases := []testCase{}
	for i := 0; i < 100; i++ {
		a := rand.Float64()
		b := rand.Float64()
		want := a + b
		testCases = append(testCases, testCase{
			name: fmt.Sprintf("Add(%f, %f)", a, b),
			a:    a,
			b:    b,
			want: want,
		})
	}

	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		// Floats are hard to compare, therefore this is a possible method to still compare these values
		if !closeEnough(tc.want, got, 0.000001) {
			t.Errorf("%s - Add(%f, %f): want %f, got %f", tc.name, tc.a, tc.b, tc.want, got)
		}
		// Inital method of comparing floats
		// if tc.want != got {
		// 	t.Errorf("%s - Add(%f, %f): want %f, got %f", tc.name, tc.a, tc.b, tc.want, got)
		// }
	}
}

func TestAddMany(t *testing.T) {
	t.Parallel()
	type testCase struct {
		name string
		a    []float64
		want float64
	}
	testCases := []testCase{
		{name: "Add 4 numbers", a: []float64{1, 2, 3, 4}, want: 10},
		{name: "Add 1 number", a: []float64{10}, want: 10},
		{name: "Add 2 numbers", a: []float64{199, 1}, want: 200},
		{name: "Add 10 numbers", a: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}, want: 45},
	}
	for _, tc := range testCases {
		got := calculator.AddMany(tc.a...)
		if tc.want != got {
			t.Errorf("%s - Add(%v): want %f, got %f", tc.name, tc.a, tc.want, got)
		}
	}
}

func ExampleSubtract() {
	fmt.Println(calculator.Subtract(10, 5))
	// Output: 5
}

// func TestSubtract(t *testing.T) {
// 	t.Parallel()
// 	var want float64 = 2
// 	got := calculator.Subtract(4, 2)
// 	if want != got {
// 		t.Errorf("want %f, got %f", want, got)
// 	}
// }

func TestSubtract(t *testing.T) {
	t.Parallel()
	type testCase struct {
		name string
		a, b float64
		want float64
	}

	testCases := []testCase{
		{name: "Substract two possitive number into negative result", a: 0, b: 1, want: -1},
		{name: "Substract the first number in half", a: 10, b: 5, want: 5},
		{name: "Substract 10 from 0", a: 0, b: 10, want: -10},
		{name: "Substract using two negative numebrs", a: -5, b: -5, want: 0},
		{name: "Substract using two floats", a: 5.5, b: 2.1, want: 3.4},
	}

	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("%s - Subtract(%f, %f): want %f, got %f", tc.name, tc.a, tc.b, tc.want, got)
		}
	}
}

func TestSubtractMany(t *testing.T) {
	t.Parallel()
	type testCase struct {
		name string
		a    []float64
		want float64
	}

	testCases := []testCase{
		{name: "Substract 4 possitive number into negative result", a: []float64{10, 8, 6, 4}, want: -8},
		{name: "Substract possitive numbers", a: []float64{10, 1, 2, 3}, want: 4},
		{name: "Substract 10 with zeros", a: []float64{10, 0, 0, 0, 0, 0}, want: 10},
		{name: "Substract mulitple floats", a: []float64{11.101012, 0.928464, 1.12345678}, want: 9.04909122},
	}

	for _, tc := range testCases {
		got := calculator.SubtractMany(tc.a...)
		if !closeEnough(tc.want, got, 0.00001) {
			t.Errorf("%s - Subtract(%v): want %f, got %f", tc.name, tc.a, tc.want, got)
		}
	}
}

func ExampleMultiply() {
	fmt.Println(calculator.Multiply(2, 5))
	// Output: 10
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	type testCase struct {
		name string
		a, b float64
		want float64
	}

	testCases := []testCase{
		{name: "Multiply zero with one", a: 0, b: 1, want: 0},
		{name: "Multiply two possitive numbers", a: 10, b: 5, want: 50},
		{name: "Multiply two negative numbers", a: -5, b: -5, want: 25},
		{name: "Multiply one negative and one possitive number", a: -100, b: 20, want: -2000},
		{name: "Multiply two floats", a: 2.5, b: 2.5, want: 6.25},
	}

	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("%s - Multiply(%f, %f): want %f, got %f", tc.name, tc.a, tc.b, tc.want, got)
		}
	}
}

func TestMultiplyMany(t *testing.T) {
	t.Parallel()
	type testCase struct {
		name string
		a    []float64
		want float64
	}

	testCases := []testCase{
		{name: "Multiply zeros with one", a: []float64{0.0, 0, 0, 0, 0, 1}, want: 0},
		{name: "Multiply possitive numbers", a: []float64{10, 5, 2, 1}, want: 100},
		{name: "Multiply negative numbers", a: []float64{-5, -5, -2, -1}, want: 50},
		{name: "Multiply floats", a: []float64{2.5, 2.5, 5.0386252, 1.1028464}, want: 34.730185},
	}

	for _, tc := range testCases {
		got := calculator.MultiplyMany(tc.a...)
		if !closeEnough(tc.want, got, 0.00001) {
			t.Errorf("%s - Multiply(%v): want %f, got %f", tc.name, tc.a, tc.want, got)
		}
	}
}

func ExampleDivide() {
	fmt.Println(calculator.Divide(10, 5))
	// Output: 2
}

func TestDivide(t *testing.T) {
	t.Parallel()
	type testCase struct {
		name        string
		a, b        float64
		want        float64
		errExpected bool
	}

	testCases := []testCase{
		{name: "Divide two possitive numbers", a: 10, b: 2, want: 5, errExpected: false},
		{name: "Divide two equeal numbers", a: 1, b: 1, want: 1, errExpected: false},
		{name: "Divide by zero", a: 10, b: 0, want: 999, errExpected: true},
		{name: "Divide two negative numbers", a: -1, b: -2, want: 0.5, errExpected: false},
		{name: "Divide two possive numbers", a: 100, b: 10, want: 10, errExpected: false},
		{name: "Divide one negative and one possive number", a: -100, b: 10, want: -10, errExpected: false},
		{name: "Divide two negative numbers", a: -5, b: -5, want: 1, errExpected: false},
		{name: "Divide one float and one int", a: 20.5, b: 5, want: 4.1, errExpected: false},
	}

	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)
		errReceived := err != nil
		if tc.errExpected != errReceived {
			t.Fatalf("%s - Divide(%f, %f): unexpected error status: %v", tc.name, tc.a, tc.b, errReceived)
		}
		if !tc.errExpected && tc.want != got {
			t.Errorf("%s - Divide(%f, %f): want %f, got %f", tc.name, tc.a, tc.b, tc.want, got)
		}
	}
}

func TestDivideMany(t *testing.T) {
	t.Parallel()
	type testCase struct {
		name        string
		a           []float64
		want        float64
		errExpected bool
	}

	testCases := []testCase{
		{name: "Divide possitive numbers", a: []float64{10, 2, 1.0}, want: 5, errExpected: false},
		{name: "Divide equeal numbers", a: []float64{1, 1.0, 1, 1.0}, want: 1, errExpected: false},
		{name: "Divide by zeros", a: []float64{10, 0, 0.0, 0.0, 0, 0}, want: 999, errExpected: true},
		{name: "Divide by negative numbers", a: []float64{-1, -2, -10, -5.2}, want: 0.009615, errExpected: false},
		{name: "Divide one float and one int", a: []float64{20.5, 5, 1.0, 1}, want: 4.1, errExpected: false},
	}

	for _, tc := range testCases {
		got, err := calculator.DivideMany(tc.a...)
		errReceived := err != nil
		if tc.errExpected != errReceived {
			t.Fatalf("%s - Divide(%v): unexpected error status: %v", tc.name, tc.a, errReceived)
		}
		if !tc.errExpected && !closeEnough(tc.want, got, 0.000001) {
			t.Errorf("%s - Divide(%v): want %f, got %f", tc.name, tc.a, tc.want, got)
		}
	}
}

func ExampleSqrt() {
	fmt.Println(calculator.Sqrt(100.0))
	// Output: 10
}

func TestSqrt(t *testing.T) {
	t.Parallel()
	type testCase struct {
		name        string
		a           float64
		want        float64
		errExpected bool
	}
	testCases := []testCase{
		{name: "Sqrt of negative number", a: -11.101, want: 0, errExpected: true},
		{name: "Sqrt of zero", a: 0, want: 0, errExpected: false},
		{name: "Sqrt of positive number", a: 100, want: 10, errExpected: false},
		{name: "Sqrt of small number", a: 0.00001, want: 0.003162278, errExpected: false},
		{name: "Sqrt of big number", a: 9999999999, want: 99999.999995, errExpected: false},
	}
	for _, tc := range testCases {
		got, err := calculator.Sqrt(tc.a)
		errReceived := err != nil
		if tc.errExpected != errReceived {
			t.Fatalf("%s - Sqrt(%f): unexpected error status: %v", tc.name, tc.a, errReceived)
		}
		if !tc.errExpected && !closeEnough(tc.want, got, 0.00001) {
			t.Errorf("%s - Sqrt(%f): want %f, got %f", tc.name, tc.a, tc.want, got)
		}
	}
}
