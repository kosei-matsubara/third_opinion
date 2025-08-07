package usecase

import (
	"github.com/third_opinion/backend/internal/domain"
)

// TodoUsecase handles todo business logic
type TodoUsecase struct {
	todoRepo domain.TodoRepository
}

// NewTodoUsecase creates a new TodoUsecase
func NewTodoUsecase(todoRepo domain.TodoRepository) *TodoUsecase {
	return &TodoUsecase{
		todoRepo: todoRepo,
	}
}

// GetAllTodos returns all todos
func (u *TodoUsecase) GetAllTodos() ([]*domain.TodoResponse, error) {
	todos, err := u.todoRepo.GetAll()
	if err != nil {
		return nil, err
	}

	responses := make([]*domain.TodoResponse, len(todos))
	for i, todo := range todos {
		responses[i] = todo.ToResponse()
	}

	return responses, nil
}

// GetTodoByID returns a todo by ID
func (u *TodoUsecase) GetTodoByID(id int) (*domain.TodoResponse, error) {
	todo, err := u.todoRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return todo.ToResponse(), nil
}

// CreateTodo creates a new todo
func (u *TodoUsecase) CreateTodo(req *domain.TodoCreateRequest) (*domain.TodoResponse, error) {
	todo := &domain.Todo{
		Title:       req.Title,
		Description: req.Description,
	}

	createdTodo, err := u.todoRepo.Create(todo)
	if err != nil {
		return nil, err
	}

	return createdTodo.ToResponse(), nil
}

// UpdateTodo updates an existing todo
func (u *TodoUsecase) UpdateTodo(id int, req *domain.TodoUpdateRequest) (*domain.TodoResponse, error) {
	updatedTodo, err := u.todoRepo.Update(id, req)
	if err != nil {
		return nil, err
	}

	return updatedTodo.ToResponse(), nil
}

// DeleteTodo deletes a todo
func (u *TodoUsecase) DeleteTodo(id int) error {
	return u.todoRepo.Delete(id)
}

// ToggleTodoComplete toggles the completed status of a todo
func (u *TodoUsecase) ToggleTodoComplete(id int) (*domain.TodoResponse, error) {
	todo, err := u.todoRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	newCompleted := !todo.Completed
	updateReq := &domain.TodoUpdateRequest{
		Completed: &newCompleted,
	}

	return u.UpdateTodo(id, updateReq)
}
