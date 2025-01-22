package main

import (
	"fmt"
	api "jobproblem1/api"
	"log"
	"net/http"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func main() {
	// Create service instance.
	rss := map[string]ReceiptMetaData{}
	rss["hi"] = ReceiptMetaData{
		receipt: api.Receipt{
			Retailer: "chc",
			Total:    "100",
		},
		points: 10,
		id:     "hi",
	}

	service := &ReceiptService{
		receipts: rss,
	}
	// Create generated server.
	srv, err := api.NewServer(service)
	if err != nil {
		log.Fatal(err)
	}
	println("Server Starting")
	if err := http.ListenAndServe(":8080", srv); err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/", greet)
	http.ListenAndServe(":8080", nil)
}
