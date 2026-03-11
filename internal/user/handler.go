package user

import (
	"ShopAPI/internal/utils"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// CreateUser godoc
// @Summary Create user
// @Description Create a new user
// @Tags users
// @Accept json
// @Produce json
// @Param user body ModelCreateUser true "User data"
// @Success 201 {object} ModelUser
// @Router /users [post]
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var inp ModelCreateUser
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&inp); err != nil {
		utils.ResponseError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	err := CheckUser(&inp)
	if err != nil {
		utils.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	task, err := h.store.CreateUser(inp)
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.ResponseWithJson(w, http.StatusCreated, task)

}

// UpdateUser godoc
// @Summary Update user
// @Description Update user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param user body ModelUpdateUser true "Updated user data"
// @Success 200 {object} ModelUpdateUser
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /users/{id} [put]
func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var inp ModelUpdateUser
	id := chi.URLParam(r, "id")
	if err := json.NewDecoder(r.Body).Decode(&inp); err != nil {
		utils.ResponseError(w, http.StatusBadRequest, "Bad Request")
		return
	}
	task, err := h.store.UpdateUser(inp, id)
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.ResponseWithJson(w, http.StatusOK, task)
}

// DeleteUser godoc
// @Summary Delete user
// @Description Delete user by ID
// @Tags users
// @Produce json
// @Param id path string true "User ID"
// @Success 204
// @Failure 404 {object} map[string]string
// @Router /users/{id} [delete]
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	err := h.store.DeleteUser(idParam)
	if err != nil {
		utils.ResponseError(w, http.StatusNotFound, err.Error())
		return
	}
	utils.ResponseWithJson(w, http.StatusNoContent, nil)
}

// GetUser godoc
// @Summary Get user
// @Description Get user by ID
// @Tags users
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} ModelUser
// @Failure 404 {object} map[string]string
// @Router /users/{id} [get]
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	user, err := h.store.GetUserById(idParam)
	if err != nil {
		utils.ResponseError(w, http.StatusNotFound, err.Error())
		return
	}
	utils.ResponseWithJson(w, http.StatusOK, user)
}

// GetAllUsers godoc
// @Summary Get all users
// @Description Get list of all users
// @Tags users
// @Produce json
// @Success 200 {array} ModelUser
// @Failure 500 {object} map[string]string
// @Router /users [get]
func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.store.GetAllUsers("")
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.ResponseWithJson(w, http.StatusOK, users)
}
