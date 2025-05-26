package controllers

import (
	"net/http"

	"github.com/jdlms/go-webapp/views"
)

func StaticHandler(tpl views.Template) http.HandlerFunc {
	return func( w http.ResponseWriter, r *http.Request){
		tpl.Execute(w, nil)
	}
}