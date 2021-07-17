package repo

import (
	"github.com/google/uuid"
	"grhamm.com/todo/data"
	"grhamm.com/todo/entity"
)

var todoList []entity.Todo

func Insert(todo entity.Todo) entity.Todo {
	return data.InsertTodo(todo)
}

func Index() []entity.Todo {
	return data.FindTodo()
}

func UpdateToFinished(id uuid.UUID) {
	data.SetFinishTodo(id)
}
