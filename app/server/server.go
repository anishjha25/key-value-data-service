package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"key-value-data-service/app/config"
	"key-value-data-service/app/external"
	"key-value-data-service/app/server/api"
	"log"
	"net/http"
	"time"
)

func NewServer() *http.Server {
	storage := external.NewStorageHTTPClient(config.AppConfig.StorageHost)
	h := api.NewHandler(storage)
	r := mux.NewRouter()

	r.HandleFunc("/api/kv/{key}", h.Set).Methods(http.MethodPut)
	r.HandleFunc("/api/kv/{key}", h.Get).Methods(http.MethodGet)
	r.HandleFunc("/api/kv/{key}", h.Delete).Methods(http.MethodDelete)

	r.Use(RequestLoggerMiddleware(r))

	return &http.Server{
		Addr:         fmt.Sprintf("%s:%d", config.AppConfig.Host, config.AppConfig.Port),
		Handler:      r,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
}

func RequestLoggerMiddleware(r *mux.Router) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			defer func() {
				log.Printf(
					"[%s] %s %s %s",
					req.Method,
					req.Host,
					req.URL.Path,
					req.URL.RawQuery,
				)
			}()

			next.ServeHTTP(w, req)
		})
	}
}
