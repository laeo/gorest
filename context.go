package gorest

import (
	"strconv"
)

type Context struct {
	Params map[string]string
}

func (c *Context) Get(k string) string {
	if v, ok := c.Params[k]; ok {
		return string(v)
	} else {
		return string("")
	}
}

func (c *Context) GetInteger(k string) int {
	if v, ok := c.Params[k]; ok {
		i, _ := strconv.Atoi(v)
		return i
	} else {
		return 0
	}
}
