package todoctrl

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"vincent.com/golangginrest/dao"
	"vincent.com/golangginrest/model"
)

var db = dao.DB

// Response is the common struct to cient
type Response struct {
	Status     int    `json:"status"  example:"201"`
	Message    string `json:"message"  example:"created sucessfully"`
	ResourceID uint   `json:"resourceId"  example:"1"`
}

// CreateTodo add a new todo
// @Summary create a todo item
// @Description create a todo item with title
// @ID create-to-do
// @Accept  json
// @Produce  json
// @Param  item body model.TodoModel true "new todo item"
// @Success 201 {object} todoctrl.Response
// @Router /todo [post]
func CreateTodo(c *gin.Context) {
	completed, _ := strconv.Atoi(c.PostForm("completed"))
	todo := model.TodoModel{Title: c.PostForm("title"), Completed: completed}
	db.Save(&todo)
	// c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Todo item created successfully!", "resourceId": todo.ID})
	c.JSON(http.StatusCreated, Response{Status: http.StatusCreated, Message: "Todo item created successfully!", ResourceID: todo.ID})
}

// FetchAllTodo fetch all todos
func FetchAllTodo(c *gin.Context) {
	var todos []model.TodoModel
	var _todos []model.TransformedTodo
	db.Find(&todos)
	if len(todos) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}
	//transforms the todos for building a good response
	for _, item := range todos {
		completed := false
		if item.Completed == 1 {
			completed = true
		} else {
			completed = false
		}
		_todos = append(_todos, model.TransformedTodo{ID: item.ID, Title: item.Title, Completed: completed})
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _todos})
}

// FetchSingleTodo fetch a single todo
func FetchSingleTodo(c *gin.Context) {
	var todo model.TodoModel
	todoID := c.Param("id")
	db.First(&todo, todoID)
	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}
	completed := false
	if todo.Completed == 1 {
		completed = true
	} else {
		completed = false
	}
	_todo := model.TransformedTodo{ID: todo.ID, Title: todo.Title, Completed: completed}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _todo})
}

// UpdateTodo update a todo
func UpdateTodo(c *gin.Context) {
	var todo model.TodoModel
	todoID := c.Param("id")
	db.First(&todo, todoID)
	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}
	db.Model(&todo).Update("title", c.PostForm("title"))
	completed, _ := strconv.Atoi(c.PostForm("completed"))
	db.Model(&todo).Update("completed", completed)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo updated successfully!"})
}

// DeleteTodo remove a todo
func DeleteTodo(c *gin.Context) {
	var todo model.TodoModel
	todoID := c.Param("id")
	db.First(&todo, todoID)
	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}
	db.Delete(&todo)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo deleted successfully!"})
}
