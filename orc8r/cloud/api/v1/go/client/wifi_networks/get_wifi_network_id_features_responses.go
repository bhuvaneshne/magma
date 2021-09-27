// Code generated by go-swagger; DO NOT EDIT.

package wifi_networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/api/v1/go/models"
)

// GetWifiNetworkIDFeaturesReader is a Reader for the GetWifiNetworkIDFeatures structure.
type GetWifiNetworkIDFeaturesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWifiNetworkIDFeaturesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWifiNetworkIDFeaturesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetWifiNetworkIDFeaturesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetWifiNetworkIDFeaturesOK creates a GetWifiNetworkIDFeaturesOK with default headers values
func NewGetWifiNetworkIDFeaturesOK() *GetWifiNetworkIDFeaturesOK {
	return &GetWifiNetworkIDFeaturesOK{}
}

/*GetWifiNetworkIDFeaturesOK handles this case with default header values.

Feature flags of the network
*/
type GetWifiNetworkIDFeaturesOK struct {
	Payload *models.NetworkFeatures
}

func (o *GetWifiNetworkIDFeaturesOK) Error() string {
	return fmt.Sprintf("[GET /wifi/{network_id}/features][%d] getWifiNetworkIdFeaturesOK  %+v", 200, o.Payload)
}

func (o *GetWifiNetworkIDFeaturesOK) GetPayload() *models.NetworkFeatures {
	return o.Payload
}

func (o *GetWifiNetworkIDFeaturesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NetworkFeatures)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWifiNetworkIDFeaturesDefault creates a GetWifiNetworkIDFeaturesDefault with default headers values
func NewGetWifiNetworkIDFeaturesDefault(code int) *GetWifiNetworkIDFeaturesDefault {
	return &GetWifiNetworkIDFeaturesDefault{
		_statusCode: code,
	}
}

/*GetWifiNetworkIDFeaturesDefault handles this case with default header values.

Unexpected Error
*/
type GetWifiNetworkIDFeaturesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get wifi network ID features default response
func (o *GetWifiNetworkIDFeaturesDefault) Code() int {
	return o._statusCode
}

func (o *GetWifiNetworkIDFeaturesDefault) Error() string {
	return fmt.Sprintf("[GET /wifi/{network_id}/features][%d] GetWifiNetworkIDFeatures default  %+v", o._statusCode, o.Payload)
}

func (o *GetWifiNetworkIDFeaturesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetWifiNetworkIDFeaturesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}