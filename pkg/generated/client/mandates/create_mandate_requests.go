// Code generated by go-swagger; DO NOT EDIT.
// HELLO :form3: !

package mandates

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/go-form3/pkg/generated/models"
)

// Client.CreateMandate creates a new CreateMandateRequest object
// with the default values initialized.
func (c *Client) CreateMandate() *CreateMandateRequest {
	var ()
	return &CreateMandateRequest{

		MandateCreation: models.MandateCreationWithDefaults(c.Defaults),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type CreateMandateRequest struct {

	/*MandateCreationRequest*/

	*models.MandateCreation

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *CreateMandateRequest) FromJson(j string) *CreateMandateRequest {

	var m models.MandateCreation
	if err := json.Unmarshal([]byte(j), &m); err != nil {
		log.Fatal(err)
	}

	o.MandateCreation = &m

	return o
}

func (o *CreateMandateRequest) WithMandateCreationRequest(mandateCreationRequest models.MandateCreation) *CreateMandateRequest {

	o.MandateCreation = &mandateCreationRequest

	return o
}

func (o *CreateMandateRequest) WithoutMandateCreationRequest() *CreateMandateRequest {

	o.MandateCreation = &models.MandateCreation{}

	return o
}

//////////////////
// WithContext adds the context to the create mandate Request
func (o *CreateMandateRequest) WithContext(ctx context.Context) *CreateMandateRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the create mandate Request
func (o *CreateMandateRequest) WithHTTPClient(client *http.Client) *CreateMandateRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *CreateMandateRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// ISBODYPARAM
	if o.MandateCreation != nil {
		if err := r.SetBodyParam(o.MandateCreation); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
