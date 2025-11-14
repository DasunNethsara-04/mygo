package mygo

import "net/http"

// Handler function type
type HandlerFunc func(ctx *Context)

type Route struct {
	Method  string
	Path    string
	Handler HandlerFunc
}

type Router struct {
	routes []Route
}

func NewRouter() *Router {
	return &Router{
		routes: []Route{},
	}
}

// register GET route
func (r *Router) GET(path string, handler HandlerFunc) {
	r.routes = append(r.routes, Route{
		Method:  http.MethodGet,
		Path:    path,
		Handler: handler,
	})
}

// register POST route
func (r *Router) POST(path string, handler HandlerFunc) {
	r.routes = append(r.routes, Route{
		Method:  http.MethodPost,
		Path:    path,
		Handler: handler,
	})
}

// register PUT route
func (r *Router) PUT(path string, handler HandlerFunc) {
	r.routes = append(r.routes, Route{
		Method:  http.MethodPut,
		Path:    path,
		Handler: handler,
	})
}

// register DELETE route
func (r *Router) DELETE(path string, handler HandlerFunc) {
	r.routes = append(r.routes, Route{
		Method:  http.MethodDelete,
		Path:    path,
		Handler: handler,
	})
}

// this is the main entry point for all HTTP endpoint
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for _, route := range r.routes {
		if route.Method == req.Method && route.Path == req.URL.Path {
			// create context
			ctx := &Context{
				Writer:  w,
				Request: req,
				Params:  map[string]string{},
			}

			// call handler
			route.Handler(ctx)
			return
		}
	}
	http.NotFound(w, req)
}
