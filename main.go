package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func executeTemplate(w http.ResponseWriter, filePath string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	tpl, err := template.ParseFiles(filePath)
	if err != nil {
		log.Printf("Parsing error: %v", err)

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
	executeTemplate(w, "templates/home.gohtml")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, "templates/contact.gohtml")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	pathParam := chi.URLParam(r, "slug")
	if pathParam != "" {
		fmt.Fprintf(w, "<p>This is the article path param: %s</p>", pathParam)
	}
	queryParam := r.URL.Query().Get("id")
	if queryParam != "" {
		fmt.Fprintf(w, "<p>This is the query param: %s</p>", queryParam)
	}

	fmt.Fprint(w, "<h1>Welcome to the FAQs page")
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faqs", faqHandler)
	r.Get("/faqs/{slug}", faqHandler)

	// page not found
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
