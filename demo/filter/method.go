package filter

import (
	"net/http"
	"strings"
)

func Method(w http.ResponseWriter, r *http.Request, p map[string]interface{}) bool {
	method, ok := p["type"].(string)
	if ok && strings.ToLower(r.Method) != strings.ToLower(method) {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return false
	}
	return true
}
