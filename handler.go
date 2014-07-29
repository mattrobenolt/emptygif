package emptygif

import "net/http"

// An http.HandlerFunc for serving up an empty gif response
func Handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/gif")
	w.Write(PIXEL)
}
