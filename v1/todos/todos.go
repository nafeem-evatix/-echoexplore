package todos

import (
	"github.com/labstack/echo/v4"
)

func Initialize(group *echo.Group) {
	makeHTTPHandlers(group)
}

func makeHTTPHandlers(group *echo.Group) {
	group.GET("todos", getAllTodosHandler)
	group.POST("todos", postTodoHandler)
	group.GET("todos/:id", getSingleTodoHandler)
	group.PUT("todos/:id", updateTodoHandler)
	group.DELETE("todos/:id", deleteTodoHandler)
}
