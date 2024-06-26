// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PostConfigBgpPolicyApplyHandlerFunc turns a function with the right signature into a post config bgp policy apply handler
type PostConfigBgpPolicyApplyHandlerFunc func(PostConfigBgpPolicyApplyParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostConfigBgpPolicyApplyHandlerFunc) Handle(params PostConfigBgpPolicyApplyParams) middleware.Responder {
	return fn(params)
}

// PostConfigBgpPolicyApplyHandler interface for that can handle valid post config bgp policy apply params
type PostConfigBgpPolicyApplyHandler interface {
	Handle(PostConfigBgpPolicyApplyParams) middleware.Responder
}

// NewPostConfigBgpPolicyApply creates a new http.Handler for the post config bgp policy apply operation
func NewPostConfigBgpPolicyApply(ctx *middleware.Context, handler PostConfigBgpPolicyApplyHandler) *PostConfigBgpPolicyApply {
	return &PostConfigBgpPolicyApply{Context: ctx, Handler: handler}
}

/*
	PostConfigBgpPolicyApply swagger:route POST /config/bgp/policy/apply postConfigBgpPolicyApply

# Apply BGP Policy in neighbor

Apply BGP Policy in neighbor
*/
type PostConfigBgpPolicyApply struct {
	Context *middleware.Context
	Handler PostConfigBgpPolicyApplyHandler
}

func (o *PostConfigBgpPolicyApply) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostConfigBgpPolicyApplyParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
