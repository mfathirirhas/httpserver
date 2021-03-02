package httpserver

import (
	"fmt"
	"net/http"
)

type Group struct {
	server      *Server
	prefix      string
	middlewares []Middleware
}

func (s *Server) Group(prefix string, middlewares ...Middleware) *Group {
	return &Group{
		server:      s,
		prefix:      prefix,
		middlewares: middlewares,
	}
}

func (g *Group) GET(path string, handler http.HandlerFunc, middlewares ...Middleware) {
	g.server.handlers.GET(fmt.Sprintf("%s%s", g.prefix, path), f(g.server.recoverPanic(g.chainMiddlewares(handler, middlewares...))))
}

func (g *Group) HEAD(path string, handler http.HandlerFunc, middlewares ...Middleware) {
	g.server.handlers.HEAD(fmt.Sprintf("%s%s", g.prefix, path), f(g.server.recoverPanic(g.chainMiddlewares(handler, middlewares...))))
}

func (g *Group) POST(path string, handler http.HandlerFunc, middlewares ...Middleware) {
	g.server.handlers.POST(fmt.Sprintf("%s%s", g.prefix, path), f(g.server.recoverPanic(g.chainMiddlewares(handler, middlewares...))))
}

func (g *Group) PUT(path string, handler http.HandlerFunc, middlewares ...Middleware) {
	g.server.handlers.PUT(fmt.Sprintf("%s%s", g.prefix, path), f(g.server.recoverPanic(g.chainMiddlewares(handler, middlewares...))))
}

func (g *Group) DELETE(path string, handler http.HandlerFunc, middlewares ...Middleware) {
	g.server.handlers.DELETE(fmt.Sprintf("%s%s", g.prefix, path), f(g.server.recoverPanic(g.chainMiddlewares(handler, middlewares...))))
}

func (g *Group) PATCH(path string, handler http.HandlerFunc, middlewares ...Middleware) {
	g.server.handlers.PATCH(fmt.Sprintf("%s%s", g.prefix, path), f(g.server.recoverPanic(g.chainMiddlewares(handler, middlewares...))))
}

func (g *Group) OPTIONS(path string, handler http.HandlerFunc, middlewares ...Middleware) {
	g.server.handlers.OPTIONS(fmt.Sprintf("%s%s", g.prefix, path), f(g.server.recoverPanic(g.chainMiddlewares(handler, middlewares...))))
}

// FILES serve files from 1 directory dynamically in a group path.
// @filePath: must end with '/*filepath' as placeholder for filename to be accessed.
// @rootPath: root directory where @filepath locate.
func (g *Group) FILES(filePath string, rootPath string, middlewares ...Middleware) {
	g.server.FILES(fmt.Sprintf("%s%s", g.prefix, filePath), rootPath, middlewares...)
}
