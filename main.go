package main

import (
	"log"
	"net/http"

	"github.com/davidbyttow/govips/v2/vips"
	"github.com/srthk29/pdf-preview/internal/api"
)

func main() {
	vips.Startup(nil)
	defer vips.Shutdown()

	mux := http.NewServeMux()
	mux.HandleFunc("/pdfpreview", api.PdfPreview)

	log.Println("server running on http://localhost:9999")
	if err := http.ListenAndServe(":9999", mux); err != nil {
		log.Fatal(err)
	}
}
