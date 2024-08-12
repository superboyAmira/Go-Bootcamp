package handlers

import (
	"net/http"
	"os"

	"github.com/russross/blackfriday/v2"
)

func MainIndexHandler(w http.ResponseWriter, req *http.Request) {
	form, err := os.ReadFile("../../../web/templates/Index.md")
	if err != nil {
		http.Error(w, "Cannot open template Index", http.StatusInternalServerError)
	}
	html := blackfriday.Run(form)

	w.Header().Set("Content-Type", "text/html")
	w.Write(html)
}