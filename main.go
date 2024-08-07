package main

import (
	"fmt"
	"net/http"
	"github.com/go-chi/chi/v5"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h2>Home. Hello, world</h2>")
}

func contactHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h2>Contact page</h2>")
}

func faqHanlder(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h2>FAQ page</h2>")
}

func main() {
	r := chi.NewRouter()
	r.Get("/{name}", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHanlder)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	fmt.Println("Starting the server at 3000")
	http.ListenAndServe(":3000", r)
}
