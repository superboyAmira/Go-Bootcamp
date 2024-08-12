package handlers

import (
	"net/http"
	"os"

	"github.com/russross/blackfriday/v2"
)

func AdminPanelHandler(w http.ResponseWriter, req *http.Request) {
	form, err := os.ReadFile("../../../web/templates/AdminPanel.md")
	if err != nil {
		http.Error(w, "Cannot open template AdminPanel", http.StatusInternalServerError)
	}
	html := blackfriday.Run(form)

	w.Header().Set("Content-Type", "text/html")
	w.Write(html)
}