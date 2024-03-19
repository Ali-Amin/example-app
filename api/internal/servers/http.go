package servers

import (
	"fmt"
	"net/http"
)

type HTTPServer struct {
	Port int
}

func NewHTTPServer(port int) *HTTPServer {
	return &HTTPServer{Port: port}
}

func (h *HTTPServer) ServeAndListen(router http.Handler) error {
	err := http.ListenAndServe(
		fmt.Sprintf("0.0.0.0:%d", h.Port),
		router,
	)
	if err != nil {
		return err
	}
	return nil
}
