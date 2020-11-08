// Package classification delivery API.
//
// this API Adds, Updates, Modifies and deletes any delivery for Bybrisk Maps Platform
//
// This documentation is for our internal use only
// external use of API is strictly forbidden
//
// Terms Of Service:
//
// there are no TOS at this moment and any Bybrisk member or Client is free to use API
//
//     Schemes: http, https
//     Host: api.bybrisk.com
//     BasePath: /
//     Version: 0.0.1
//     Contact: Shashank P. Sharma<imctobybrisk@gmail.com> http://csol99.blogspot.com
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
// swagger:meta

package docs

import "github.com/bybrisk/delivery-api/data"

// swagger:route POST /delivery delivery-methods POST
// delivery API uploads a new delivery to bybrisk maps Platform.
// responses:
//   200: deliveryResponse

// This is the response body.
// swagger:response deliveryResponse
type foobarResponseWrapper struct {
	// in:body
	Body data.DeliveryDetail
}

// swagger:parameters idOfdeliveryEndpoint
type foobarParamsWrapper struct {
	// This is the schema of your request body.
	// in:body
	Body data.DeliveryDetail
}