package pkg

import (
	"fmt"
	"net/http"
	"time"
)

func SetupServer(port int, handler http.Handler) http.Server {
	server := http.Server{
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		Addr:         fmt.Sprintf("0.0.0.0:%d", port),
	}
	return server
}
