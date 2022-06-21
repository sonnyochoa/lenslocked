package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/sonnyochoa/lenslocked/controllers"
	"github.com/sonnyochoa/lenslocked/views"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", controllers.StaticHandler(views.Must(views.Parse(filepath.Join("templates", "home.gohtml")))))
	r.Get("/contact", controllers.StaticHandler(views.Must(views.Parse(filepath.Join("templates", "contact.gohtml")))))
	r.Get("/faq", controllers.StaticHandler(views.Must(views.Parse(filepath.Join("templates", "faq.gohtml")))))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found.", http.StatusNotFound)
	})
	fmt.Println("Starting the serve on :3000...")
	http.ListenAndServe(":3000", r)
}
