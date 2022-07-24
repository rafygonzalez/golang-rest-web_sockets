package handlers

import (
	"encoding/json"
	"gows/models"
	"gows/repository"
	"gows/server"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
	"github.com/segmentio/ksuid"
)

type UpsertProductRequest struct {
	ProductName string `json:"product_name"`
}

type ProductResponse struct {
	Id          string `json:"id"`
	ProductName string `json:"product_name"`
}

type ProductUpdateResponse struct {
	Message string `json:"message"`
}

func InsertProductHandler(s server.Server) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		tokenString := strings.TrimSpace(r.Header.Get("Authorization"))
		token, err := jwt.ParseWithClaims(tokenString, &models.AppClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(s.Config().JWTSecret), nil
		})
		if err != nil {
			http.Error(rw, err.Error(), http.StatusUnauthorized)
			return
		}
		if claims, ok := token.Claims.(*models.AppClaims); ok && token.Valid {
			productRequest := UpsertProductRequest{}
			if err := json.NewDecoder(r.Body).Decode(&productRequest); err != nil {
				http.Error(rw, err.Error(), http.StatusBadRequest)
			}
			id, err := ksuid.NewRandom()
			if err != nil {
				http.Error(rw, err.Error(), http.StatusInternalServerError)
			}
			product := models.Product{
				Id:          id.String(),
				ProductName: productRequest.ProductName,
				CreatedBy:   claims.UserId,
			}
			err = repository.InsertProduct(r.Context(), &product)
			if err != nil {
				http.Error(rw, err.Error(), http.StatusInternalServerError)
			}
			rw.Header().Set("Content-Type", "application/json")
			json.NewEncoder(rw).Encode(ProductResponse{
				Id:          product.Id,
				ProductName: product.ProductName,
			})
		} else {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func GetProductByIdHandler(s server.Server) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		product, err := repository.GetProductById(r.Context(), params["id"])
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		rw.Header().Set("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(product)
	}
}

func UpdateProductHandler(s server.Server) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		tokenString := strings.TrimSpace(r.Header.Get("Authorization"))
		token, err := jwt.ParseWithClaims(tokenString, &models.AppClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(s.Config().JWTSecret), nil
		})
		if err != nil {
			http.Error(rw, err.Error(), http.StatusUnauthorized)
			return
		}
		if claims, ok := token.Claims.(*models.AppClaims); ok && token.Valid {
			productRequest := UpsertProductRequest{}
			if err := json.NewDecoder(r.Body).Decode(&productRequest); err != nil {
				http.Error(rw, err.Error(), http.StatusBadRequest)
			}
			params := mux.Vars(r)
			product := models.Product{
				Id:          params["id"],
				ProductName: productRequest.ProductName,
				CreatedBy:   claims.UserId,
			}
			err = repository.UpdateProduct(r.Context(), &product)
			if err != nil {
				http.Error(rw, err.Error(), http.StatusInternalServerError)
			}
			rw.Header().Set("Content-Type", "application/json")
			json.NewEncoder(rw).Encode(ProductUpdateResponse{
				Message: "Product Updated",
			})
		} else {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func DeleteProductHandler(s server.Server) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		tokenString := strings.TrimSpace(r.Header.Get("Authorization"))
		token, err := jwt.ParseWithClaims(tokenString, &models.AppClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(s.Config().JWTSecret), nil
		})
		if err != nil {
			http.Error(rw, err.Error(), http.StatusUnauthorized)
			return
		}
		if _, ok := token.Claims.(*models.AppClaims); ok && token.Valid {
			params := mux.Vars(r)
			err = repository.DeleteProductById(r.Context(), params["id"])
			if err != nil {
				http.Error(rw, err.Error(), http.StatusInternalServerError)
			}
			rw.Header().Set("Content-Type", "application/json")
			json.NewEncoder(rw).Encode(ProductUpdateResponse{
				Message: "Product Deleted",
			})
		} else {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
