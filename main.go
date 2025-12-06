package main

import (
	"log"
	"net/http"

	"github.com/srthk29/pdf-preview/internal/api"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/pdfpreview", api.PdfPreview)

	log.Println("server running on http://localhost:9999")
	if err := http.ListenAndServe(":9999", mux); err != nil {
		log.Fatal(err)
	}
}
