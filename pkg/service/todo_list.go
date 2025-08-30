package service

import (
	"github.com/MarkLover/todo-app"
	"github.com/MarkLover/todo-app/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(userId int, list todo.Todolist) (int, error) {
	return s.repo.Create(userId, list)
}
