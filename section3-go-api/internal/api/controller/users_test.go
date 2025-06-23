package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetUsersController(t *testing.T) {
	type response struct {
		Users []mockUser
	}

	tests := []struct {
		name       string
		statusCode int
		users      []mockUser
	}{
		{
			name:       "GetUsersController OK",
			statusCode: http.StatusOK,
			users:      mockedUsers,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			GetUsersController(c)

			assert.Equal(t, tt.statusCode, w.Code)

			var res response
			err := json.Unmarshal(w.Body.Bytes(), &res)

			assert.NoError(t, err)

			assert.EqualValues(t, tt.users, res.Users)
		})
	}
}
