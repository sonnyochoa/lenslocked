package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
)

func executeTemplate(w http.ResponseWriter, filepath string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tpl, err := template.ParseFiles(filepath)
	if err != nil {
		log.Printf("parsing tempalte: %v", err)
		http.Error(w, "There was an error parsing the template.", http.StatusInternalServerError)
		return
	}
	err = tpl.Execute(w, nil)
	if err != nil {
		log.Printf("executing template: %v", err)
		http.Error(w, "There was an error executing the template.", http.StatusInternalServerError)
		return
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "home.gohtml")
	executeTemplate(w, tplPath)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "contact.gohtml")
	executeTemplate(w, tplPath)
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	var faqHTML string = "<body><main><h1 class=\"faq-heading\">FAQ'S</h1><section class=\"faq-container\"><div class=\"faq-one\"><!-- faq question --><h1 class=\"faq-page\">What is an FAQ Page?</h1><!-- faq answer --><div class=\"faq-body\"> <p>Lorem ipsum dolor, sit amet consectetur adipisicing elit. Velit saepe sequi, illum facere necessitatibus cum aliquam id illo omnis maxime, totam soluta voluptate amet ut sit ipsum aperiam. Perspiciatis, porro!</p></div></div><hr class=\"hr-line\"><div class=\"faq-two\"><!-- faq question --><h1 class=\"faq-page\">Why do you need an FAQ page?</h1><!-- faq answer --><div class=\"faq-body\"><p>Lorem ipsum dolor, sit amet consectetur adipisicing elit. Velit saepe sequi, illum facere necessitatibus cum aliquam id illo omnis maxime, totam soluta voluptate amet ut sit ipsum aperiam. Perspiciatis, porro!</p></div></div><hr class=\"hr-line\"><div class=\"faq-three\"><!-- faq question --><h1 class=\"faq-page\">Does it improves the user experience of a website?</h1><!-- faq answer --><div class=\"faq-body\"><p>Lorem ipsum dolor, sit amet consectetur adipisicing elit. Velit saepe sequi, illum facere necessitatibus cum aliquam id illo omnis maxime, totam soluta voluptate amet ut sit ipsum aperiam. Perspiciatis, porro!</p></div></div></section></main></body>"
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, faqHTML)
}

func main() {
	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found.", http.StatusNotFound)
	})
	fmt.Println("Starting the serve on :3000...")
	http.ListenAndServe(":3000", r)
}
