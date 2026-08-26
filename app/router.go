package app

import (
	"net/http"

	"github.com/bogdanCap/go-api/internal/product/controller"

	"github.com/go-chi/chi/v5"
)

func NewRouter(
	productController controller.ProductController,
) http.Handler {

	r := chi.NewRouter()

	RegisterMiddleware(r)
	RegisterRoutes(r, productController)

	return r
}
