package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type response[T any] struct {
	Data  T      `json:"data"`
	Error string `json:"error"`
}

func TestPostUserController(t *testing.T) {
	user := mockUser{Id: len(mockedUsers) + 1, Email: "email"}
	tests := []struct {
		name          string
		payload       interface{}
		statusCode    int
		responseData  *mockUser
		responseError string
	}{
		{
			name:          "Should fail when request body wrong format",
			payload:       mockUser{},
			statusCode:    http.StatusBadRequest,
			responseData:  nil,
			responseError: "Key: 'mockUser.Email' Error:Field validation for 'Email' failed on the 'required' tag",
		},
		{
			name:          "Should success when request body correct",
			payload:       user,
			statusCode:    http.StatusCreated,
			responseData:  &user,
			responseError: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			c.Request = BuildRequest(tt.payload)

			PostUserController(c)

			assert.Equal(t, tt.statusCode, w.Code)

			var res response[mockUser]
			if err := json.Unmarshal(w.Body.Bytes(), &res); err != nil {
				t.Fatalf("%s", err.Error())
			}

			if tt.responseData != nil {
				assert.EqualValues(t, *tt.responseData, res.Data)
			} else {
				assert.Empty(t, res.Data)
			}
			assert.Equal(t, tt.responseError, res.Error)
		})
	}
}

func TestGetUsersController(t *testing.T) {
	tests := []struct {
		name       string
		statusCode int
		resData    []mockUser
		error      string
	}{
		{
			name:       "GetUsersController OK",
			statusCode: http.StatusOK,
			resData:    mockedUsers,
			error:      "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			GetUsersController(c)

			assert.Equal(t, tt.statusCode, w.Code)

			var res response[[]mockUser]
			if err := json.Unmarshal(w.Body.Bytes(), &res); err != nil {
				t.Fatalf("%s", err.Error())
			}

			assert.Equal(t, res.Error, tt.error)
			assert.EqualValues(t, res.Data, tt.resData)
		})
	}
}
