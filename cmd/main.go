package main

import (
	"fmt"
	"log"
	"net/http"

	api "github.com/john6fa/receipt-processor-challenge/internal/api"

	"github.com/gorilla/mux"
)

func main() {
	handler := api.NewHandler()

	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintln(w, "Receipt Processor API")
		fmt.Fprintln(w, "/receipts/process (POST): Process a receipt and store in memory.")
		fmt.Fprintln(w, "/receipts/{id}/points (GET): Get the number of points awarded for a receipt.")
	})
	router.HandleFunc("/receipts/process", handler.ProcessReceipts).Methods("POST")
	router.HandleFunc("/receipts/{id}/points", handler.GetPoints).Methods("GET")

	log.Println("Starting server on port 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
