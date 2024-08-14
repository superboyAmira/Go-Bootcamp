package handlers

import (
	"html/template"
	"net/http"
)

func AdminPanelHandler(w http.ResponseWriter, req *http.Request) {
	tmpl := template.Must(template.ParseFiles("../../web/templates/admin.html"))
	w.Header().Set("Content-Type", "text/html")
	tmpl.Execute(w, tmpl)
}
