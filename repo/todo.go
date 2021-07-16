package repo

import (
	"github.com/google/uuid"
	"grhamm.com/todo/entity"
)

var todoList []entity.Todo

func Insert(todo entity.Todo) entity.Todo {
	todo.Id = uuid.New()
	todoList = append(todoList, todo)

	return todo
}

func Index() []entity.Todo {
	return todoList
}

func UpdateToFinished(id uuid.UUID) entity.Todo {
	for i, todo := range todoList {
		if id == todo.Id {
			todoList[i].Finished = true
			return todoList[i]
		}
	}

	return entity.Todo{}
}
