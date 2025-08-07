package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/third_opinion/backend/internal/domain"
	"github.com/third_opinion/backend/internal/usecase"
)

// TodoHandler handles todo HTTP requests
type TodoHandler struct {
	todoUsecase *usecase.TodoUsecase
}

// NewTodoHandler creates a new TodoHandler
func NewTodoHandler(todoUsecase *usecase.TodoUsecase) *TodoHandler {
	return &TodoHandler{
		todoUsecase: todoUsecase,
	}
}

// GetAllTodos handles GET /api/todos
func (h *TodoHandler) GetAllTodos(c *gin.Context) {
	todos, err := h.todoUsecase.GetAllTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to get todos",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   todos,
	})
}

// GetTodoByID handles GET /api/todos/:id
func (h *TodoHandler) GetTodoByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid todo ID",
			"message": "ID must be a number",
		})
		return
	}

	todo, err := h.todoUsecase.GetTodoByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Todo not found",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   todo,
	})
}

// CreateTodo handles POST /api/todos
func (h *TodoHandler) CreateTodo(c *gin.Context) {
	var req domain.TodoCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request",
			"message": err.Error(),
		})
		return
	}

	todo, err := h.todoUsecase.CreateTodo(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to create todo",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"data":   todo,
	})
}

// UpdateTodo handles PUT /api/todos/:id
func (h *TodoHandler) UpdateTodo(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid todo ID",
			"message": "ID must be a number",
		})
		return
	}

	var req domain.TodoUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request",
			"message": err.Error(),
		})
		return
	}

	todo, err := h.todoUsecase.UpdateTodo(id, &req)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Todo not found",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   todo,
	})
}

// DeleteTodo handles DELETE /api/todos/:id
func (h *TodoHandler) DeleteTodo(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid todo ID",
			"message": "ID must be a number",
		})
		return
	}

	err = h.todoUsecase.DeleteTodo(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Todo not found",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Todo deleted successfully",
	})
}

// ToggleTodoComplete handles PATCH /api/todos/:id/toggle
func (h *TodoHandler) ToggleTodoComplete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid todo ID",
			"message": "ID must be a number",
		})
		return
	}

	todo, err := h.todoUsecase.ToggleTodoComplete(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Todo not found",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   todo,
	})
}
