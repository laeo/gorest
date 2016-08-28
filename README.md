# What's this?

`GOREST` is a simple RESTful framework, because of write by golang, so i named it "GOREST".

## How to use?

For example:

```go
g := gorest.Boot()

g.Get("/users/:id", func(w http.ResponseWriter, r *http.Request, c gorest.Context){
    w.Write([]byte("User ID: " + c.Get("id")))
})

g.Run(":3000")
```

## License

Based on `MIT` license.
