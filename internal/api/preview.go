package api

import (
	"io"
	"log"
	"net/http"

	"github.com/davidbyttow/govips/v2/vips"
)

func init() {
	vips.LoggingSettings(func(domain string, level vips.LogLevel, msg string) {
		log.Println(domain, level, msg)
	}, vips.LogLevelDebug)

	// Disable the cache so that after GC, libvips does not hold reference to any object
	vips.Startup(&vips.Config{
		ConcurrencyLevel: 0,
		MaxCacheFiles:    0,
		MaxCacheMem:      0,
		MaxCacheSize:     0,
		ReportLeaks:      false,
		CacheTrace:       false,
		CollectStats:     false,
	})
}

func PdfPreview(w http.ResponseWriter, r *http.Request) {
	previewBytes, err := previewImage(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(previewBytes)
}

func previewImage(reader io.Reader) ([]byte, error) {
	img, err := vips.NewImageFromReader(reader)
	if err != nil {
		return nil, err
	}

	ep := vips.NewJpegExportParams()
	ep.StripMetadata = true
	ep.Quality = 75
	ep.Interlace = true
	ep.OptimizeCoding = true
	ep.SubsampleMode = vips.VipsForeignSubsampleAuto
	ep.TrellisQuant = true
	ep.OvershootDeringing = true
	ep.OptimizeScans = true
	ep.QuantTable = 3

	imageBytes, _, err := img.ExportJpeg(ep)
	if err != nil {
		return nil, err
	}

	return imageBytes, nil
}
