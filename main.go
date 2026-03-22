package main

import (
	"fmt"
	"net/http"
	"os"

	"jatinporiya/handler"
)

func main() {
	var port string = "8080"

	http.HandleFunc("/", Home)
	http.HandleFunc("/api/v1/cta", handler.CreateCTA)

	fmt.Println("Server is Running on PORT", port)
	fmt.Fprintf(os.Stdout, "http://localhost:%s\n", port)

	http.ListenAndServe(":8080", nil)
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Aroha Health - Server is Running.</h1>")
	fmt.Fprintf(w, "<h2>Call to Action - API</h2>")
	fmt.Fprintf(w, "<a href='/api/v1/cta'><button>CTA</button></a>")
}
