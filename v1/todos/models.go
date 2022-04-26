package todos

import (
	"errors"
)

type Todo struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	IsCompleted bool   `json:"isCompleted"`
}

var todos = []Todo{
	{
		Id:          0,
		Title:       "todo 1",
		IsCompleted: false,
	},
	{
		Id:          1,
		Title:       "todo 2",
		IsCompleted: true,
	},
}

func findTodoById(id int) *Todo {
	for _, todo := range todos {
		if todo.Id == id {
			return &todo
		}
	}

	return nil
}

func updateTodoById(id int, updatedTodo Todo) error {
	for i, todo := range todos {
		if todo.Id == id {
			todos[i] = updatedTodo
		}

		return nil
	}

	return errors.New("todo not found")
}

func deleteTodoById(id int) error {
	if id < 0 || id >= len(todos) {
		return errors.New("todo not found")
	}

	todos = append(todos[:id], todos[id+1:]...)
	return nil
}
