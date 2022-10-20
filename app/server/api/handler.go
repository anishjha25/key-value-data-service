package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"key-value-data-service/app/dto"
	"key-value-data-service/app/external"
	"net/http"
	"strings"
)

type handler struct {
	storage external.StorageHTTPClient
}

func (h *handler) Set(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	key := strings.TrimSpace(params["key"])

	if key == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var kv dto.KeyValue

	err := json.NewDecoder(r.Body).Decode(&kv)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.storage.Set(key, kv.Value)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	key := strings.TrimSpace(params["key"])

	if key == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	value, err := h.storage.Get(key)
	if err != nil {
		if err.Error() == "not found" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp := dto.KeyValue{Value: value}
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	key := strings.TrimSpace(params["key"])

	if key == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := h.storage.Delete(key)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func NewHandler(store external.StorageHTTPClient) *handler {
	return &handler{
		storage: store,
	}
}
