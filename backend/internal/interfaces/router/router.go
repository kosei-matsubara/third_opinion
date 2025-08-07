package router

import (
	"github.com/gin-gonic/gin"
	"github.com/third_opinion/backend/internal/infrastructure/repository"
	"github.com/third_opinion/backend/internal/interfaces/handler"
	"github.com/third_opinion/backend/internal/interfaces/middleware"
	"github.com/third_opinion/backend/internal/usecase"
)

// SetupRouter configures and returns the Gin router
func SetupRouter() *gin.Engine {
	// Create router
	r := gin.Default()

	// Add CORS middleware
	r.Use(middleware.CORSMiddleware())

	// Initialize dependencies
	todoRepo := repository.NewTodoMemoryRepository()
	todoUsecase := usecase.NewTodoUsecase(todoRepo)
	todoHandler := handler.NewTodoHandler(todoUsecase)

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "OK",
			"message": "Todo API is running!",
		})
	})

	// Root endpoint
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message":     "Welcome to Todo API",
			"version":     "1.0.0",
			"description": "A simple todo management API built with Go and Gin",
		})
	})

	// API routes group
	api := r.Group("/api")
	{
		// Todo routes
		todos := api.Group("/todos")
		{
			todos.GET("", todoHandler.GetAllTodos)
			todos.POST("", todoHandler.CreateTodo)
			todos.GET("/:id", todoHandler.GetTodoByID)
			todos.PUT("/:id", todoHandler.UpdateTodo)
			todos.DELETE("/:id", todoHandler.DeleteTodo)
			todos.PATCH("/:id/toggle", todoHandler.ToggleTodoComplete)
		}
	}

	return r
}
