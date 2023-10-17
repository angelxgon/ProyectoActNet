package apis

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"

	"github.com/angelxgon/ServerBifrost/models"
)

func Sale(w http.ResponseWriter, r *http.Request) {

	var input models.SaleRequest

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Decode error! please check your JSON formating.")
		return
	}

	// Response JSON

	var req models.Commands
	req.TypeTrans = "Sale"
	req.Sale.Amount = "50.50"
	req.Sale.Mounths = "6"
	req.Sale.TransactionId = "123456"

	jsonData, err2 := json.Marshal(req)
	if err2 != nil {
		fmt.Printf("could not marshal json: %s\n", err2)
		return
	}

	c, err := net.Dial("tcp", "10.100.163.37:8095") //pinPad
	if err != nil {
		fmt.Println(err)
		return
	}

	c.Write([]byte(jsonData))

	reply := make([]byte, 1024)

	_, err = c.Read(reply)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("reply:", string(reply))

	c.Close()

	data := req

	jsonDataResponse, err2 := json.Marshal(data)
	if err2 != nil {
		fmt.Printf("could not marshal json: %s\n", err2)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(jsonDataResponse))

}

func Refund(w http.ResponseWriter, r *http.Request) {

	var input models.RefundRequest

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Decode error! please check your JSON formating.")
		return
	}

	var data models.RefundResponse

	jsonDataResponse, err2 := json.Marshal(data)
	if err2 != nil {
		fmt.Printf("could not marshal json: %s\n", err2)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(jsonDataResponse))

}

func Cancellation(w http.ResponseWriter, r *http.Request) {

	var input models.CancellationRequest

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Decode error! please check your JSON formating.")
		return
	}

	var data models.CancellationResponse

	jsonDataResponse, err2 := json.Marshal(data)
	if err2 != nil {
		fmt.Printf("could not marshal json: %s\n", err2)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(jsonDataResponse))

}

func Echo(w http.ResponseWriter, r *http.Request) {

	var input models.EchoPinPadRequest

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Decode error! please check your JSON formating.")
		return
	}

	var data models.EchoPinPadResponse

	jsonDataResponse, err2 := json.Marshal(data)
	if err2 != nil {
		fmt.Printf("could not marshal json: %s\n", err2)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(jsonDataResponse))

}

func StatusPinPad(w http.ResponseWriter, r *http.Request) {

	var input models.StatusPinPadRequest

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Decode error! please check your JSON formating.")
		return
	}

	var data models.StatusPinPadResponse

	jsonDataResponse, err2 := json.Marshal(data)
	if err2 != nil {
		fmt.Printf("could not marshal json: %s\n", err2)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(jsonDataResponse))

}

func ConfigPinPad(w http.ResponseWriter, r *http.Request) {

	var input models.ConfigPinPadRequest

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Decode error! please check your JSON formating.")
		return
	}

	var data models.ConfigPinPadResponse

	jsonDataResponse, err2 := json.Marshal(data)
	if err2 != nil {
		fmt.Printf("could not marshal json: %s\n", err2)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(jsonDataResponse))

}

func InitPinPad(w http.ResponseWriter, r *http.Request) {

	var input models.InitPinPadRequest

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Decode error! please check your JSON formating.")
		return
	}

	var data models.InitPinPadResponse

	jsonDataResponse, err2 := json.Marshal(data)
	if err2 != nil {
		fmt.Printf("could not marshal json: %s\n", err2)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(jsonDataResponse))

}

func ResetPinPad(w http.ResponseWriter, r *http.Request) {

	var input models.ResetPinPadRequest

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Decode error! please check your JSON formating.")
		return
	}

	var data models.ResetPinPadResponse

	jsonDataResponse, err2 := json.Marshal(data)
	if err2 != nil {
		fmt.Printf("could not marshal json: %s\n", err2)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(jsonDataResponse))

}
