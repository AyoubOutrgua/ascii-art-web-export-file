package functions

import (
	"fmt"
	"net/http"
)

func HandlerExport(w http.ResponseWriter, r *http.Request) {
	asciArtGenerated := r.FormValue("asciigen")
	if len(asciArtGenerated) == 0 {
		HandleError(w, "Bad Request!", http.StatusBadRequest)
		return
	}
	dataExport := []byte(asciArtGenerated)
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(dataExport)))
	w.Header().Set("Content-Disposition", "attachment; filename=\"asci-art.txt\"")
	w.Write(dataExport)
}
