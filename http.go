package gorest

import (
	"log"
	"net/http"
	"regexp"
	"strings"
)

type rest struct {
	routes map[string]map[string]func(http.ResponseWriter, *http.Request, Context)
}

func (r *rest) ServeHTTP(rw http.ResponseWriter, rq *http.Request) {
	log.Println(rq.RemoteAddr, rq.Method, rq.URL.RequestURI())

	routes, ok := r.routes[rq.Method]
	if ok == false {
		http.NotFound(rw, rq)
		return
	}

	c := Context{make(map[string]string)}

	if fne, ok := routes[rq.URL.Path]; ok {
		fne(rw, rq, c)
	} else {
		for p, fn := range routes {
			p = "^" + p + "$"
			re := regexp.MustCompile(p)
			if re.MatchString(rq.URL.Path) {
				ns := re.SubexpNames()[1:]
				vs := re.FindStringSubmatch(rq.URL.Path)[1:]
				if len(ns) != len(vs) {
					panic(string("URL parameter mismatch"))
				}

				for i, k := range ns {
					c.Params[k] = vs[i]
				}

				fn(rw, rq, c)
				return
			}
		}

		http.NotFound(rw, rq)
	}
}

func (r *rest) Run(s string) {
	log.Println("Starting HTTP Service on", s)
	err := http.ListenAndServe(s, r)
	if err != nil {
		log.Fatal(err)
	}
}

func (r *rest) On(m string, p string, fn func(http.ResponseWriter, *http.Request, Context)) {
	m = strings.ToUpper(m)

	if _, ok := r.routes[m]; ok == false {
		r.routes[m] = make(map[string]func(http.ResponseWriter, *http.Request, Context))
	}

	if strings.Contains(p, "/:") {
		re := regexp.MustCompile("/:(\\w+)")
		p = re.ReplaceAllString(p, "/(?P<$1>\\d+)")
	}

	r.routes[m][p] = fn
}

func (r *rest) Get(p string, fn func(http.ResponseWriter, *http.Request, Context)) {
	r.On("GET", p, fn)
}

func (r *rest) Post(p string, fn func(http.ResponseWriter, *http.Request, Context)) {
	r.On("POST", p, fn)
}

func (r *rest) Put(p string, fn func(http.ResponseWriter, *http.Request, Context)) {
	r.On("PUT", p, fn)
}

func (r *rest) Delete(p string, fn func(http.ResponseWriter, *http.Request, Context)) {
	r.On("DELETE", p, fn)
}

func (r *rest) Head(p string, fn func(http.ResponseWriter, *http.Request, Context)) {
	r.On("HEAD", p, fn)
}

func (r *rest) Patch(p string, fn func(http.ResponseWriter, *http.Request, Context)) {
	r.On("PATCH", p, fn)
}

func (r *rest) Option(p string, fn func(http.ResponseWriter, *http.Request, Context)) {
	r.On("OPTION", p, fn)
}

func (r *rest) Provide(p Provider) {

}
