package handlers

import "net/http"

type RESTHandler interface {
	Handle(http.ResponseWriter, *http.Request) error
}