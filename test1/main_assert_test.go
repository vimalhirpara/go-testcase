package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type TestStruct struct {
	arg1, arg2, expected int
}

var PositiveStructs = []TestStruct{
	{2, 1, 2},
	{4, 3, 12},
	{6, 5, 30},
	{8, 7, 56},
	{10, 9, 90},
}

var NegativeStructs = []TestStruct{
	{2, 1, 1},
	{4, 3, 15},
	{6, 5, 50},
	{8, 7, 72},
	{10, 9, 80},
}

// Test Case Without Any Package
func TestMultiplyWithoutPackage(t *testing.T) {

	mergedStruct := append(PositiveStructs, NegativeStructs...)

	for _, val := range mergedStruct {

		result, _ := Multiply(val.arg1, val.arg2)

		if val.expected == result {
			t.Logf("%d x %d = %d (Pass)", val.arg1, val.arg2, val.expected)
		} else {
			t.Logf("%d x %d = %d (Fail)", val.arg1, val.arg2, val.expected)
		}
	}
}

func TestMultiplySuccess(t *testing.T) {

	for _, val := range PositiveStructs {

		// Get result from actual function
		result, err := Multiply(val.arg1, val.arg2)

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, val.expected, result)
	}
}

func TestMultiplyFail(t *testing.T) {

	asserts := assert.New(t)
	for _, val := range NegativeStructs {

		// Get result from actual function
		result, err := Multiply(val.arg1, val.arg2)

		asserts.Nil(err)
		asserts.NotNil(result)
		asserts.NotEqual(val.expected, result)
	}
}

func TestMultiplyAllScenario(t *testing.T) {

	asserts := assert.New(t)

	t.Run("Positive", func(t *testing.T) {

		for _, val := range PositiveStructs {

			// Get result from actual function
			result, err := Multiply(val.arg1, val.arg2)

			asserts.Nil(err)
			asserts.NotNil(result)
			asserts.Equal(val.expected, result)
		}
	})

	t.Run("Negative", func(t *testing.T) {

		for _, val := range NegativeStructs {

			// Get result from actual function
			result, err := Multiply(val.arg1, val.arg2)

			asserts.Nil(err)
			asserts.NotNil(result)
			asserts.NotEqual(val.expected, result)
		}
	})
}

func TestMultiplySkip(t *testing.T) {

	asserts := assert.New(t)

	t.Logf("Total test case: %d", len(PositiveStructs))

	for i := 0; i < len(PositiveStructs); i++ {

		// Get result from actual function
		result, err := Multiply(PositiveStructs[i].arg1, PositiveStructs[i].arg2)

		asserts.Nil(err)
		asserts.NotNil(result)
		asserts.Equal(PositiveStructs[i].expected, result)

		if i >= 3 {
			t.Skip("skip the rest test case because not required to test further")
		}
		t.Logf("%d x %d = %d (Pass)", PositiveStructs[i].arg1, PositiveStructs[i].arg2, PositiveStructs[i].expected)
	}

	// if len(os.Getenv("GOPATH")) == 0 {
	// 	t.Skip("skip this test because go path is not set")
	// }
	// t.Log("Go path:", os.Getenv("GOPATH"))
}

func TestMultiplyCleanup(t *testing.T) {

	asserts := assert.New(t)

	t.Cleanup(func() {
		t.Log("Cleanup")
	})

	t.Log("Running test cases...")

	for _, val := range PositiveStructs {

		// Get result from actual function
		result, err := Multiply(val.arg1, val.arg2)

		asserts.Nil(err)
		asserts.NotNil(result)
		asserts.Equal(val.expected, result)

		t.Logf("%d x %d = %d", val.arg1, val.arg2, val.expected)
	}
}

func TestParallel(t *testing.T) {

	t.Run("One", func(t *testing.T) {
		t.Parallel()
		time.Sleep(2 * time.Second)
	})

	t.Run("Two", func(t *testing.T) {
		t.Parallel()
		time.Sleep(2 * time.Second)
	})

	t.Run("Three", func(t *testing.T) {
		t.Parallel()
		time.Sleep(2 * time.Second)
	})
}
