package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

type testingUserError string

func (e testingUserError) Error() string {
	return e.Message()
}

func (e testingUserError) Message() string {
	return string(e)
}

func errPanic(_ http.ResponseWriter, _ *http.Request) error {
	panic(123)
}

func errUserError(_ http.ResponseWriter, _ *http.Request) error {
	return testingUserError("user error")
}

func errNotFound(_ http.ResponseWriter, _ *http.Request) error {
	return os.ErrNotExist
}

func errNoPermission(_ http.ResponseWriter, _ *http.Request) error {
	return os.ErrPermission
}

func errUnknown(_ http.ResponseWriter, _ *http.Request) error {
	return os.ErrInvalid
}

func noError(writer http.ResponseWriter,
	_ *http.Request) error {
	fmt.Fprintln(writer, "no error")
	return nil
}

var tests = []struct {
	h       appHandler
	code    int
	message string
}{
	{errPanic, 500, "Internal Server Error"},
	{errUserError, 400, "user error"},
	{errNotFound, 404, "Not Found"},
	{errNoPermission, 403, "Forbidden"},
	{errUnknown, 500, "Internal Server Error"},
	{noError, 200, "no error"},
}

// 通过函数调用去测试
func TestErrWrapper(t *testing.T) {
	for _, tt := range tests {
		f := errWrapper(tt.h)
		response := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "http://www.baidu.com", nil)
		f(response, request)
		verifyResponse(response.Result(), tt.code, tt.message, t)
	}
}

// 通过http server去测试
func TestErrWrapperInServer(t *testing.T) {
	for _, tt := range tests {
		f := errWrapper(tt.h)
		server := httptest.NewServer(http.HandlerFunc(f))
		response, _ := http.Get(server.URL)
		verifyResponse(response, tt.code, tt.message, t)
	}
}

func verifyResponse(resp *http.Response, expectedCode int, expectedMsg string, t *testing.T) {
	b, _ := io.ReadAll(resp.Body)
	body := strings.Trim(string(b), "\n")
	if resp.StatusCode != expectedCode || body != expectedMsg {
		t.Errorf("expect (%d, %s); "+
			"got (%d, %s)",
			expectedCode, expectedMsg,
			resp.StatusCode, body)
	}
}
