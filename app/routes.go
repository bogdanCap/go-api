package app

import (
	"github.com/bogdanCap/go-api/internal/product/controller"

	"github.com/go-chi/chi/v5"
	//"net/http"
)

func RegisterRoutes(
	r *chi.Mux,
	productController controller.ProductController,
) {

	// Health check endpoint for Kubernetes
	//r.Route("/health", func(r chi.Router) {
	//	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	//		w.WriteHeader(http.StatusOK)
	//	})
	//})
	r.Route("/products", func(r chi.Router) {

		r.Get("/", productController.List)

		/*
			r.Get("/{id}", productController.Get)

			r.Post("/", productController.Create)

			r.Put("/{id}", productController.Update)

			r.Delete("/{id}", productController.Delete)*/
	})
}
