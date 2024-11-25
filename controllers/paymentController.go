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
	merchantID := r.URL.Query().Get("merchant_id")

	// Validasi pelanggan
	var customerFound bool
	var customer models.Customer
	for _, c := range customers {
		if c.Email == email && c.IsLoggedIn {
			customerFound = true
			customer = c
			break
		}
	}

	if !customerFound {
		views.SendJSONResponse(w, http.StatusUnauthorized, "Customer not found or not logged in")
		return
	}

	// Validasi merchant
	var merchantFound bool
	var merchant models.Merchant
	for _, m := range merchants {
		if m.ID == merchantID {
			merchantFound = true
			merchant = m
			break
		}
	}

	if !merchantFound {
		views.SendJSONResponse(w, http.StatusBadRequest, "Merchant not found")
		return
	}

	// Proses pembayaran
	action := fmt.Sprintf("Paid amount: %s to merchant: %s", amount, merchant.Name)
	history = append(history, models.History{
		CustomerID: email,
		Action:     action,
		Timestamp:  fmt.Sprintf("%v", utils.GetCurrentTime()),
	})

	utils.WriteJSONFile("data/history.json", history)

	// Respon pembayaran
	paymentMessage := fmt.Sprintf("Payment successful for %s to %s for amount %s", customer.Name, merchant.Name, amount)
	views.SendJSONResponse(w, http.StatusOK, paymentMessage)
}
