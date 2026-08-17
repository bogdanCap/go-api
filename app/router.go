package app

import (
	"net/http"

	"test/internal/controller"

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
