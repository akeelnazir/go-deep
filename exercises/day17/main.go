package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
)

type TestResponse struct {
	Status string `json:"status"`
	Data   string `json:"data"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, World!"))
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(TestResponse{
		Status: "ok",
		Data:   "test data",
	})
}

func errorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("Internal Server Error"))
}

func main() {
	fmt.Println("=== Day 17: Testing Web Applications ===")

	fmt.Println("\n--- Testing HTTP Handlers ---")

	req := httptest.NewRequest("GET", "/hello", nil)
	w := httptest.NewRecorder()
	helloHandler(w, req)

	fmt.Printf("Status: %d\n", w.Code)
	fmt.Printf("Body: %s\n", w.Body.String())

	fmt.Println("\n--- Testing JSON Handler ---")

	req = httptest.NewRequest("GET", "/json", nil)
	w = httptest.NewRecorder()
	jsonHandler(w, req)

	fmt.Printf("Status: %d\n", w.Code)
	fmt.Printf("Content-Type: %s\n", w.Header().Get("Content-Type"))
	fmt.Printf("Body: %s\n", w.Body.String())

	fmt.Println("\n--- Testing Error Handler ---")

	req = httptest.NewRequest("GET", "/error", nil)
	w = httptest.NewRecorder()
	errorHandler(w, req)

	fmt.Printf("Status: %d\n", w.Code)
	fmt.Printf("Body: %s\n", w.Body.String())

	fmt.Println("\n--- Testing with Mock Server ---")

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Mock response"))
	}))
	defer server.Close()

	resp, err := http.Get(server.URL)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Mock server status: %d\n", resp.StatusCode)
		defer resp.Body.Close()
	}

	fmt.Println("\n=== Day 17 Complete ===")
	fmt.Println("Next: Learn about file and IO operations on Day 18.")
}
