package routes

import (
	"github.com/IkrmMrbsy/BankMerchantAPI/controllers"
	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	// Routes untuk pelanggan
	r.HandleFunc("/login", controllers.Login).Methods("POST")
	r.HandleFunc("/logout", controllers.Logout).Methods("POST")

	// Routes untuk pembayaran
	r.HandleFunc("/payment", controllers.MakePayment).Methods("POST")

	// Routes untuk merchant
	r.HandleFunc("/merchants", controllers.GetMerchants).Methods("GET")

	return r
}
