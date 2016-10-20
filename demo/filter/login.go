package filter

import (
	"net/http"

	"github.com/simplejia/wsp/demo/controller"
)

func Login(w http.ResponseWriter, r *http.Request, p map[string]interface{}) bool {
	c := p["__C__"].(controller.IBase)

	// ...

	c.SetParam("user", "user_info")

	return true
}
