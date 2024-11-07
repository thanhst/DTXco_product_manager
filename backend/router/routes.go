package router

import (
	"net/http"
	"product_manage/controller"

	"product_manage/middleware"

	"github.com/gorilla/mux"
)

func NewRouter(userController *controller.UserController, productController *controller.ProductController, wsController *controller.WebSocketController) *mux.Router {
	router := mux.NewRouter()

	protectedRoutes := router.PathPrefix("/api").Subrouter()
	protectedRoutes.Use(middleware.JWTMiddleware)
	protectedRoutes.HandleFunc("/products/create", productController.CreateProduct).Methods("POST")
	protectedRoutes.HandleFunc("/products/update", productController.UpdateProduct).Methods("PUT")
	protectedRoutes.HandleFunc("/products/delete", productController.DeleteProduct).Methods("DELETE")
	protectedRoutes.HandleFunc("/products/get", productController.GetAllProducts).Methods("GET")
	protectedRoutes.HandleFunc("/products/get/{id}", productController.GetProductById).Methods("GET")
	// protectedRoutes.HandleFunc("/products/get", productController.GetAllProducts).Methods("GET")
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../html/main.html") // Đường dẫn đến file HTML
	}).Methods("GET")
	protectedRoutes.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	}).Methods("GET")

	router.HandleFunc("/register", userController.Register).Methods("POST")
	router.HandleFunc("/login", userController.Login).Methods("POST")
	router.HandleFunc("/ws", wsController.HandleWebSocket)

	return router
}
