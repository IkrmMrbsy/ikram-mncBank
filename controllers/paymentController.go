package controllers

import (
	"fmt"
	"net/http"

	"github.com/IkrmMrbsy/BankMerchantAPI/models"
	"github.com/IkrmMrbsy/BankMerchantAPI/utils"
	"github.com/IkrmMrbsy/BankMerchantAPI/views"
)

var history []models.History

func InitPaymentController() {
	utils.ReadJSONFile("data/history.json", &history)
}

func MakePayment(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	amount := r.URL.Query().Get("amount")

	// Validasi pelanggan
	var customerFound bool
	var customer models.Customer
	for _, c := range customers {
		if c.Email == email && c.IsLoggedIn {
			customerFound = true
			customer = c // Menyimpan customer yang ditemukan
			break
		}
	}

	//meriksa hanya pelanggan yang terdaftar dan sudah login yang dapat melakukan pembayaran
	if !customerFound || !customer.IsLoggedIn {
		views.SendJSONResponse(w, http.StatusUnauthorized, "Customer not found or not logged in")
		return
	}

	// Proses pembayaran
	//mencatat dan menyimpan riwayat pembayaran pelanggan ke dalam file history.json
	action := fmt.Sprintf("Paid amount: %s", amount)
	history = append(history, models.History{
		CustomerID: email,
		Action:     action,
		Timestamp:  fmt.Sprintf("%v", utils.GetCurrentTime()),
	})

	utils.WriteJSONFile("data/history.json", history)

	// Menambahkan keterangan pembayaran dengan nama pengguna dan jumlah pembayaran
	paymentMessage := fmt.Sprintf("Payment successful for %s for amount %s", customer.Name, amount)

	views.SendJSONResponse(w, http.StatusOK, paymentMessage)
}
