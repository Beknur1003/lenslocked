package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"lens/controllers"
	"lens/templates"
	"lens/views"
	"net/http"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "home.gohtml", "tailwind.gohtml"))))

	// parse contact template
	r.Get("/contact", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "contact.gohtml", "tailwind.gohtml"))))

	// parse faq template
	r.Get("/faq", controllers.FAQ(
		views.Must(views.ParseFS(templates.FS, "faq.gohtml", "tailwind.gohtml"))))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)

	})

	fmt.Println("Starting the server on :3000 port")

	http.ListenAndServe(":3000", r)
}
