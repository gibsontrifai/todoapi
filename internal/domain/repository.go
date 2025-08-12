package domain

type TodoRepository interface {
	Save(todo Todo) error
	FindByID(id string) (Todo, error)
	Update(todo Todo) error
	Delete(id string) error
	FindAll() ([]Todo, error)
}
