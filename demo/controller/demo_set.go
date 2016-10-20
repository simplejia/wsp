package controller

import (
	"net/http"
)

// @prefilter("Login", {"Method":{"type":"post"}})
// @postfilter("Boss")
func (demo *Demo) Set(w http.ResponseWriter, r *http.Request) {
}
