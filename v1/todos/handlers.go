package todos

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func getSingleTodoHandler(ctx echo.Context) error {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			echo.HTTPError{Message: "string to int conversion failed"})
		return nil
	}

	todo := findTodoById(id)
	if todo == nil {
		ctx.JSON(
			http.StatusNotFound,
			echo.HTTPError{Message: "todo not found"})
		return nil
	}

	return ctx.JSON(http.StatusOK, todo)
}

func postTodoHandler(ctx echo.Context) error {
	var req Todo
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusInternalServerError, echo.HTTPError{Message: "binding failed"})
		return nil
	}

	todos = append(todos, req)

	return ctx.JSON(http.StatusOK, &req)
}

func getAllTodosHandler(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, &todos)
}

func updateTodoHandler(ctx echo.Context) error {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			echo.HTTPError{Message: "string to int conversion failed"})
		return nil
	}

	var req Todo
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusInternalServerError, echo.HTTPError{Message: "binding failed"})
		return nil
	}

	updateErr := updateTodoById(id, req)
	if updateErr != nil {
		ctx.JSON(http.StatusInternalServerError, echo.HTTPError{Message: updateErr.Error()})
		return nil
	}

	return ctx.JSON(http.StatusOK, &req)
}

func deleteTodoHandler(ctx echo.Context) error {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			echo.HTTPError{Message: "string to int conversion failed"})
		return nil
	}

	deleteErr := deleteTodoById(id)
	if deleteErr != nil {
		ctx.JSON(http.StatusInternalServerError, echo.HTTPError{Message: deleteErr.Error()})
		return nil
	}

	return ctx.JSON(http.StatusOK, &todos)
}
