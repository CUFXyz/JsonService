package srv

import (
	"jsonservice/src/jsfs"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	indexpage := jsfs.Htmlfile{Name: "index.html"}
	container := jsfs.ReadHTML(&indexpage)
	w.Write(container)
}

func DefaultServerSetup(sm *http.ServeMux) *http.Server {
	sm.HandleFunc("/", index)
	return &http.Server{
		Addr:    ":9090",
		Handler: sm,
	}
}

func StartServer() {
	TheSrv := DefaultServerSetup(http.NewServeMux())
	TheSrv.ListenAndServe()
}
