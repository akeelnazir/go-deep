package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

// Section 1: Basic Router
type Router struct {
	routes map[string]http.HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		routes: make(map[string]http.HandlerFunc),
	}
}

func (r *Router) Register(path string, handler http.HandlerFunc) {
	r.routes[path] = handler
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	handler, ok := r.routes[req.URL.Path]
	if !ok {
		http.NotFound(w, req)
		return
	}
	handler(w, req)
}

// Section 2: Middleware type and common implementations
type Middleware func(http.Handler) http.Handler

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("[%s] %s %s", r.Method, r.URL.Path, r.RemoteAddr)
		next.ServeHTTP(w, r)
		duration := time.Since(start)
		log.Printf("Request completed in %v", duration)
	})
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}

		if !isValidToken(token) {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func isValidToken(token string) bool {
	return token == "valid-token"
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func requestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := fmt.Sprintf("%d", time.Now().UnixNano())
		ctx := context.WithValue(r.Context(), "request-id", requestID)
		w.Header().Set("X-Request-ID", requestID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Section 3: Middleware chaining
func chain(handler http.Handler, middlewares ...Middleware) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}
	return handler
}

// Section 4: Context propagation
func contextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "user", "alice")
		ctx = context.WithValue(ctx, "role", "admin")

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func contextAwareHandler(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(string)
	role := r.Context().Value("role").(string)

	fmt.Fprintf(w, "User: %s, Role: %s", user, role)
}

// Section 5: Recovery middleware
func recoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic: %v", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	})
}

func main() {
	fmt.Println("=== Day 12: Routing and Middleware ===")

	// Section 1: Basic Router
	fmt.Println("\n--- Basic Router ---")
	router := NewRouter()
	router.Register("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Home Page")
	})
	router.Register("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "About Page")
	})
	router.Register("/contact", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Contact Page")
	})
	fmt.Println("Router created with 3 routes: /, /about, /contact")

	// Section 2: Middleware demonstration
	fmt.Println("\n--- Middleware Chaining ---")
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from chained middleware")
	})

	_ = chain(mux,
		loggingMiddleware,
		corsMiddleware,
	)
	fmt.Println("Created middleware chain: logging -> CORS -> handler")

	// Section 3: Context propagation
	fmt.Println("\n--- Context Propagation ---")
	contextMux := http.NewServeMux()
	contextMux.HandleFunc("/user", contextAwareHandler)

	_ = chain(contextMux,
		contextMiddleware,
		requestIDMiddleware,
	)
	fmt.Println("Created context chain: request-id -> context -> handler")

	// Section 4: Recovery middleware
	fmt.Println("\n--- Recovery Middleware ---")
	recoveryMux := http.NewServeMux()
	recoveryMux.HandleFunc("/safe", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Safe operation")
	})

	_ = chain(recoveryMux,
		recoveryMiddleware,
		loggingMiddleware,
	)
	fmt.Println("Created recovery chain: recovery -> logging -> handler")

	fmt.Println("\n=== Day 12 Complete ===")
	fmt.Println("Next: Learn about REST API design on Day 13.")
}
