// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PostLinksHandlerFunc turns a function with the right signature into a post links handler
type PostLinksHandlerFunc func(PostLinksParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostLinksHandlerFunc) Handle(params PostLinksParams) middleware.Responder {
	return fn(params)
}

// PostLinksHandler interface for that can handle valid post links params
type PostLinksHandler interface {
	Handle(PostLinksParams) middleware.Responder
}

// NewPostLinks creates a new http.Handler for the post links operation
func NewPostLinks(ctx *middleware.Context, handler PostLinksHandler) *PostLinks {
	return &PostLinks{Context: ctx, Handler: handler}
}

/* PostLinks swagger:route POST /links postLinks

creates a short link for specified link

*/
type PostLinks struct {
	Context *middleware.Context
	Handler PostLinksHandler
}

func (o *PostLinks) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostLinksParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}