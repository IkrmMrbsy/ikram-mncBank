package controllers

import (
	"net/http"

	"github.com/IkrmMrbsy/BankMerchantAPI/models"
	"github.com/IkrmMrbsy/BankMerchantAPI/utils"
	"github.com/IkrmMrbsy/BankMerchantAPI/views"
)

var merchants []models.Merchant

// InitMerchantController - Membaca data merchant dari file JSON
func InitMerchantController() {
	utils.ReadJSONFile("data/merchants.json", &merchants)
}

// GetMerchants - Mengembalikan daftar semua merchant
func GetMerchants(w http.ResponseWriter, r *http.Request) {
	views.SendJSONResponse(w, http.StatusOK, merchants)
}
