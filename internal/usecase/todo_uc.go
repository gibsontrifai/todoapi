package usecase

import (
	"errors"
	"time"

	"github.com/username/todoapi/internal/domain"
	//"github.com/username/myapi/todoapi/internal/domain"
)

type TodoUsecase struct {
	Repo domain.TodoRepository
}

func NewTodoUsecase(repo domain.TodoRepository) *TodoUsecase {
	return &TodoUsecase{Repo: repo}
}
func (uc *TodoUsecase) CreateTodo(title, description string) (domain.Todo, error) {
	if title == "" {
		return domain.Todo{}, errors.New("title is required")
	}
	t := domain.NewTodo("", title, description)
	return t, uc.Repo.Save(t)
}
func (uc *TodoUsecase) ListTodos() ([]domain.Todo, error) {
	return uc.Repo.FindAll()
}
func (uc *TodoUsecase) GetTodoByID(id string) (domain.Todo, error) {
	return uc.Repo.FindByID(id)
}
func (uc *TodoUsecase) UpdateTodo(id, title, description string) (domain.Todo, error) {
	t, err := uc.Repo.FindByID(id)
	if err != nil {
		return domain.Todo{}, err
	}
	if title != "" {
		//t.Title = todo.Title
		t.Title = title
	}
	if description != "" {
		t.Description = description
	}
	t.UpdatedAt = time.Now()
	if err := uc.Repo.Update(t); err != nil {
		return domain.Todo{}, err
	}
	return t, nil
}
func (uc *TodoUsecase) DeleteTodo(id string) error {
	return uc.Repo.Delete(id)
}
