package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
)

// Section 1: Simple HTTP Handler
func simpleHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello, World!")
}

// Section 2: HTTP Methods Handler
func methodHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fmt.Fprintf(w, "GET request received")
	case http.MethodPost:
		fmt.Fprintf(w, "POST request received")
	case http.MethodPut:
		fmt.Fprintf(w, "PUT request received")
	case http.MethodDelete:
		fmt.Fprintf(w, "DELETE request received")
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// Section 3: Query Parameters Handler
func queryHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	name := query.Get("name")
	age := query.Get("age")

	if name == "" {
		http.Error(w, "name parameter required", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Name: %s, Age: %s", name, age)
}

// Section 4: JSON Request Handler
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func jsonRequestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Created user: %s (%s)", user.Name, user.Email)
}

// Section 5: JSON Response Handler
func jsonResponseHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"status": "success",
		"data": map[string]string{
			"message": "Hello from API",
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// Section 6: Status Codes Handler
func statusHandler(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("code")

	switch status {
	case "201":
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "Resource created")
	case "400":
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Bad request")
	case "404":
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Not found")
	case "500":
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Server error")
	default:
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "OK")
	}
}

// Section 7: Headers Handler
func headersHandler(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	userAgent := r.Header.Get("User-Agent")

	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("X-Custom-Header", "custom-value")

	fmt.Fprintf(w, "Content-Type: %s\nUser-Agent: %s\n", contentType, userAgent)
}

// Section 8: Routing with ServeMux
func setupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", simpleHandler)
	mux.HandleFunc("/method", methodHandler)
	mux.HandleFunc("/query", queryHandler)
	mux.HandleFunc("/json-request", jsonRequestHandler)
	mux.HandleFunc("/json-response", jsonResponseHandler)
	mux.HandleFunc("/status", statusHandler)
	mux.HandleFunc("/headers", headersHandler)

	return mux
}

// Section 9: HTTP Client - GET Request
func exampleGetRequest() (string, error) {
	mux := setupRoutes()
	server := httptest.NewServer(mux)
	defer server.Close()

	resp, err := http.Get(server.URL + "/json-response")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// Section 10: HTTP Client - POST Request
func examplePostRequest() (string, error) {
	mux := setupRoutes()
	server := httptest.NewServer(mux)
	defer server.Close()

	user := User{Name: "Alice", Email: "alice@example.com"}
	jsonData, _ := json.Marshal(user)

	resp, err := http.Post(
		server.URL+"/json-request",
		"application/json",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// Section 11: Custom HTTP Client
func exampleCustomClient() (string, error) {
	mux := setupRoutes()
	server := httptest.NewServer(mux)
	defer server.Close()

	client := &http.Client{}

	req, _ := http.NewRequest("GET", server.URL+"/headers", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "CustomClient/1.0")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func main() {
	fmt.Println("=== Day 11: Web Fundamentals and HTTP ===")

	// Section 1: Simple Handler
	fmt.Println("\n--- Simple HTTP Handler ---")
	mux := setupRoutes()
	server := httptest.NewServer(mux)
	defer server.Close()

	resp, _ := http.Get(server.URL + "/")
	body, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Printf("Response: %s\n", string(body))

	// Section 2: HTTP Methods
	fmt.Println("\n--- HTTP Methods ---")
	for _, method := range []string{"GET", "POST", "PUT", "DELETE"} {
		req, _ := http.NewRequest(method, server.URL+"/method", nil)
		resp, _ := http.DefaultClient.Do(req)
		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		fmt.Printf("%s: %s\n", method, string(body))
	}

	// Section 3: Query Parameters
	fmt.Println("\n--- Query Parameters ---")
	resp, _ = http.Get(server.URL + "/query?name=Alice&age=30")
	body, _ = io.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Printf("Response: %s\n", string(body))

	// Section 4: JSON Request
	fmt.Println("\n--- JSON Request ---")
	result, _ := examplePostRequest()
	fmt.Printf("Response: %s\n", strings.TrimSpace(result))

	// Section 5: JSON Response
	fmt.Println("\n--- JSON Response ---")
	result, _ = exampleGetRequest()
	fmt.Printf("Response: %s\n", strings.TrimSpace(result))

	// Section 6: Status Codes
	fmt.Println("\n--- Status Codes ---")
	for _, code := range []string{"201", "400", "404"} {
		resp, _ := http.Get(server.URL + "/status?code=" + code)
		fmt.Printf("Status %s: %d\n", code, resp.StatusCode)
		resp.Body.Close()
	}

	// Section 7: Headers
	fmt.Println("\n--- Headers ---")
	result, _ = exampleCustomClient()
	fmt.Printf("Response:\n%s\n", strings.TrimSpace(result))

	// Section 8: Routing
	fmt.Println("\n--- Routing with ServeMux ---")
	routes := []string{"/", "/method", "/query", "/json-response"}
	for _, route := range routes {
		resp, _ := http.Get(server.URL + route)
		fmt.Printf("Route %s: Status %d\n", route, resp.StatusCode)
		resp.Body.Close()
	}

	fmt.Println("\n=== Day 11 Complete ===")
	fmt.Println("Next: Learn about routing and middleware on Day 12.")
}
