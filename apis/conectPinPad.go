package apis

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/angelxgon/ServerBifrost/models"
	"github.com/gorilla/websocket"
)

func Conect(w http.ResponseWriter, r *http.Request) {

	var input models.SaleRequest

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Decode error! please check your JSON formating.")
		return
	}

	log.Printf("Received: %s\n", input.Amount)
	fmt.Println("Received: %s\n", input.Amount)

	socketUrl := "ws://localhost:8080/ws?username=servidor"
	conn, _, err := websocket.DefaultDialer.Dial(socketUrl, nil)
	if err != nil {
		log.Fatal("Error connecting to Websocket Server:", err)
		return
	}

	fmt.Println("Se conecta al websocket %s", conn.LocalAddr())

	// Response JSON

	data := map[string]interface{}{
		"intValue":    1234,
		"boolValue":   true,
		"stringValue": "hello!",
		"objectValue": map[string]interface{}{
			"arrayValue": []int{1, 2, 3, 4},
		},
	}

	jsonData, err2 := json.Marshal(data)
	if err2 != nil {
		fmt.Printf("could not marshal json: %s\n", err2)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintf(w, string(jsonData))

	envio := conn.WriteMessage(websocket.TextMessage, []byte("Hello from GolangDocs!"))
	if envio != nil {
		log.Println("Error during writing to websocket:", envio)
		fmt.Println("Error during writing to websocket:", envio)
		return
	}
	fmt.Println("Termina Proceso")
}
