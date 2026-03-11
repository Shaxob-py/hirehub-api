package project

import (
	"ShopAPI/internal/utils"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (h *ProjectHandle) CreateProject(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var inpProject ModelCreateProject

	if err := json.NewDecoder(r.Body).Decode(&inpProject); err != nil {
		utils.ResponseError(w, http.StatusBadRequest, "Bad Request")
		return
	}
	created, err := h.store.ProjectCreate(inpProject)
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, "Server Error")
		return
	}
	utils.ResponseWithJson(w, http.StatusCreated, created)

}

func (h *ProjectHandle) GetProject(w http.ResponseWriter, r *http.Request) {
	project, err := h.store.ProjectGetAll()

	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, "cannot fetch projects")
		return
	}
	utils.ResponseWithJson(w, http.StatusOK, project)
}

func (h *ProjectHandle) GetProjectById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	project, err := h.store.ProjectFindByID(id)
	if err != nil {
		utils.ResponseError(w, http.StatusNotFound, "Not Found")
		return
	}
	utils.ResponseWithJson(w, http.StatusOK, project)
}

func (h *ProjectHandle) UpdateProject(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var inp ModelUpdateProject
	id := chi.URLParam(r, "id")
	if err := json.NewDecoder(r.Body).Decode(&inp); err != nil {
		utils.ResponseError(w, http.StatusBadRequest, "Bad Request")
		return
	}
	inp.ID = id
	project, err := h.store.ProjectUpdate(inp)
	if err != nil {
		utils.ResponseError(w, http.StatusBadRequest, "Bad Request")
		return
	}
	utils.ResponseWithJson(w, http.StatusOK, project)
}

func (h *ProjectHandle) DeleteProject(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := h.store.ProjectDelete(id)
	if err != nil {
		utils.ResponseError(w, http.StatusNotFound, "Project not found")

		return
	}
	w.WriteHeader(http.StatusNoContent)
}
