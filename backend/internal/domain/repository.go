package domain

// TodoRepository defines methods for todo data access
type TodoRepository interface {
	GetAll() ([]*Todo, error)
	GetByID(id int) (*Todo, error)
	Create(todo *Todo) (*Todo, error)
	Update(id int, updateReq *TodoUpdateRequest) (*Todo, error)
	Delete(id int) error
}
