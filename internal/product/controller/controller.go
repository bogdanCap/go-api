package controller

import (
	"encoding/json"
	"net/http"
	//"github.com/bogdanCap/go-api/internal/application"
	"github.com/bogdanCap/go-api/internal/product/application/service"
	"github.com/bogdanCap/go-api/internal/product/application/port"
	"github.com/bogdanCap/go-api/internal/product/validation"
	"github.com/bogdanCap/go-api/internal/product/dto"
	"github.com/bogdanCap/go-api/internal/product/mapper"
	"github.com/bogdanCap/go-api/internal/shared/pagination"
	//"github.com/bogdanCap/go-api/internal/port/service"
	"github.com/bogdanCap/go-api/internal/response"

	"github.com/go-playground/validator/v10"
)

type ProductController struct {
	service port.IProductService
}

func NewProductController(service service.ProductService) ProductController {
	return ProductController{
		service: service,
	}
}

func (c *ProductController) List(w http.ResponseWriter, r *http.Request) {

	var validate = validator.New()

	query := r.URL.Query()

	var req dto.ProductListFilterDTO

	// name
	if value := query.Get("name"); value != "" {
		req.Name = &value
	}

	limitRow := query.Get("limit")
	req.Limit = pagination.Limit(&limitRow)

	//if offsetRow := query.Get("offset"); offsetRow != "" {
	offsetRow := query.Get("offset")
	req.Offset = pagination.Offset(&offsetRow)

	// validation
	if err := validate.Struct(req); err != nil {
		response.WriteValidationError(
			w,
			err,
			validation.ProductValidationMessages,
		)

		return
	}

	product, err := c.service.List(r.Context(), req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	productsResponseDTO := mapper.ProductsToResponse(product)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(productsResponseDTO)
}
