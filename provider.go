package gorest

import (
	"net/http"
)

type Provider interface {
	Index(http.ResponseWriter, *http.Request, Context)
	Get(http.ResponseWriter, *http.Request, Context)
	Post(http.ResponseWriter, *http.Request, Context)
	Put(http.ResponseWriter, *http.Request, Context)
	Delete(http.ResponseWriter, *http.Request, Context)
	Head(http.ResponseWriter, *http.Request, Context)
	Patch(http.ResponseWriter, *http.Request, Context)
	Options(http.ResponseWriter, *http.Request, Context)
}
