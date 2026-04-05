package routers

import (
	"todo-api/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
    beego.Router("/", &controllers.RootController{})
	beego.Router("api/v1/todos", &controllers.TodoController{}, "get:GetAll;post:Post")
	beego.Router("api/v1/todos/:id", &controllers.TodoController{}, "get:GetOne;put:Put;delete:Delete;patch:Patch")
}