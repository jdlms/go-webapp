package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/jdlms/go-webapp/controllers"
	"github.com/jdlms/go-webapp/templates"
	"github.com/jdlms/go-webapp/views"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	// these are variadic functions which accept multiple strings
	// the order of the html files does matter
	tpl := views.Must(views.ParseFS(templates.FS, "layout.gohtml", "home.gohtml"))
	r.Get("/", controllers.StaticHandler(tpl))

	tpl = views.Must(views.ParseFS(templates.FS, "layout.gohtml", "contact.gohtml"))
	r.Get("/contact", controllers.StaticHandler(tpl))

	tpl = views.Must(views.ParseFS(templates.FS, "layout.gohtml", "faqs.gohtml"))
	r.Get("/faqs", controllers.FAQ(tpl))

	usersC := controllers.Users{}
	usersC.Templates.New = views.Must(views.ParseFS(templates.FS, "layout.gohtml", "signup.gohtml"))
	r.Get("/signup", usersC.New)
	r.Post("/users", usersC.Create)
	// page not found
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
