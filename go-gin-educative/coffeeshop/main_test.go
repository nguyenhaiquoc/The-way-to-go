package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

var r = gin.Default()

func init() {
	gin.SetMode(gin.TestMode)
	r.GET("/ping", ping)
}

func TestPing(t *testing.T) {
	type testCase struct {
		method               string
		path                 string
		expectedResponseCode int
		expectedBody         string
	}

	cases := []testCase{
		{"GET", "/ping", http.StatusOK, "{\"message\":\"Welcome to the Coffeeshop!\"}"},
	}

	for _, tc := range cases {
		req := getRequest(tc.method, tc.path)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		body, _ := ioutil.ReadAll(w.Body)
		if tc.expectedResponseCode != w.Code {
			t.Errorf("Expected '%v', but got '%v'", tc.expectedResponseCode, w.Code)
		}
		if tc.expectedBody != string(body) {
			t.Errorf("Expected '%v', but got '%v'", tc.expectedBody, string(body))
		}
	}
}

func getRequest(method, path string) *http.Request {
	req, _ := http.NewRequest(method, path, nil)
	return req
}
