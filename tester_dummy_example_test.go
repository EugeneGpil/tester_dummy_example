package main

import (
	"testing"

	"github.com/EugeneGpil/tester"
)

func Test_assert_len(t *testing.T) {
	tester.SetTester(t)

	sliceOfStrings := []string{
		"hello1",
		"hello2",
	}
	tester.AssertLen(sliceOfStrings, 2)

	sliceOfInts := []int{1, 2, 3}
	tester.AssertLen(sliceOfInts, 3)

	sliceOfSomeStructs := []someStruct{
		{1},
		{2},
	}
	tester.AssertLen(sliceOfSomeStructs, 2)
}

func Test_assert_not_nil(t *testing.T) {
	tester.SetTester(t)

	tester.AssertNotNil(1)
	tester.AssertNotNil("nil")
}

func Test_assert_same(t *testing.T) {
	tester.SetTester(t)

	tester.AssertSame("test", "test")
	tester.AssertSame(1, 1)
	tester.AssertSame([]byte("Hello"), []byte("Hello"))
}

type someStruct struct {
	SomeValue int
}
