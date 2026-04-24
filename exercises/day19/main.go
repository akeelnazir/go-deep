package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price,omitempty"`
}

func marshalJSON(data interface{}) (string, error) {
	bytes, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func unmarshalJSON(jsonStr string, v interface{}) error {
	return json.Unmarshal([]byte(jsonStr), v)
}

func marshalJSONIndent(data interface{}) (string, error) {
	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func main() {
	fmt.Println("=== Day 19: Serialization and Encoding ===")

	fmt.Println("\n--- JSON Marshaling ---")
	person := Person{Name: "Alice", Age: 30, Email: "alice@example.com"}
	jsonStr, _ := marshalJSON(person)
	fmt.Printf("Marshaled: %s\n", jsonStr)

	fmt.Println("\n--- JSON Marshaling with Indent ---")
	jsonIndent, _ := marshalJSONIndent(person)
	fmt.Printf("Marshaled (indented):\n%s\n", jsonIndent)

	fmt.Println("\n--- JSON Unmarshaling ---")
	jsonData := `{"name":"Bob","age":25,"email":"bob@example.com"}`
	var p Person
	unmarshalJSON(jsonData, &p)
	fmt.Printf("Unmarshaled: %+v\n", p)

	fmt.Println("\n--- JSON with Omitempty ---")
	product1 := Product{ID: 1, Name: "Laptop", Price: 999.99}
	product2 := Product{ID: 2, Name: "Mouse"}
	
	json1, _ := marshalJSON(product1)
	json2, _ := marshalJSON(product2)
	
	fmt.Printf("Product 1: %s\n", json1)
	fmt.Printf("Product 2: %s\n", json2)

	fmt.Println("\n--- JSON Array ---")
	people := []Person{
		{Name: "Alice", Age: 30, Email: "alice@example.com"},
		{Name: "Bob", Age: 25, Email: "bob@example.com"},
	}
	jsonArray, _ := marshalJSONIndent(people)
	fmt.Printf("Array:\n%s\n", jsonArray)

	fmt.Println("\n=== Day 19 Complete ===")
	fmt.Println("Next: Learn about CLI applications on Day 20.")
}
