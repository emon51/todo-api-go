package controllers
import (
	beego "github.com/beego/beego/v2/server/web"
	"strconv"
) 

type TodoController struct {
	beego.Controller
}


func (c *TodoController) GetAll() {
	c.Data["json"] = map[string]string {
		"message": "All tasks",
	}
	c.ServeJSON()
}


func (c *TodoController) Post()  {
	c.Data["json"] = map[string]string {
		"message": "Todo creating...",
	}
	c.ServeJSON()
}


func (c *TodoController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Data["json"] = map[string]string {
			"error": "Invalid Id",
		}

		c.ServeJSON()
		return 
	}
	
	c.Data["json"] = map[string]interface{} {
		"message": "Getting todo", 
		"id": id,
	}

	c.ServeJSON()
}


func (c *TodoController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Data["json"] = map[string]string {
			"error": "invalid Id", 
		}
		c.ServeJSON()
		return 
	}
	c.Data["json"] = map[string]interface{} {
		"message": "Updating todo", 
		"id": id, 
	}
	c.ServeJSON()
}


func (c *TodoController) Delete(){
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Data["json"] = map[string]interface{} {
			"error": "Invalid Id", 
		}
		c.ServeJSON()
		return 
	}
	c.Data["json"] = map[string]interface{} {
		"message": "Deleting todo", 
		"id": id,
	}
	c.ServeJSON()
}


func (c *TodoController) Patch(){
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Data["json"] = map[string]interface{} {
			"error": "Invalid Id", 
		}
		c.ServeJSON()
		return 
	}
	c.Data["json"] = map[string]interface{} {
		"message": "Partial updating todo", 
		"id": id,
	}
	c.ServeJSON()
}

















