package gorest

import (
	"net/http"
	"regexp"
)

func Boot() *rest {
	return &rest{make(map[string]map[*regexp.Regexp]func(http.ResponseWriter, *http.Request, Context))}
}
