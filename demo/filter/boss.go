package filter

import (
	"net/http"
	"runtime/debug"
	"time"

	"github.com/simplejia/clog/api"
)

func Boss(w http.ResponseWriter, r *http.Request, m map[string]interface{}) bool {
	defer func() {
		recover()
	}()

	err := m["__E__"]
	elapse := m["__T__"].(time.Time)

	// ...

	if err != nil {
		clog.Error("Boss() err: %v, stack: %s", err, debug.Stack())
	} else {
		clog.Info("Boss() elapse: %v", elapse)
	}
	return true
}
