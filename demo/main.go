//go:generate wsp

package main

import (
	"fmt"

	"github.com/simplejia/clog/api"
	"github.com/simplejia/lc"

	"net/http"

	"github.com/simplejia/wsp/demo/conf"
)

func init() {
	lc.Init(1e5)

	clog.Init(conf.C.Clog.Name, "", conf.C.Clog.Level, conf.C.Clog.Mode)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})
}

func main() {
	clog.Info("main()")

	s := &http.Server{
		Addr: fmt.Sprintf("%s:%d", conf.C.App.Ip, conf.C.App.Port),
	}
	err := s.ListenAndServe()
	clog.Error("main() s.ListenAndServe %v", err)
}
