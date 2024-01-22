// main.go
package main

import (
	"fmt"
	"net/http"

	"github.com/Anshu-rai89/key-value-store/keyvaluestore"
)

// ... (previously provided HandleSet and HandleGet functions)

func main() {
	// Create a new instance of KeyValueStore
	kv := keyvaluestore.NewKeyValueStore()

	// Set up HTTP handlers
	http.HandleFunc("/set", HandleSet(kv))
	http.HandleFunc("/get", HandleGet(kv))

	// Start the HTTP server
	port := 8080
	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("Starting key-value store on http://localhost%s\n", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}
