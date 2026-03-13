package user

import (
	"ShopAPI/internal/utils"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// CreateUserHandler godoc
// Register user
// @Summary Register new user
// @Tags auth
// @Accept json
// @Produce json
// @Param user body ModelCreateUser true "User data"
// @Success 201 {object} ModelCreateUser
// @Failure 400 {object} map[string]string
// @Router /api/v1/auth/register [post]
func (h *UserHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
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

// LoginUserHandler godoc
// Login user
// @Summary Login user
// @Tags auth
// @Accept json
// @Produce json
// @Param user body ModelUserLogin true "Login"
// @Success 200 {string} string
// @Router /api/v1/auth/login [post]
func (h *UserHandler) LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user ModelUserLogin
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.ResponseError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := h.store.CheckUserEmail(user); err != nil {
		utils.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	tokenString, err := utils.CreateToken(user.Email)
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	utils.ResponseWithJson(w, http.StatusOK, tokenString)
	return
}

// UpdateUserHandler godoc
// @Summary Update user
// @Description Update user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param user body ModelUpdateUser true "Updated user data"
// @Success 200 {object} ModelUser
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/v1/users/{id} [put]
func (h *UserHandler) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	var inp ModelUpdateUser
	id := chi.URLParam(r, "id")
	if err := json.NewDecoder(r.Body).Decode(&inp); err != nil {
		utils.ResponseError(w, http.StatusBadRequest, "Bad Request")
		return
	}
	task, err := h.store.UpdateUser(inp, id)
	if err != nil {
		utils.ResponseError(w, http.StatusNotFound, "Not Found")
		return
	}
	utils.ResponseWithJson(w, http.StatusOK, task)
}

// DeleteUserHandler godoc
// @Summary Delete user
// @Description Delete user by ID
// @Tags users
// @Produce json
// @Param id path string true "User ID"
// @Success 204
// @Failure 404 {object} map[string]string
// @Router /api/v1/users/{id} [delete]
func (h *UserHandler) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	err := h.store.DeleteUser(idParam)
	if err != nil {
		utils.ResponseError(w, http.StatusNotFound, err.Error())
		return
	}

	utils.ResponseWithJson(w, http.StatusNoContent, nil)
}

// GetUserHandler godoc
// @Summary Get user
// @Description Get user by ID
// @Tags users
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} ModelUser
// @Failure 404 {object} map[string]string
// @Router /api/v1/users/{id} [get]
func (h *UserHandler) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	user, err := h.store.GetUserById(idParam)
	if err != nil {
		utils.ResponseError(w, http.StatusNotFound, err.Error())
		return
	}

	utils.ResponseWithJson(w, http.StatusOK, user)
}

// GetAllUsersHandler godoc
// @Summary Get all users
// @Description Get list of all users
// @Tags users
// @Produce json
// @Success 200 {array} ModelUser
// @Failure 500 {object} map[string]string
// @Router /api/v1/users [get]
func (h *UserHandler) GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := h.store.GetAllUsers("string")
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.ResponseWithJson(w, http.StatusOK, users)
}
