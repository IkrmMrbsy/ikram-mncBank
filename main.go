package main

import (
	"log"
	"net/http"

	"github.com/IkrmMrbsy/BankMerchantAPI/controllers"
)

func main() {
	controllers.InitAuthController()
	controllers.InitPaymentController()

	http.HandleFunc("/login", controllers.Login)
	http.HandleFunc("/logout", controllers.Logout)
	http.HandleFunc("/payment", controllers.MakePayment)

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
