package rasweb

import "github.com/julienschmidt/httprouter"

// RouteHandler interface for abstracting out route registery and handling
type RouteHandler interface {
	Register(h *httprouter.Router) error
}
