package routes

import (
	"github.com/IkrmMrbsy/BankMerchantAPI/controllers"
	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/login", controllers.Login).Methods("POST")
	r.HandleFunc("/logout", controllers.Logout).Methods("POST")
	r.HandleFunc("/payment", controllers.MakePayment).Methods("POST")

	return r
}
