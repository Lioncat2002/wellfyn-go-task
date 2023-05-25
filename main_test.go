package main_test

import (
	"bytes"
	"encoding/json"
	"main/models/database"
	"main/providers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert/v2"
)

type Data struct {
	Token string `json:"token"`
}

var data Data

func TestSignUpRoute(t *testing.T) {

	router := providers.InitRouter()
	w := httptest.NewRecorder()
	database.DB.Exec("DELETE FROM users WHERE email = ?", "helo@example.com")
	body := "{\"email\":\"helo@example.com\",\"password\":\"123456\"}"
	req, _ := http.NewRequest("POST", "/api/signup", bytes.NewBufferString(body))
	router.ServeHTTP(w, req)
	t.Log(w.Body.String())
	err := json.Unmarshal(w.Body.Bytes(), &data)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	assert.Equal(t, 200, w.Code)

}

func TestLogin(t *testing.T) {
	router := providers.InitRouter()
	w := httptest.NewRecorder()
	body := "{\"email\":\"helo@example.com\",\"password\":\"123456\"}"
	req, _ := http.NewRequest("POST", "/api/login", bytes.NewBufferString(body))
	router.ServeHTTP(w, req)
	err := json.Unmarshal(w.Body.Bytes(), &data)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	assert.Equal(t, 200, w.Code)
	assert.NotEqual(t, "", data.Token)
}

func TestCurrentUser(t *testing.T) {
	router := providers.InitRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/user/current?token="+data.Token, nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestUpdateUser(t *testing.T) {
	router := providers.InitRouter()
	w := httptest.NewRecorder()
	body := `{"name":"Tigercat","imgurl":"https://picsum.photos/200","phno":"+9123456789"}`
	req, _ := http.NewRequest("POST", "/api/user/current?token="+data.Token, bytes.NewBufferString(body))
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}
