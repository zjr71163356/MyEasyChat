package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"myeasychat/user/api/internal/config"
	"myeasychat/user/api/internal/svc"
	"myeasychat/user/api/internal/types"
)

func TestgetUserInfoHandler(t *testing.T) {
	// new service context
	c := config.Config{}
	svcCtx := svc.NewServiceContext(c)
	// init mock service context here

	tests := []struct {
		name       string
		reqBody    interface{}
		wantStatus int
		wantResp   string
		setupMocks func()
	}{
		{
			name:       "invalid request body",
			reqBody:    "invalid",
			wantStatus: http.StatusBadRequest,
			wantResp:   "unsupported type", // Adjust based on actual error response
			setupMocks: func() {
				// No setup needed for this test case
			},
		},
		{
			name:    "handler error",
			reqBody: types.GetUserInfoReq{
				//TODO: add fields here
			},
			wantStatus: http.StatusBadRequest,
			wantResp:   "error", // Adjust based on actual error response
			setupMocks: func() {
				// Mock login logic to return an error
			},
		},
		{
			name:    "handler successful",
			reqBody: types.GetUserInfoReq{
				//TODO: add fields here
			},
			wantStatus: http.StatusOK,
			wantResp:   `{"code":0,"msg":"success","data":{}}`, // Adjust based on actual success response
			setupMocks: func() {
				// Mock login logic to return success
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMocks()
			var reqBody []byte
			var err error
			reqBody, err = json.Marshal(tt.reqBody)
			require.NoError(t, err)
			req, err := http.NewRequest("POST", "/ut", bytes.NewBuffer(reqBody))
			require.NoError(t, err)
			req.Header.Set("Content-Type", "application/json")

			rr := httptest.NewRecorder()
			handler := getUserInfoHandler(svcCtx)
			handler.ServeHTTP(rr, req)
			t.Log(rr.Body.String())
			assert.Equal(t, tt.wantStatus, rr.Code)
			assert.Contains(t, rr.Body.String(), tt.wantResp)
		})
	}
}
