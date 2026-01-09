// cmd/api/main.go
package main

import (
	"log"
/* 
	entity "github.com/GabrielSathler/articles-backend/internal/controller/model/request"
	responseEntity "github.com/GabrielSathler/articles-backend/internal/controller/model/response" */
	"github.com/GabrielSathler/articles-backend/internal/controller/routes"
	database "github.com/GabrielSathler/articles-backend/internal/db"
	"github.com/GabrielSathler/articles-backend/internal/repository"
	"github.com/GabrielSathler/articles-backend/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Conectar ao banco de dados
	database.Connect()

	// 2. Executar migrações
	/* if err := database.DB.AutoMigrate(&entity.UserRequest{}, &responseEntity.UserResponse{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
 */
	// 3. Criar as camadas seguindo a arquitetura (injeção de dependências)
	userRepository := repository.NewUserRepository(database.DB)
	userService := service.NewUserService(userRepository)

	// 4. Configurar o router
	router := gin.Default()

	// 5. Configurar as rotas passando o service
	routes.SetupRoutes(&router.RouterGroup, userService)

	// 6. Iniciar o servidor
	log.Println("Starting server on port 8080...")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
