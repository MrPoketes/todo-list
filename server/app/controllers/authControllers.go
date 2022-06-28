package controllers

import (
	"github.com/revel/revel"
)

type AuthControllers struct {
	*revel.Controller
}

func (c AuthControllers) Index() revel.Result {
	test := make(map[string]interface{})
	test["test"] = "hello"
	return c.RenderJSON(test)
}
