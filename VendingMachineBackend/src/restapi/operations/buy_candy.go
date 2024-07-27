// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BuyCandyHandlerFunc turns a function with the right signature into a buy candy handler
type BuyCandyHandlerFunc func(BuyCandyParams) middleware.Responder

// Handle executing the request and returning a response
func (fn BuyCandyHandlerFunc) Handle(params BuyCandyParams) middleware.Responder {
	return fn(params)
}

// BuyCandyHandler interface for that can handle valid buy candy params
type BuyCandyHandler interface {
	Handle(BuyCandyParams) middleware.Responder
}

// NewBuyCandy creates a new http.Handler for the buy candy operation
func NewBuyCandy(ctx *middleware.Context, handler BuyCandyHandler) *BuyCandy {
	return &BuyCandy{Context: ctx, Handler: handler}
}

/*
	BuyCandy swagger:route POST /buy_candy buyCandy

BuyCandy buy candy API
*/
type BuyCandy struct {
	Context *middleware.Context
	Handler BuyCandyHandler
}

func (o *BuyCandy) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewBuyCandyParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)
}

// BuyCandyBadRequestBody buy candy bad request body
//
// swagger:model BuyCandyBadRequestBody
type BuyCandyBadRequestBody struct {
	// error
	Error string `json:"error,omitempty"`
}

// Validate validates this buy candy bad request body
func (o *BuyCandyBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this buy candy bad request body based on context it is used
func (o *BuyCandyBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *BuyCandyBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *BuyCandyBadRequestBody) UnmarshalBinary(b []byte) error {
	var res BuyCandyBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// BuyCandyBody buy candy body
//
// swagger:model BuyCandyBody
type BuyCandyBody struct {

	// number of candy
	// Required: true
	CandyCount *int64 `json:"candyCount"`

	// kind of candy
	// Required: true
	CandyType *string `json:"candyType"`

	// amount of money put into vending machine
	// Required: true
	Money *int64 `json:"money"`
}

// Validate validates this buy candy body
func (o *BuyCandyBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCandyCount(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateCandyType(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMoney(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *BuyCandyBody) validateCandyCount(formats strfmt.Registry) error {

	if err := validate.Required("order"+"."+"candyCount", "body", o.CandyCount); err != nil {
		return err
	}

	return nil
}

func (o *BuyCandyBody) validateCandyType(formats strfmt.Registry) error {

	if err := validate.Required("order"+"."+"candyType", "body", o.CandyType); err != nil {
		return err
	}

	return nil
}

func (o *BuyCandyBody) validateMoney(formats strfmt.Registry) error {

	if err := validate.Required("order"+"."+"money", "body", o.Money); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this buy candy body based on context it is used
func (o *BuyCandyBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *BuyCandyBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *BuyCandyBody) UnmarshalBinary(b []byte) error {
	var res BuyCandyBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// BuyCandyCreatedBody buy candy created body
//
// swagger:model BuyCandyCreatedBody
type BuyCandyCreatedBody struct {

	// change
	Change int64 `json:"change,omitempty"`

	// thanks
	Thanks string `json:"thanks,omitempty"`
}

// Validate validates this buy candy created body
func (o *BuyCandyCreatedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this buy candy created body based on context it is used
func (o *BuyCandyCreatedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *BuyCandyCreatedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *BuyCandyCreatedBody) UnmarshalBinary(b []byte) error {
	var res BuyCandyCreatedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// BuyCandyPaymentRequiredBody buy candy payment required body
//
// swagger:model BuyCandyPaymentRequiredBody
type BuyCandyPaymentRequiredBody struct {

	// error
	Error string `json:"error,omitempty"`
}

// Validate validates this buy candy payment required body
func (o *BuyCandyPaymentRequiredBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this buy candy payment required body based on context it is used
func (o *BuyCandyPaymentRequiredBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *BuyCandyPaymentRequiredBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *BuyCandyPaymentRequiredBody) UnmarshalBinary(b []byte) error {
	var res BuyCandyPaymentRequiredBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
