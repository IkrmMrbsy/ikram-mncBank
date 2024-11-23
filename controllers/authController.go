package controllers

import (
	"net/http"

	"github.com/IkrmMrbsy/BankMerchantAPI/models"
	"github.com/IkrmMrbsy/BankMerchantAPI/utils"
	"github.com/IkrmMrbsy/BankMerchantAPI/views"
)

var customers []models.Customer

// fungsi untuk membaca data pelanggan dari file customers.json
func InitAuthController() {
	utils.ReadJSONFile("data/customers.json", &customers)
}

// Login
// Mengautentikasi pelanggan berdasarkan email dan password
func Login(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	password := r.URL.Query().Get("password")

	for i, customer := range customers {
		if customer.Email == email && customer.Password == password {
			customers[i].IsLoggedIn = true
			utils.WriteJSONFile("data/customers.json", customers)
			views.SendJSONResponse(w, http.StatusOK, "Login successful")
			return
		}
	}

	views.SendJSONResponse(w, http.StatusUnauthorized, "Invalid credentials")
}

// Logout
// Menghapus status login pelanggan berdasarkan email yang dikirim oleh klien.
func Logout(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")

	// Cari pelanggan berdasarkan email
	for i, customer := range customers {
		if customer.Email == email {
			// Jika pelanggan sudah login
			if customer.IsLoggedIn {
				customers[i].IsLoggedIn = false
				utils.WriteJSONFile("data/customers.json", customers)
				views.SendJSONResponse(w, http.StatusOK, "Logout successful")
			} else {
				// Jika pelanggan tidak login
				views.SendJSONResponse(w, http.StatusBadRequest, "User is not logged in")
			}
			return
		}
	}

	// Jika pelanggan tidak ditemukan
	views.SendJSONResponse(w, http.StatusNotFound, "User not found")
}
