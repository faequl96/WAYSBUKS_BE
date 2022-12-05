package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	dto "waysbucks_api/dto/result"
	toppingdto "waysbucks_api/dto/topping"
	"waysbucks_api/models"
	"waysbucks_api/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
)

type handlerTopping struct {
	ToppingRepository repositories.ToppingRepository
}

func HandlerTopping(ToppingRepository repositories.ToppingRepository) *handlerTopping {
	return &handlerTopping{ToppingRepository}
}

func (h *handlerTopping) HandlerGetToppings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	toppings, err := h.ToppingRepository.RepoGetToppings()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	var ToppingResponses []toppingdto.ToppingResponse
	for _, topping := range toppings {
		ToppingResponse := toppingdto.ToppingResponse{
			ID:    topping.ID,
			Title: topping.Title,
			Price: topping.Price,
			Image: os.Getenv("PATH_FILE") + topping.Image,
		}

		ToppingResponses = append(ToppingResponses, ToppingResponse)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: ToppingResponses}
	json.NewEncoder(w).Encode(response)

}

func (h *handlerTopping) HandlerGetToppingByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	toppingFromDB, err := h.ToppingRepository.RepoGetToppingByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	ToppingResponse := toppingdto.ToppingResponse{
		ID:    toppingFromDB.ID,
		Title: toppingFromDB.Title,
		Price: toppingFromDB.Price,
		Image: os.Getenv("PATH_FILE") + toppingFromDB.Image,
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: ToppingResponse}
	json.NewEncoder(w).Encode(response)

}

func (h *handlerTopping) HandlerCreateTopping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userRole := userInfo["role"]
	// userId := int(userInfo["id"].(float64))

	if userRole != "admin" {
		w.WriteHeader(http.StatusUnauthorized)
		response := dto.ErrorResult{Code: http.StatusUnauthorized, Message: "only admin can add topping"}
		json.NewEncoder(w).Encode(response)
		return
	}

	dataContex := r.Context().Value("dataFile")
	filename := dataContex.(string)

	price, _ := strconv.Atoi(r.FormValue("price"))
	request := toppingdto.ToppingRequest{
		Title: r.FormValue("title"),
		Price: price,
		Image: filename,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	topping := models.Topping{
		Title: request.Title,
		Price: request.Price,
		Image: filename,
	}

	toppingReturn, err := h.ToppingRepository.RepoCreateTopping(topping)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	ToppingResponse := toppingdto.ToppingResponse{
		ID:    toppingReturn.ID,
		Title: toppingReturn.Title,
		Price: toppingReturn.Price,
		Image: os.Getenv("PATH_FILE") + toppingReturn.Image,
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: ToppingResponse}
	json.NewEncoder(w).Encode(response)

}

func (h *handlerTopping) HandlerUpdateTopping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userRole := userInfo["role"]

	if userRole != "admin" {
		w.WriteHeader(http.StatusUnauthorized)
		response := dto.ErrorResult{Code: http.StatusUnauthorized, Message: "only admin can edit topping"}
		json.NewEncoder(w).Encode(response)
		return
	}

	dataContex := r.Context().Value("dataFile")
	filename := dataContex.(string)

	price, _ := strconv.Atoi(r.FormValue("price"))
	request := toppingdto.ToppingRequest{
		Title: r.FormValue("title"),
		Price: price,
		Image: filename,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	toppingFromDB, err := h.ToppingRepository.RepoGetToppingByID(int(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	if request.Title != "" {
		toppingFromDB.Title = request.Title
	}

	if request.Price != 0 {
		toppingFromDB.Price = request.Price
	}

	if request.Image != "empty" {
		toppingFromDB.Image = request.Image
	}

	toppingReturn, err := h.ToppingRepository.RepoUpdateTopping(toppingFromDB)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	ToppingResponse := toppingdto.ToppingResponse{
		ID:    toppingReturn.ID,
		Title: toppingReturn.Title,
		Price: toppingReturn.Price,
		Image: os.Getenv("PATH_FILE") + toppingReturn.Image,
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: ToppingResponse}
	json.NewEncoder(w).Encode(response)

}

func (h *handlerTopping) HandlerDeleteTopping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userRole := userInfo["role"]

	if userRole != "admin" {
		w.WriteHeader(http.StatusUnauthorized)
		response := dto.ErrorResult{Code: http.StatusUnauthorized, Message: "only admin can delete topping"}
		json.NewEncoder(w).Encode(response)
		return
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	toppingFromDB, err := h.ToppingRepository.RepoGetToppingByID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	toppingReturn, err := h.ToppingRepository.RepoDeleteTopping(toppingFromDB)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	ToppingDeleteResponse := toppingdto.ToppingDeleteResponse{
		ID:    toppingReturn.ID,
		Title: toppingReturn.Title,
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: ToppingDeleteResponse}
	json.NewEncoder(w).Encode(response)

}
