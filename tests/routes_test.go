package tests

import (
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/lmhuong711/go-go-be/server"

	"github.com/stretchr/testify/assert"
)

var app = server.New()

func TestRoutes(t *testing.T) {
	tests := []struct {
		testName     string
		route        string
		method       string
		body         io.Reader
		expectedCode int
		expectedBody string
	}{
		{
			testName:     "an invalid path",
			route:        "/api/v1/invalid-path",
			method:       "GET",
			body:         nil,
			expectedCode: 404,
			expectedBody: "{\"error\":\"Cannot GET /api/v1/invalid-path\"}",
		},
		{
			testName:     "another invalid path",
			route:        "/api/v1/another-invalid-path",
			method:       "POST",
			body:         nil,
			expectedCode: 404,
			expectedBody: "{\"error\":\"Cannot POST /api/v1/another-invalid-path\"}",
		},
		{
			testName:     "Get all students",
			route:        "/api/v1/students",
			method:       "GET",
			body:         nil,
			expectedBody: "",
			expectedCode: 200,
		},
		{
			testName:     "Get detail student",
			route:        "/api/v1/student/1",
			method:       "GET",
			body:         nil,
			expectedBody: "",
			expectedCode: 404,
		},
		{
			testName: "Create a student",
			route:    "/api/v1/students",
			method:   "POST",
			body: strings.NewReader(
				`{"first_name":"Phineas","last_name":"Flynn","phone":"0987654320","email":"meowmeowgogo@gmail.com"}`,
			),
			expectedBody: "",
			expectedCode: 200,
		},
		{
			testName: "Update a student",
			route:    "/api/v1/students/1",
			method:   "PUT",
			body: strings.NewReader(
				`{"id":1,"first_name":"Phineass","last_name":"Flynn","phone":"0987754220","email":"phineaaamewo@gmail.com"}`,
			),
			expectedBody: "",
			expectedCode: 200,
		},
		{
			testName:     "Delete a student",
			route:        "/api/v1/students/1",
			method:       "DELETE",
			body:         nil,
			expectedBody: `{"data":null,"message":"Student deleted successfully","success":true}`,
			expectedCode: 200,
		},
	}

	InitTestDB()

	for _, test := range tests {
		req, _ := http.NewRequest(
			test.method,
			test.route,
			test.body,
		)
		req.Close = true
		req.Header.Set("Content-Type", "application/json")

		res, err := app.Test(req, -1)
		assert.Nilf(t, err, test.testName)
		assert.Equalf(t, test.expectedCode, res.StatusCode, test.testName)

		if test.expectedBody != "" {
			body, err := ioutil.ReadAll(res.Body)
			assert.Nilf(t, err, test.testName)
			assert.Equalf(t, test.expectedBody, string(body), test.testName)
		}
		defer res.Body.Close()

	}
}
