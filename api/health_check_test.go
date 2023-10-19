package api_test

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"strings"
	"testing"

	"core_api/api"

	"github.com/stretchr/testify/assert"
)

func TestHealthApi(t *testing.T) {
	serv := api.NewServer("127.0.0.1", 80)
	router := serv.Setup()
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/health", nil)
	router.ServeHTTP(w, req)
	res := SwapResponseToMapStruct(w.Body.Bytes())

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, true, CheckBodyOkErrors(res))
}

func makeRequestTest() map[string]interface{} {
	mockRequest := strings.NewReader(`{"Username": "test", "Password": "test", "is_admin": 1,"email":"test@co.com"}`)
	jsonValue, _ := json.Marshal(mockRequest)
	serv := api.NewServer("127.0.0.1", 80)
	router := serv.Setup()
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/health", bytes.NewBuffer(jsonValue))
	router.ServeHTTP(w, req)
	return SwapResponseToMapStruct(w.Body.Bytes())
}
