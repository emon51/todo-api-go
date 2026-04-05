package controllers

import (
	"encoding/json"
	"strconv"
	"strings"
	"todo-api/models"

	beego "github.com/beego/beego/v2/server/web"
)

// TodoController handles all todo endpoints
type TodoController struct {
	beego.Controller
}

// ErrorResponse is a helper to send error message with status code
func (c *TodoController) ErrorResponse(status_code int, error_message string) {
	c.Ctx.ResponseWriter.WriteHeader(status_code)
	c.Data["json"] = map[string]interface{}{
		"error_message": error_message,
		"status_code":   status_code,
	}
	c.ServeJSON()
}

// GetAll
func (c *TodoController) GetAll() {
	todos := models.GetAllTodos()

	if todos == nil {
		todos = []models.Todo{}
	}

	c.Data["json"] = map[string]interface{}{
		"data": todos,
	}
	c.ServeJSON()
}

// Post
func (c *TodoController) Post() {
	var body map[string]interface{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &body); err != nil {
		c.ErrorResponse(400, "Invalid JSON body.")
		return
	}
	// Validate Id
	todo_id, ok := body["id"].(float64)
	if !ok {
		c.ErrorResponse(400, "Id must be Integer.")
		return
	}

	// Validate title
	title, ok := body["title"].(string)
	if !ok || strings.TrimSpace(title) == "" {
		c.ErrorResponse(400, "Title is required and cannot be empty.")
		return
	}
	if len(title) > 255 {
		c.ErrorResponse(400, "Title cannot exceed 255 characters.")
		return
	}

	description, _ := body["description"].(string)

	todo := models.CreateTodo(int(todo_id), strings.TrimSpace(title), description)

	c.Ctx.ResponseWriter.WriteHeader(201)
	c.Data["json"] = todo
	c.ServeJSON()
}

// GetOne
func (c *TodoController) GetOne() {
	id, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if err != nil {
		c.ErrorResponse(400, "Invalid ID.")
		return
	}

	todo, index := models.GetTodoByID(id)
	if index == -1 {
		c.ErrorResponse(404, "Todo not found.")
		return
	}

	c.Data["json"] = todo
	c.ServeJSON()
}

// Put
func (c *TodoController) Put() {
	id, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if err != nil {
		c.ErrorResponse(400, "Invalid ID.")
		return
	}

	_, index := models.GetTodoByID(id)
	if index == -1 {
		c.ErrorResponse(404, "Todo not found.")
		return
	}

	var body map[string]interface{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &body); err != nil {
		c.ErrorResponse(400, "Invalid JSON body.")
		return
	}

	title, ok := body["title"].(string)
	if !ok || strings.TrimSpace(title) == "" {
		c.ErrorResponse(400, "Title is required and cannot be empty.")
		return
	}
	if len(title) > 255 {
		c.ErrorResponse(400, "Title cannot exceed 255 characters.")
		return
	}

	description, _ := body["description"].(string)
	isCompleted, _ := body["is_completed"].(bool)

	todo := models.UpdateTodo(index, strings.TrimSpace(title), description, isCompleted)
	c.Data["json"] = todo
	c.ServeJSON()
}

// Patch
func (c *TodoController) Patch() {
	id, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if err != nil {
		c.ErrorResponse(400, "Invalid ID.")
		return
	}

	_, index := models.GetTodoByID(id)
	if index == -1 {
		c.ErrorResponse(404, "Todo not found.")
		return
	}

	var body map[string]interface{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &body); err != nil {
		c.ErrorResponse(400, "Invalid JSON body.")
		return
	}

	todo := models.PatchTodo(index, body)
	c.Data["json"] = todo
	c.ServeJSON()
}

// Delete
func (c *TodoController) Delete() {
	id, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if err != nil {
		c.ErrorResponse(400, "Invalid ID.")
		return
	}

	_, index := models.GetTodoByID(id)
	if index == -1 {
		c.ErrorResponse(404, "Todo not found.")
		return
	}

	models.DeleteTodo(id)
	c.Ctx.ResponseWriter.WriteHeader(204)
}