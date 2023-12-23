package service

import (
	"github.com/4akislav/todo-app"
	"github.com/4akislav/todo-app/pkg/repository"
)

type TodoListServise struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListServise {
	return &TodoListServise{repo: repo}
}

func (s *TodoListServise) Create(userId int, list todo.TodoList) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *TodoListServise) GetAll(userId int) ([]todo.TodoList, error) {
	return s.repo.GetAll(userId)
}

func (s *TodoListServise) GetById(userId, listId int) (todo.TodoList, error) {
	return s.repo.GetById(userId, listId)
}

func (s *TodoListServise) Delete(userId, listId int) error {
	return s.repo.Delete(userId, listId)
}

func (s *TodoListServise) Update(userId, listId int, input todo.UpdateListInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userId, listId, input)
}
