package repository

import (
	"fmt"
	"sync"
	"time"

	"github.com/third_opinion/backend/internal/domain"
)

// TodoMemoryRepository implements TodoRepository with in-memory storage
type TodoMemoryRepository struct {
	todos  map[int]*domain.Todo
	mutex  sync.RWMutex
	lastID int
}

// NewTodoMemoryRepository creates a new TodoMemoryRepository
func NewTodoMemoryRepository() *TodoMemoryRepository {
	return &TodoMemoryRepository{
		todos:  make(map[int]*domain.Todo),
		lastID: 0,
	}
}

// GetAll returns all todos
func (r *TodoMemoryRepository) GetAll() ([]*domain.Todo, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	todos := make([]*domain.Todo, 0, len(r.todos))
	for _, todo := range r.todos {
		todos = append(todos, todo)
	}

	return todos, nil
}

// GetByID returns a todo by ID
func (r *TodoMemoryRepository) GetByID(id int) (*domain.Todo, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	todo, exists := r.todos[id]
	if !exists {
		return nil, fmt.Errorf("todo with ID %d not found", id)
	}

	return todo, nil
}

// Create creates a new todo
func (r *TodoMemoryRepository) Create(todo *domain.Todo) (*domain.Todo, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.lastID++
	now := time.Now()

	newTodo := &domain.Todo{
		ID:          r.lastID,
		Title:       todo.Title,
		Description: todo.Description,
		Completed:   false,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	r.todos[newTodo.ID] = newTodo
	return newTodo, nil
}

// Update updates an existing todo
func (r *TodoMemoryRepository) Update(id int, updateReq *domain.TodoUpdateRequest) (*domain.Todo, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	todo, exists := r.todos[id]
	if !exists {
		return nil, fmt.Errorf("todo with ID %d not found", id)
	}

	// Update fields if provided
	if updateReq.Title != nil {
		todo.Title = *updateReq.Title
	}
	if updateReq.Description != nil {
		todo.Description = *updateReq.Description
	}
	if updateReq.Completed != nil {
		todo.Completed = *updateReq.Completed
	}

	todo.UpdatedAt = time.Now()
	r.todos[id] = todo

	return todo, nil
}

// Delete removes a todo
func (r *TodoMemoryRepository) Delete(id int) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.todos[id]; !exists {
		return fmt.Errorf("todo with ID %d not found", id)
	}

	delete(r.todos, id)
	return nil
}
