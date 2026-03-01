package user

import (
	"ShopAPI/internal/utils"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strings"
)

type UserHandler struct {
	store *User
}

func NewUserHandler(store *User) *UserHandler {
	return &UserHandler{
		store: store,
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var inp ModelCreateUser
	if err := json.NewDecoder(r.Body).Decode(&inp); err != nil {
		utils.ResponseError(w, http.StatusBadRequest, "Bad Request")
		return
	}

	if inp.Skill == "" || inp.About == "" || !strings.HasSuffix(inp.Email, "gmail.com") {
		utils.ResponseError(w, http.StatusBadRequest, "Bad Request")
		return
	}

	task, err := h.store.CreateUser(inp)
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.ResponseWithJson(w, http.StatusCreated, task)

}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var inp ModelUpdateUser
	if err := json.NewDecoder(r.Body).Decode(&inp); err != nil {
		utils.ResponseError(w, http.StatusBadRequest, "Bad Request")
		return
	}
	task, err := h.store.UpdateUser(inp)
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.ResponseWithJson(w, http.StatusOK, task)
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	err := h.store.DeleteUser(idParam)
	if err != nil {
		utils.ResponseError(w, http.StatusNotFound, err.Error())
		return
	}
	utils.ResponseWithJson(w, http.StatusNoContent, nil)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	user, err := h.store.GetUserById(idParam)
	if err != nil {
		utils.ResponseError(w, http.StatusNotFound, err.Error())
		return
	}
	utils.ResponseWithJson(w, http.StatusOK, user)
}

func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.store.GetAllUsers("")
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.ResponseWithJson(w, http.StatusOK, users)
}
