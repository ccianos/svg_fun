package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Handler function to serve SVG markup
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Generate SVG markup
		svg := fmt.Sprintf(`<svg width="200" height="50" xmlns="http://www.w3.org/2000/svg">
  <text x="20" y="30" font-family="Arial" font-size="20" fill="black">Hello World</text>
</svg>`)

		// Set content type to SVG
		w.Header().Set("Content-Type", "image/svg+xml")

		// Write SVG markup to response
		fmt.Fprint(w, svg)
	})

	// Start HTTP server on port 8080
	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}

