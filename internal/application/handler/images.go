package handler

import (
	"net/http"
)

func ListImages(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
	return
}
