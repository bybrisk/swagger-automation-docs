// Package classification delivery.
//
// Documentation of our delivery API.
//
//     Schemes: http
//     BasePath: /
//     Version: 1.0.0
//     Host: api.bybrisk.com
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Security:
//     - basic
//
//    SecurityDefinitions:
//    basic:
//      type: basic
//
// swagger:meta
package docs

import "github.com/bybrisk/delivery-api/data"

// swagger:route POST /delivery delivery-tag idOfdeliveryEndpoints
// delivery API uploads a new delivery to bybrisk system.
// responses:
//   200: deliveryResponse

// This text will appear as description of your response body.
// swagger:response deliveryResponse
type foobarResponseWrapper struct {
	// in:body
	Body data.DeliveryDetail
}

// swagger:parameters idOfdeliveryEndpoint
type foobarParamsWrapper struct {
	// This text will appear as description of your request body.
	// in:body
	Body data.DeliveryDetail
}