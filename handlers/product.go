package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	productdto "waysbucks_api/dto/product"
	dto "waysbucks_api/dto/result"
	"waysbucks_api/models"
	"waysbucks_api/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
)

type handlerProduct struct {
	ProductRepository repositories.ProductRepository
}

func HandlerProduct(ProductRepository repositories.ProductRepository) *handlerProduct {
	return &handlerProduct{ProductRepository}
}

func (h *handlerProduct) HandlerGetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	products, err := h.ProductRepository.RepoGetProducts()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	var ProductResponses []productdto.ProductResponse
	for _, product := range products {
		ProductResponse := productdto.ProductResponse{
			ID:    product.ID,
			Title: product.Title,
			Price: product.Price,
			Image: product.Image,
		}

		ProductResponses = append(ProductResponses, ProductResponse)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: ProductResponses}
	json.NewEncoder(w).Encode(response)

}

func (h *handlerProduct) HandlerGetProductByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	productFromDB, err := h.ProductRepository.RepoGetProductByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	ProductResponse := productdto.ProductResponse{
		ID:    productFromDB.ID,
		Title: productFromDB.Title,
		Price: productFromDB.Price,
		Image: productFromDB.Image,
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: ProductResponse}
	json.NewEncoder(w).Encode(response)

}

func (h *handlerProduct) HandlerCreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userRole := userInfo["role"]

	if userRole != "admin" {
		w.WriteHeader(http.StatusUnauthorized)
		response := dto.ErrorResult{Code: http.StatusUnauthorized, Message: "only admin can add product"}
		json.NewEncoder(w).Encode(response)
		return
	}

	dataContex := r.Context().Value("dataFile")
	filename := dataContex.(string)

	price, _ := strconv.Atoi(r.FormValue("price"))
	request := productdto.ProductRequest{
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

	product := models.Product{
		Title: request.Title,
		Price: request.Price,
		Image: os.Getenv("PATH_FILE") + filename,
	}

	productReturn, err := h.ProductRepository.RepoCreateProduct(product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	ProductResponse := productdto.ProductResponse{
		ID:    productReturn.ID,
		Title: productReturn.Title,
		Price: productReturn.Price,
		Image: productReturn.Image,
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: ProductResponse}
	json.NewEncoder(w).Encode(response)

}

func (h *handlerProduct) HandlerUpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userRole := userInfo["role"]

	if userRole != "admin" {
		w.WriteHeader(http.StatusUnauthorized)
		response := dto.ErrorResult{Code: http.StatusUnauthorized, Message: "only admin can edit product"}
		json.NewEncoder(w).Encode(response)
		return
	}

	dataContex := r.Context().Value("dataFile")
	filename := dataContex.(string)

	price, _ := strconv.Atoi(r.FormValue("price"))
	request := productdto.ProductRequest{
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

	productFromDB, err := h.ProductRepository.RepoGetProductByID(int(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	if request.Title != "" {
		productFromDB.Title = request.Title
	}

	if request.Price != 0 {
		productFromDB.Price = request.Price
	}

	if request.Image != "empty" {
		productFromDB.Image = os.Getenv("PATH_FILE") + request.Image
	}

	productReturn, err := h.ProductRepository.RepoUpdateProduct(productFromDB)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	ProductResponse := productdto.ProductResponse{
		ID:    productReturn.ID,
		Title: productReturn.Title,
		Price: productReturn.Price,
		Image: productReturn.Image,
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: ProductResponse}
	json.NewEncoder(w).Encode(response)

}

func (h *handlerProduct) HandlerDeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userRole := userInfo["role"]

	if userRole != "admin" {
		w.WriteHeader(http.StatusUnauthorized)
		response := dto.ErrorResult{Code: http.StatusUnauthorized, Message: "only admin can delete product"}
		json.NewEncoder(w).Encode(response)
		return
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	productFromDB, err := h.ProductRepository.RepoGetProductByID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	productReturn, err := h.ProductRepository.RepoDeleteProduct(productFromDB)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	ProductDeleteResponse := productdto.ProductDeleteResponse{
		ID:    productReturn.ID,
		Title: productReturn.Title,
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: ProductDeleteResponse}
	json.NewEncoder(w).Encode(response)

}
