package handler

import (
	"bytes"
	"errors"
	fhub "github.com/antonchaban/file-hub-go"
	"github.com/antonchaban/file-hub-go/pkg/service"
	"github.com/antonchaban/file-hub-go/pkg/service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"net/http/httptest"
	"testing"
)

func TestHandler_signUp(t *testing.T) {
	type mockBehavior func(s *mock_service.MockAuthorization, user fhub.User)

	testTable := []struct {
		name                 string
		inputBody            string
		inputUser            fhub.User
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "OK",
			inputBody: `{"username": "test", "password": "test"}`,
			inputUser: fhub.User{
				Username: "test",
				Password: "test",
			},
			mockBehavior: func(s *mock_service.MockAuthorization, user fhub.User) {
				s.EXPECT().CreateUser(user).Return(1, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"id":1}`,
		},
		{
			name:                 "Wrong Input",
			inputBody:            `{"username": "username"}`,
			inputUser:            fhub.User{},
			mockBehavior:         func(r *mock_service.MockAuthorization, user fhub.User) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"invalid input body"}`,
		},
		{
			name:      "Service Error",
			inputBody: `{"username": "username", "password": "qwerty"}`,
			inputUser: fhub.User{
				Username: "username",
				Password: "qwerty",
			},
			mockBehavior: func(r *mock_service.MockAuthorization, user fhub.User) {
				r.EXPECT().CreateUser(user).Return(0, errors.New("something went wrong"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"something went wrong"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			auth := mock_service.NewMockAuthorization(c)
			testCase.mockBehavior(auth, testCase.inputUser)

			services := &service.Service{Authorization: auth}
			handler := Handler{services: services}

			r := gin.New()
			r.POST("/sign-up", handler.signUp)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/sign-up",
				bytes.NewBufferString(testCase.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedResponseBody, w.Body.String())
		})
	}
}
