package handler

import (
	"bytes"
	fhub "github.com/antonchaban/file-hub-go"
	"github.com/antonchaban/file-hub-go/pkg/service"
	mock_service "github.com/antonchaban/file-hub-go/pkg/service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestFolder_CreateFolder(t *testing.T) {
	type mockBehavior func(s *mock_service.MockFolder, folder fhub.Folder, user int)

	testTable := []struct {
		name                 string
		inputBody            string
		inputFolder          fhub.Folder
		inputUser            int
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "OK",
			inputBody: `{"folder_name": "test"}`,
			inputFolder: fhub.Folder{
				FolderName: "test",
			},
			inputUser: 1,
			mockBehavior: func(s *mock_service.MockFolder, folder fhub.Folder, user int) {
				s.EXPECT().CreateFolder(folder, user).Return(1, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"id":1}`,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			folder := mock_service.NewMockFolder(c)
			testCase.mockBehavior(folder, testCase.inputFolder, testCase.inputUser)

			services := &service.Service{Folder: folder}
			handler := Handler{services: services}

			r := gin.New()
			r.POST("/api/folders", handler.createFolder)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/api/folders",
				bytes.NewBufferString(testCase.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedResponseBody, w.Body.String())
		})
	}
}
