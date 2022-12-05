package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	dto "waysbucks_api/dto/result"
	userdto "waysbucks_api/dto/user"
	"waysbucks_api/pkg/bcrypt"
	"waysbucks_api/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
)

type handlerUser struct {
	UserRepository repositories.UserRepository
}

func HandlerUser(UserRepository repositories.UserRepository) *handlerUser {
	return &handlerUser{UserRepository}
}

func (h *handlerUser) HandlerGetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userRole := userInfo["role"]

	if userRole != "admin" {
		w.WriteHeader(http.StatusUnauthorized)
		response := dto.ErrorResult{Code: http.StatusUnauthorized, Message: "To show all data user, you must login as Admin"}
		json.NewEncoder(w).Encode(response)
		return
	}

	usersFromDB, err := h.UserRepository.RepoGetUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	var UserResponses []userdto.UserResponse
	for _, user := range usersFromDB {
		UserResponse := userdto.UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
			Image: os.Getenv("PATH_FILE") + user.Image,
		}

		UserResponses = append(UserResponses, UserResponse)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: UserResponses}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerUser) HandlerGetUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userRole := userInfo["role"]
	userId := int(userInfo["id"].(float64))

	if userId != id && userRole != "admin" {
		w.WriteHeader(http.StatusUnauthorized)
		response := dto.ErrorResult{Code: http.StatusUnauthorized, Message: "To show this data user, you must login as Admin or login with this user account!"}
		json.NewEncoder(w).Encode(response)
		return
	}

	userFromDB, err := h.UserRepository.RepoGetUserByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	UserResponse := userdto.UserResponse{
		ID:    userFromDB.ID,
		Name:  userFromDB.Name,
		Email: userFromDB.Email,
		Image: os.Getenv("PATH_FILE") + userFromDB.Image,
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: UserResponse}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerUser) HandlerUpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dataContex := r.Context().Value("dataFile")
	filename := dataContex.(string)

	request := userdto.UserUpdateRequest{
		Name:     r.FormValue("name"),
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
		Image:    filename,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	passwordHashed, err := bcrypt.HashingPassword(request.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userRole := userInfo["role"]
	userId := int(userInfo["id"].(float64))

	if userId != id && userRole != "admin" {
		w.WriteHeader(http.StatusUnauthorized)
		response := dto.ErrorResult{Code: http.StatusUnauthorized, Message: "You can't edit this data user, you must login as Admin or login with this user account!"}
		json.NewEncoder(w).Encode(response)
		return
	}

	user, err := h.UserRepository.RepoGetUserByID(int(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	if request.Name != "" {
		user.Name = request.Name
	}

	if request.Email != "" {
		user.Email = request.Email
	}

	if request.Password != "" {
		user.Password = passwordHashed
	}

	if request.Image != "empty" {
		user.Image = request.Image
	}

	data, err := h.UserRepository.RepoUpdateUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	UserResponse := userdto.UserUpdateResponse{
		ID:       data.ID,
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
		Image:    os.Getenv("PATH_FILE") + data.Image,
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: UserResponse}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerUser) HandlerDeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userRole := userInfo["role"]
	userId := int(userInfo["id"].(float64))

	if userId != id && userRole != "admin" {
		w.WriteHeader(http.StatusUnauthorized)
		response := dto.ErrorResult{Code: http.StatusUnauthorized, Message: "You can't delete this user account, you must login as Admin or login with this user account!"}
		json.NewEncoder(w).Encode(response)
		return
	}

	userFromDB, err := h.UserRepository.RepoGetUserByID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	userReturn, err := h.UserRepository.RepoDeleteUser(userFromDB)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	UserDeleteResponse := userdto.UserDeleteResponse{
		ID:   userReturn.ID,
		Name: userReturn.Name,
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: UserDeleteResponse}
	json.NewEncoder(w).Encode(response)
}
