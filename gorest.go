package gorest

import (
	"net/http"
)

func Boot() *rest {
	return &rest{make(map[string]map[string]func(http.ResponseWriter, *http.Request, Context))}
}
