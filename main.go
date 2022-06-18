package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to my great site</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch, email me at <a href=\"mailto:create@sonnylife.dev\">create@sonnylife.dev</a>.")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	var faqHTML string = "<body><main><h1 class=\"faq-heading\">FAQ'S</h1><section class=\"faq-container\"><div class=\"faq-one\"><!-- faq question --><h1 class=\"faq-page\">What is an FAQ Page?</h1><!-- faq answer --><div class=\"faq-body\"> <p>Lorem ipsum dolor, sit amet consectetur adipisicing elit. Velit saepe sequi, illum facere necessitatibus cum aliquam id illo omnis maxime, totam soluta voluptate amet ut sit ipsum aperiam. Perspiciatis, porro!</p></div></div><hr class=\"hr-line\"><div class=\"faq-two\"><!-- faq question --><h1 class=\"faq-page\">Why do you need an FAQ page?</h1><!-- faq answer --><div class=\"faq-body\"><p>Lorem ipsum dolor, sit amet consectetur adipisicing elit. Velit saepe sequi, illum facere necessitatibus cum aliquam id illo omnis maxime, totam soluta voluptate amet ut sit ipsum aperiam. Perspiciatis, porro!</p></div></div><hr class=\"hr-line\"><div class=\"faq-three\"><!-- faq question --><h1 class=\"faq-page\">Does it improves the user experience of a website?</h1><!-- faq answer --><div class=\"faq-body\"><p>Lorem ipsum dolor, sit amet consectetur adipisicing elit. Velit saepe sequi, illum facere necessitatibus cum aliquam id illo omnis maxime, totam soluta voluptate amet ut sit ipsum aperiam. Perspiciatis, porro!</p></div></div></section></main></body>"
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, faqHTML)
}

// func pathHandler(w http.ResponseWriter, r *http.Request) {
// 	switch r.URL.Path {
// 	case "/":
// 		homeHandler(w, r)
// 	case "/contact":
// 		contactHandler(w, r)
// 	default:
// 		// TODO: handle the page not found error
// 		http.Error(w, "Page not found.", http.StatusNotFound)
// 	}
// }

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	case "/faq":
		faqHandler(w, r)
	default:
		// TODO: handle the page not found error
		http.Error(w, "Page not found.", http.StatusNotFound)
	}
}

func main() {
	var router Router
	fmt.Println("Starting the serve on :3000...")
	http.ListenAndServe(":3000", router)
}
