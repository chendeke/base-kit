package middleware

import (
	"context"
	"git.aimap.io/location/genproto/location-apis/mojo/core"
	"github.com/chendeke/config"
	"github.com/go-kit/kit/endpoint"
	"net/http"
)

type WithValidate interface {
	Validate() error
}

func Validate(entity interface{}) error {
	elem, ok := entity.(WithValidate)
	if !ok {
		return nil
	}
	return elem.Validate()
}

type Validator struct{}

func NewValidator() *Validator {
	return &Validator{}
}

type ValidatorConfig struct {
	CheckResponse bool `json:"checkResponse" yaml:"checkResponse"`
}

// Validator returns an endpoint middleware that validates request and response, so should be placed
// just before the real business endpoint
func (v *Validator) Validate() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			if err = Validate(request); err != nil {
				return nil, core.NewError(http.StatusBadRequest, err.Error())
			}
			response, err = next(ctx, request)
			if err != nil {
				return nil, err
			}
			var validatorCfg ValidatorConfig
			_ = config.ScanKey("validator", &validatorCfg)
			if !validatorCfg.CheckResponse {
				return
			}
			if err = Validate(response); err != nil {
				return nil, core.NewError(http.StatusInternalServerError, err.Error())
			}
			return
		}
	}
}
