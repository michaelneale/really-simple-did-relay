package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var messageStore = make(map[string][]string)

func handleMessage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var msg struct {
			Nonce     string `json:"nonce"`
			Recipient string `json:"recipientDid"`
			Sender    string `json:"senderDid"`
			Payload   string `json:"payload"`
		}
		err := json.NewDecoder(r.Body).Decode(&msg)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		if !checkNonce(msg.Nonce, msg.Sender) {
			http.Error(w, "Bad request", http.StatusBadRequest)
		}

		if _, exists := messageStore[msg.Recipient]; !exists {
			messageStore[msg.Recipient] = make([]string, 0)
		}
		messageStore[msg.Recipient] = append(messageStore[msg.Recipient], msg.Payload)
		fmt.Fprintf(w, "Message stored for recipient %s", msg.Recipient)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	recipient := r.URL.Query().Get("recipientDid")
	if recipient == "" {
		http.Error(w, "Missing recipientDid parameter", http.StatusBadRequest)
		return
	}

	nonce := r.URL.Query().Get("nonce")
	if nonce == "" {
		http.Error(w, "Missing nonce parameter", http.StatusBadRequest)
		return
	}

	if !checkNonce(nonce, recipient) {
		http.Error(w, "Bad request", http.StatusBadRequest)
	}

	payloads, exists := messageStore[recipient]
	if !exists {
		http.Error(w, "No messages found for recipient", http.StatusNotFound)
		return
	}

	response, err := json.Marshal(payloads)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}

	delete(messageStore, recipient)
	w.Write(response)
}

func main() {
	http.HandleFunc("/", handleMessage)
	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", nil)
}

func checkNonce(nonce string, did string) bool {
	// TODO: Implement your nonce validation logic here
	// For the sake of this example, we will simply return true, assuming the nonce is valid.
	return true
}
