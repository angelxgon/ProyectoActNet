package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/angelxgon/ServerBifrost/apis"
	"github.com/angelxgon/ServerBifrost/pinpad"
	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func setupRoutes() {
	r := mux.NewRouter()
	r.HandleFunc("/", homePage)
	r.HandleFunc("/ws", pinpad.WSEndpoint)
	r.HandleFunc("/api/v1/pinpad/conect", apis.Conect).Methods("POST").Schemes("http")
	r.HandleFunc("/api/v1/pinpad/echo", apis.Echo).Methods("POST").Schemes("http")
	r.HandleFunc("/api/v1/pinpad/status", apis.StatusPinPad).Methods("POST").Schemes("http")
	r.HandleFunc("/api/v1/pinpad/config", apis.ConfigPinPad).Methods("POST").Schemes("http")
	r.HandleFunc("/api/v1/pinpad/init", apis.InitPinPad).Methods("POST").Schemes("http")
	r.HandleFunc("/api/v1/pinpad/reset", apis.ResetPinPad).Methods("POST").Schemes("http")
	r.HandleFunc("/api/v1/transaction/sale", apis.Sale).Methods("POST").Schemes("http")
	r.HandleFunc("/api/v1/transaction/refund", apis.Sale).Methods("POST").Schemes("http")
	r.HandleFunc("/api/v1/transaction/cancellation", apis.Sale).Methods("POST").Schemes("http")
	http.Handle("/", r)
}

func main() {
	fmt.Println("Servidor Iniciado en el puerto 8080")
	setupRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
