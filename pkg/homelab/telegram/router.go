package telegram

import (
	"fmt"
	"strings"
)

type Router struct {
	routes map[string]*Route
}

func NewRouter() *Router {
	return &Router{
		routes: make(map[string]*Route),
	}
}

func (r *Router) Command(name string) *Route {
	name = strings.ToLower(name)
	return r.addRoute(name)
}

func (r *Router) addRoute(name string) *Route {
	route := NewRoute()
	r.routes[name] = route
	return route
}

func (r *Router) Execute(ctx Context) error {
	command := ctx.Message().Command()
	command = strings.ToLower(command)

	route, ok := r.routes[command]
	if !ok {
		return fmt.Errorf("unknown command: %s", command)
	}

	return r.executeRoute(ctx, route)
}

func (r *Router) executeRoute(ctx Context, route *Route) error {
	args := ctx.Message().CommandArguments()
	argPos := 0

	for _, token := range strings.Fields(args) {
		if len(route.args) > 0 && argPos < len(route.args) {
			name := route.args[argPos]
			ctx.args[name] = token
			argPos++
			continue
		}

		newRoute, ok := route.sub[token]
		if !ok {
			return fmt.Errorf("unknown sub-command: %s", token)
		}

		route = newRoute
		argPos = 0
	}

	if route.handler == nil {
		return nil
	}

	return route.handler(ctx)
}

type Handler func(ctx Context) error

type Route struct {
	args    []string
	sub     map[string]*Route
	handler Handler
}

func NewRoute() *Route {
	return &Route{
		args: make([]string, 0),
		sub:  make(map[string]*Route),
	}
}

func (r *Route) Command(name string) *Route {
	route := NewRoute()
	r.sub[name] = route
	return route
}

func (r *Route) Argument(name string) {
	r.args = append(r.args, name)
}

func (r *Route) Handler(handler Handler) {
	r.handler = handler
}
