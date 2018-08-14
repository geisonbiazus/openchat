package main

import (
	"net/http"

	"github.com/geisonbiazus/openchat/internal/openchat/router"
)

func main() {
	http.ListenAndServe(":8080", router.NewMux())
}
