package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	cartdto "waysbucks_api/dto/cart"
	dto "waysbucks_api/dto/result"
	"waysbucks_api/models"
	"waysbucks_api/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
)

type handlerCart struct {
	CartRepository repositories.CartRepository
}

func HandlerCart(CartRepository repositories.CartRepository) *handlerCart {
	return &handlerCart{CartRepository}
}

func (h *handlerCart) HandlerAddCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	request := new(cartdto.AddCartRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	validation := validator.New()
	err := validation.Struct(request)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "error validation"}
		json.NewEncoder(w).Encode(response)
		return
	}

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	customerID := int(userInfo["id"].(float64))

	product, err := h.CartRepository.RepoGetProductCart(request.ProductID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "Product Not Found!"}
		json.NewEncoder(w).Encode(response)
		return
	}

	toppings, err := h.CartRepository.RepoGetToppingCart(request.ToppingsID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "Toping Not Found!"}
		json.NewEncoder(w).Encode(response)
		return
	}

	dataCart := models.Cart{
		CustomerID: customerID,
		ProductID:  product.ID,
		Toppings:   toppings,
		Price:      request.Price,
	}

	err = h.CartRepository.RepoAddCart(dataCart)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: "Failed to add cart!"}
		json.NewEncoder(w).Encode(response)
		return
	}
}

func (h *handlerCart) HandlerDeleteCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	cart, err := h.CartRepository.RepoGetCart(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.CartRepository.RepoDelCart(cart)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerCart) HandlerGetCartByIdUserIsPayedFalse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	customerID := int(userInfo["id"].(float64))

	// fmt.Println(customerID)

	cartFromDB, err := h.CartRepository.RepoGetCartByIdUserIsPayedFalse(customerID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	var cartResponse []models.CartResponse
	for _, cart := range cartFromDB {
		cartRes := models.CartResponse{
			ID:       cart.ID,
			Product:  cart.Product,
			Toppings: cart.Toppings,
			Price:    cart.Price,
		}
		cartResponse = append(cartResponse, cartRes)
	}

	fmt.Println(len(cartResponse))

	if len(cartResponse) > 0 {
		w.WriteHeader(http.StatusOK)
		response := dto.SuccessResult{Code: http.StatusOK, Data: cartResponse}
		json.NewEncoder(w).Encode(response)
	} else {
		w.WriteHeader(http.StatusOK)
		response := dto.SuccessResult{Code: http.StatusOK, Data: 0}
		json.NewEncoder(w).Encode(response)
	}
}

func (h *handlerCart) HandlerUpdateCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(cartdto.UpdateCartRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	validation := validator.New()
	err := validation.Struct(request)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "error validation"}
		json.NewEncoder(w).Encode(response)
		return
	}

	err = h.CartRepository.RepoUpdateCart(request.ID, request.IsPayed, request.TransDay, request.TransTime)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	customerID := int(userInfo["id"].(float64))

	cartFromDB, err := h.CartRepository.RepoGetCartByIdUserIsPayedFalse(customerID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	if len(cartFromDB) > 0 {
		w.WriteHeader(http.StatusOK)
		response := dto.SuccessResult{Code: http.StatusOK, Data: cartFromDB}
		json.NewEncoder(w).Encode(response)
	} else {
		w.WriteHeader(http.StatusOK)
		response := dto.SuccessResult{Code: http.StatusOK, Data: 0}
		json.NewEncoder(w).Encode(response)
	}

}
