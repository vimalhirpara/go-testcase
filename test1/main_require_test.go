package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAssertPackage(t *testing.T) {

	asserts := assert.New(t)

	mergedStruct := append(PositiveStructs, NegativeStructs...)

	t.Logf("Total test cases: %d (5 Positive, 5 Negative)", len(mergedStruct))
	t.Log(mergedStruct)

	for _, val := range mergedStruct {

		t.Logf("%d x %d = %d", val.arg1, val.arg2, val.expected)

		// Get result from actual function
		result, err := Multiply(val.arg1, val.arg2)

		asserts.Nil(err)
		asserts.NotNil(result)
		asserts.Equal(val.expected, result)
	}
}

func TestRequirePackage(t *testing.T) {

	requires := require.New(t)
	mergedStruct := append(PositiveStructs, NegativeStructs...)

	t.Logf("Total test cases: %d (5 Positive, 5 Negative)", len(mergedStruct))
	t.Log(mergedStruct)

	for _, val := range mergedStruct {

		t.Logf("%d x %d = %d", val.arg1, val.arg2, val.expected)

		// Get result from actual function
		result, err := Multiply(val.arg1, val.arg2)

		requires.Nil(err)
		requires.NotNil(result)
		requires.Equal(val.expected, result)
	}
}
