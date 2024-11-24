package tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/IkrmMrbsy/BankMerchantAPI/controllers"
	"github.com/stretchr/testify/assert"
)

func TestLoginSuccess(t *testing.T) {
	// Inisialisasi router
	router := controllers.SetupRouter()

	// Payload untuk permintaan login
	payload := []byte(`{"username":"ikram", "password":"IK2003"}`)

	// Buat permintaan POST ke endpoint /login
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	// Kirim permintaan ke router
	router.ServeHTTP(w, req)

	// Verifikasi respons
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Login successful")
}

func TestLoginFailed(t *testing.T) {
	// Inisialisasi router
	router := controllers.SetupRouter()

	// Payload untuk login gagal
	payload := []byte(`{"username":"wrong_user", "password":"wrong_pass"}`)

	// Buat permintaan POST ke endpoint /login
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	// Kirim permintaan ke router
	router.ServeHTTP(w, req)

	// Verifikasi respons
	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Contains(t, w.Body.String(), "Invalid credentials")
}
