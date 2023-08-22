package main

import (
	"net/http"
	netUrl "net/url"
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

func Test_assert_nil(t *testing.T) {
	tester.SetTester(t)
	tester.AssertNil(nil)
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

func Test_assert_response_writer_message(t *testing.T) {
	tester.SetTester(t)

	helloMessage := []byte("Hello")

	callback := func(writer http.ResponseWriter, request *http.Request) {
		writer.Write(helloMessage)
	}

	mux := http.NewServeMux()

	url := "/some/url/"

	mux.HandleFunc(url, callback)

	urlObj, err := netUrl.Parse(url)
	tester.AssertNil(err)

	request := http.Request{
		Method: http.MethodGet,
		URL:    urlObj,
	}

	handler, _ := mux.Handler(&request)

	writer := tester.GetTestResponseWriter()

	handler.ServeHTTP(writer, &request)

	tester.AssertLen(writer.GetMessages(), 1)

	responseMessage := writer.GetMessages()[0]

	tester.AssertSame(responseMessage, helloMessage)
}

type someStruct struct {
	SomeValue int
}
