package router

import (
	"github.com/alissonphp/go-update-file/internal/application/handler"
	"log"
	"net/http"
)

func SetupWebRouter() {

	mux := http.NewServeMux()
	mux.HandleFunc("/images", handler.ListImages)

	err := http.ListenAndServe(":8083", mux)
	if err != nil {
		log.Fatalf(err.Error())
	}
}
