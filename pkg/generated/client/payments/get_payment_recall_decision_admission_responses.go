// Code generated by go-swagger; DO NOT EDIT.

package payments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/go-form3/pkg/generated/models"
)

// GetPaymentRecallDecisionAdmissionReader is a Reader for the GetPaymentRecallDecisionAdmission structure.
type GetPaymentRecallDecisionAdmissionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPaymentRecallDecisionAdmissionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPaymentRecallDecisionAdmissionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPaymentRecallDecisionAdmissionOK creates a GetPaymentRecallDecisionAdmissionOK with default headers values
func NewGetPaymentRecallDecisionAdmissionOK() *GetPaymentRecallDecisionAdmissionOK {
	return &GetPaymentRecallDecisionAdmissionOK{}
}

/*GetPaymentRecallDecisionAdmissionOK handles this case with default header values.

Recall decision admission details
*/
type GetPaymentRecallDecisionAdmissionOK struct {

	//Payload
	*models.RecallDecisionAdmissionDetailsResponse
}

func (o *GetPaymentRecallDecisionAdmissionOK) Error() string {
	return fmt.Sprintf("[GET /transaction/payments/{id}/recalls/{recallId}/decisions/{decisionId}/admissions/{admissionId}][%d] getPaymentRecallDecisionAdmissionOK  %+v", 200, o)
}

func (o *GetPaymentRecallDecisionAdmissionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.RecallDecisionAdmissionDetailsResponse = new(models.RecallDecisionAdmissionDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.RecallDecisionAdmissionDetailsResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}
