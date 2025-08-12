package persistence

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sync"

	"github.com/username/todoapi/internal/domain"
	//"github.com/username/myapi/todoapi/internal/domain"
)

type FileTodoRepository struct {
	Path string
	mu   sync.RWMutex
}

func (r *FileTodoRepository) ensureFile() error {
	if _, err := os.Stat(r.Path); os.IsNotExist(err) {

		//ini untuk file kosong
		return os.WriteFile(r.Path, []byte{}, 0644)
	}
	return nil
}
func (r *FileTodoRepository) loadTodos() ([]domain.Todo, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if err := r.ensureFile(); err != nil {
		return nil, err

		// file, err := os.Open(r.Path)
		// if err != nil {
		// 	return nil, err
		// }
		// defer file.Close()

		// var todos []domain.Todo
		// if err := json.NewDecoder(file).Decode(&todos); err != nil {
		// 	return nil, err
		// }
		// return todos, nil
	}
	data, err := ioutil.ReadFile(r.Path)
	if err != nil {
		return nil, err
	}
	var todos []domain.Todo
	if err := json.Unmarshal(data, &todos); err != nil {
		return nil, err
	}
	return todos, nil
}
func (r *FileTodoRepository) saveAll(todos []domain.Todo) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if err := r.ensureFile(); err != nil {
		return err
	}
	data, err := json.Marshal(todos)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(r.Path, data, 0644)
}
