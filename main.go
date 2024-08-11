package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mrquaketotheworld/gallery/controllers"
	"github.com/mrquaketotheworld/gallery/templates"
	"github.com/mrquaketotheworld/gallery/views"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "home.gohtml"))))
	r.Get("/contact", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "contact.gohtml"))))
	r.Get("/faq", controllers.FAQ(views.Must(views.ParseFS(templates.FS, "faq.gohtml"))))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	fmt.Println("Starting the server at 3000")
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		log.Printf("Error start server: %v", err)
	}
}
