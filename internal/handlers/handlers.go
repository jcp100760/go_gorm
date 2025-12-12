package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jcp100760/api_go/internal/repository"
	"github.com/jcp100760/api_go/internal/service"
	"strconv"
)

// Handlers agrupa dependencias usadas por los controladores HTTP.
type Handlers struct {
	svc service.Service
}

// NewHandlers crea los handlers inyectando el servicio.
func NewHandlers(s service.Service) *Handlers {
	return &Handlers{svc: s}
}

// RegisterRoutes registra las rutas en el router provisto.
func (h *Handlers) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/api/saludo", h.saludoHandler).Methods("GET")
	r.HandleFunc("/api/clientes", h.getClientesHandler).Methods("GET")
	r.HandleFunc("/api/addCliente", h.addClienteHandler).Methods("POST")
}

// saludoHandler responde con un mensaje simple.
func (h *Handlers) saludoHandler(w http.ResponseWriter, r *http.Request) {
	resp := map[string]string{"mensaje": "¡Hola! Este es el endpoint /api/saludo"}
	writeJSON(w, http.StatusOK, resp)
}

// getClientesHandler obtiene la lista de clientes desde el servicio.
func (h *Handlers) getClientesHandler(w http.ResponseWriter, r *http.Request) {
	clients, err := h.svc.GetAllClients()
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, clients)
}

// addClienteHandler parsea el body y solicita al servicio guardar el cliente.
func (h *Handlers) addClienteHandler(w http.ResponseWriter, r *http.Request) {
	var req repository.Client
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "body inválido"})
		return
	}

	created, err := h.svc.AddClient(req)
	if err != nil {
		// Podríamos mapear errores específicos (p.ej., validaciones) a status codes
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusCreated, created)
}

// helper to write JSON responses
func writeJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
