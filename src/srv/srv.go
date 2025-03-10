package srv

import "net/http"

func DefaultServerSetup(sm *http.ServeMux) *http.Server {
	return &http.Server{
		Addr:    ":9090",
		Handler: sm,
	}
}

func StartServer() {
	TheSrv := DefaultServerSetup(http.NewServeMux())
	TheSrv.ListenAndServe()
}
