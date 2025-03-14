package srv

import (
	"jsonservice/src/jsfs"
	"net/http"
)

func DefaultServerSetup(sm *http.ServeMux) *http.Server {
	return &http.Server{
		Addr:    ":9090",
		Handler: sm,
	}
}

func StartServer() {
	TheSrv := DefaultServerSetup(http.NewServeMux())
	file := jsfs.Htmlfile{
		Name:   "index",
		Format: "html",
		Css:    "",
		Js:     "",
	}
	jsfs.CheckPage(&file)
	TheSrv.ListenAndServe()
}
