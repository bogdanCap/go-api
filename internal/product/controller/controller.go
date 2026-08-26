package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	//"github.com/bogdanCap/go-api/internal/application"
	"github.com/bogdanCap/go-api/internal/product/application/service"
	"github.com/bogdanCap/go-api/internal/product/application/port"
	"github.com/bogdanCap/go-api/internal/product/validation"
	"github.com/bogdanCap/go-api/internal/product/dto"
	"github.com/bogdanCap/go-api/internal/product/mapper"
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

	if value := query.Get("age"); value != "" {
		age, err := strconv.Atoi(value)
		if err != nil {
			http.Error(w, "age must be integer", http.StatusBadRequest)
			return
		}

		req.Age = &age
	}

	// name
	if value := query.Get("name"); value != "" {
		req.Name = &value
	}

	// validation
	if err := validate.Struct(req); err != nil {
		response.WriteValidationError(
			w,
			err,
			validation.ProductValidationMessages,
		)

		return
	}

	//TODO in this place dto need to map to domain layer
	//this is example and mapper for FILTER need to add in future
	//order := mapper.CreateOrderRequestToDomain(
	//		request,
	//	)
	//productDomain := mapper.ProductToDomain(req)

	//TODO in list controller response look like this
	/*
		esponse := mapper.OrdersToResponse(
			orders,
		)


		json.NewEncoder(w).Encode(
			response,
		)
	*/

	product, err := c.service.List(r.Context())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	productResponseDTO := mapper.ProductToResponse(product)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(productResponseDTO)
}
