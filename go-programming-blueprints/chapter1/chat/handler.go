package main

import "net/http"

// HandlerWithData to add data to handler
type HandlerWithData interface {
	SetData(map[string]interface{})
	http.Handler
}
