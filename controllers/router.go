package controllers

import (
	"net/http"
)

func SetupRouter() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/login", Login)
	router.HandleFunc("/logout", Logout)

	return router
}
