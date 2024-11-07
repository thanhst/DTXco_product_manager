package main

import (
	"log"
	"net/http"
	"product_manage/config"
	"product_manage/controller"
	"product_manage/repository"
	"product_manage/router"
	"product_manage/service"
	"product_manage/websocket"
)

func main() {
	config.InitDB()

	userRepo := repository.NewUserRepository(config.DB)       // Khởi tạo UserRepository với kết nối DB
	productRepo := repository.NewProductRepository(config.DB) // Khởi tạo ProductRepository với kết nối DB
	userService := service.NewUserService(userRepo)
	productService := service.NewProductService(productRepo)

	wsManager := websocket.NewWebSocketManager()
	wsController := controller.NewWebSocketController(wsManager)

	userController := controller.NewUserController(userService)
	productController := controller.NewProductController(productService, wsController)

	rs := router.NewRouter(userController, productController, wsController)

	log.Println("Server started on :8386")
	log.Fatal(http.ListenAndServe(":8386", rs))
}
