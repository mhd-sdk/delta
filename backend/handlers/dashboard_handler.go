package handlers

import (
	"delta/models"
	"delta/repositories"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type DashboardHandler struct {
	repo *repositories.DashboardRepository
}

func NewDashboardHandler(repo *repositories.DashboardRepository) *DashboardHandler {
	return &DashboardHandler{repo: repo}
}

func (h *DashboardHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dashboard models.Dashboard
	if err := json.NewDecoder(r.Body).Decode(&dashboard); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.repo.Create(&dashboard); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(dashboard)
}

func (h *DashboardHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	dashboard, err := h.repo.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dashboard)
}

func (h *DashboardHandler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	var dashboard models.Dashboard
	if err := json.NewDecoder(r.Body).Decode(&dashboard); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	dashboard.ID = id
	if err := h.repo.Update(&dashboard); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dashboard)
}

func (h *DashboardHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	if err := h.repo.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *DashboardHandler) UpdatePanel(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	dashboardID, err := uuid.Parse(vars["id"])
	if err != nil {
		http.Error(w, "invalid dashboard id", http.StatusBadRequest)
		return
	}

	var panel models.Panel
	if err := json.NewDecoder(r.Body).Decode(&panel); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.repo.UpdatePanel(dashboardID, &panel); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(panel)
}

func (h *DashboardHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	dashboards, err := h.repo.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dashboards)
}
