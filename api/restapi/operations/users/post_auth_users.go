// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PostAuthUsersHandlerFunc turns a function with the right signature into a post auth users handler
type PostAuthUsersHandlerFunc func(PostAuthUsersParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostAuthUsersHandlerFunc) Handle(params PostAuthUsersParams) middleware.Responder {
	return fn(params)
}

// PostAuthUsersHandler interface for that can handle valid post auth users params
type PostAuthUsersHandler interface {
	Handle(PostAuthUsersParams) middleware.Responder
}

// NewPostAuthUsers creates a new http.Handler for the post auth users operation
func NewPostAuthUsers(ctx *middleware.Context, handler PostAuthUsersHandler) *PostAuthUsers {
	return &PostAuthUsers{Context: ctx, Handler: handler}
}

/*
	PostAuthUsers swagger:route POST /auth/users users postAuthUsers

# Create a new user

Creates a new user in the system
*/
type PostAuthUsers struct {
	Context *middleware.Context
	Handler PostAuthUsersHandler
}

func (o *PostAuthUsers) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostAuthUsersParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
