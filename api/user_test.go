package api_test

import (
	"core_api/api"
	"core_api/utils"
	"net/http/httptest"
	"strings"
	"testing"

	"encoding/json"

	"github.com/stretchr/testify/assert"
)

func TestSigninigUserSuccess(t *testing.T) {
	serv := api.NewServer("127.0.0.1", 80)
	router := serv.Setup()
	w := httptest.NewRecorder()
	conf := utils.NewConfig()
	conf.LoadConfig("..")

	reqUser := api.UserSigninRequest{
		Username: conf.ApiAdminUsername,
		Password: conf.ApiAdminPassword,
	}

	reqBodyJson, _ := json.Marshal(reqUser)

	reqBody := strings.NewReader(string(reqBodyJson))
	req := httptest.NewRequest("POST", "/users/signin/", reqBody)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	res := SwapResponseToMapStruct(w.Body.Bytes())

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, true, CheckBodyOkErrors(res))

	jsonData, err := json.Marshal(res["data"])
	if err != nil {
		t.Error(err)
	}

	data := SwapResponseToMapStruct([]byte(jsonData))

	if token, ok := data["token"]; !ok {
		t.Error("token is not exist.")
		_ = token
	}

	if apiKey, ok := data["api_key"]; !ok {
		t.Error("api_key is not exist.")
		_ = apiKey
	}

}
