package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mfcbentes/go-api/controller"
	"github.com/mfcbentes/go-api/db"
	"github.com/mfcbentes/go-api/repository"
	"github.com/mfcbentes/go-api/usecase"
)

func init() {
	// Carrega as variáveis de ambiente do arquivo .env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	// Camada de repository
	productRepository := repository.NewProductRepository(dbConnection)

	// Camada de usecase
	productUseCase := usecase.NewProductuseCase(*productRepository)

	// Camada de controllers
	productController := controller.NewProductController(*productUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", productController.GetProducts)
	server.GET("/product/:id", productController.GetProductById)
	server.POST("/product", productController.CreateProduct)

	server.Run(":8080")
}
