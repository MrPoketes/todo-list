package controllers

import (
	"github.com/revel/revel"
)

type User struct {
	*revel.Controller
}

func (c User) GetItems() revel.Result {
	return c.RenderJSON("")
}

func (c User) PostItems() revel.Result {
	return c.RenderJSON("")
}
