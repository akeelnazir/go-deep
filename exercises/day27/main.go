package main

import (
	"fmt"
)

type Resource struct {
	Name    string
	Content string
}

var resources = make(map[string]Resource)

func registerResource(name, content string) {
	resources[name] = Resource{Name: name, Content: content}
}

func getResource(name string) *Resource {
	if r, ok := resources[name]; ok {
		return &r
	}
	return nil
}

func listResources() []string {
	names := make([]string, 0, len(resources))
	for name := range resources {
		names = append(names, name)
	}
	return names
}

func init() {
	registerResource("template.html", "<html><body>Hello</body></html>")
	registerResource("config.json", `{"version":"1.0"}`)
	registerResource("style.css", "body { color: blue; }")
}

func main() {
	fmt.Println("=== Day 27: Embedding and Code Generation ===")

	fmt.Println("\n--- Registering Resources ---")
	registerResource("script.js", "console.log('Hello');")
	fmt.Printf("Registered resources: %d\n", len(resources))

	fmt.Println("\n--- Listing Resources ---")
	for _, name := range listResources() {
		fmt.Printf("  - %s\n", name)
	}

	fmt.Println("\n--- Retrieving Resources ---")
	resource := getResource("template.html")
	if resource != nil {
		fmt.Printf("Resource: %s\n", resource.Name)
		fmt.Printf("Content: %s\n", resource.Content)
	}

	fmt.Println("\n--- Resource Content Sizes ---")
	for _, name := range listResources() {
		r := getResource(name)
		if r != nil {
			fmt.Printf("%s: %d bytes\n", name, len(r.Content))
		}
	}

	fmt.Println("\n=== Day 27 Complete ===")
	fmt.Println("Next: Learn about advanced Go on Day 28.")
}
