// Code generated by go-swagger; DO NOT EDIT.

package wifi_meshes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/api/v1/go/models"
)

// NewPutWifiNetworkIDMeshesMeshIDParams creates a new PutWifiNetworkIDMeshesMeshIDParams object
// with the default values initialized.
func NewPutWifiNetworkIDMeshesMeshIDParams() *PutWifiNetworkIDMeshesMeshIDParams {
	var ()
	return &PutWifiNetworkIDMeshesMeshIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutWifiNetworkIDMeshesMeshIDParamsWithTimeout creates a new PutWifiNetworkIDMeshesMeshIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutWifiNetworkIDMeshesMeshIDParamsWithTimeout(timeout time.Duration) *PutWifiNetworkIDMeshesMeshIDParams {
	var ()
	return &PutWifiNetworkIDMeshesMeshIDParams{

		timeout: timeout,
	}
}

// NewPutWifiNetworkIDMeshesMeshIDParamsWithContext creates a new PutWifiNetworkIDMeshesMeshIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutWifiNetworkIDMeshesMeshIDParamsWithContext(ctx context.Context) *PutWifiNetworkIDMeshesMeshIDParams {
	var ()
	return &PutWifiNetworkIDMeshesMeshIDParams{

		Context: ctx,
	}
}

// NewPutWifiNetworkIDMeshesMeshIDParamsWithHTTPClient creates a new PutWifiNetworkIDMeshesMeshIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutWifiNetworkIDMeshesMeshIDParamsWithHTTPClient(client *http.Client) *PutWifiNetworkIDMeshesMeshIDParams {
	var ()
	return &PutWifiNetworkIDMeshesMeshIDParams{
		HTTPClient: client,
	}
}

/*PutWifiNetworkIDMeshesMeshIDParams contains all the parameters to send to the API endpoint
for the put wifi network ID meshes mesh ID operation typically these are written to a http.Request
*/
type PutWifiNetworkIDMeshesMeshIDParams struct {

	/*MeshID
	  Mesh ID

	*/
	MeshID string
	/*NetworkID
	  Network ID

	*/
	NetworkID string
	/*WifiMesh
	  Mesh to add to the network

	*/
	WifiMesh *models.WifiMesh

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put wifi network ID meshes mesh ID params
func (o *PutWifiNetworkIDMeshesMeshIDParams) WithTimeout(timeout time.Duration) *PutWifiNetworkIDMeshesMeshIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put wifi network ID meshes mesh ID params
func (o *PutWifiNetworkIDMeshesMeshIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put wifi network ID meshes mesh ID params
func (o *PutWifiNetworkIDMeshesMeshIDParams) WithContext(ctx context.Context) *PutWifiNetworkIDMeshesMeshIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put wifi network ID meshes mesh ID params
func (o *PutWifiNetworkIDMeshesMeshIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put wifi network ID meshes mesh ID params
func (o *PutWifiNetworkIDMeshesMeshIDParams) WithHTTPClient(client *http.Client) *PutWifiNetworkIDMeshesMeshIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put wifi network ID meshes mesh ID params
func (o *PutWifiNetworkIDMeshesMeshIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMeshID adds the meshID to the put wifi network ID meshes mesh ID params
func (o *PutWifiNetworkIDMeshesMeshIDParams) WithMeshID(meshID string) *PutWifiNetworkIDMeshesMeshIDParams {
	o.SetMeshID(meshID)
	return o
}

// SetMeshID adds the meshId to the put wifi network ID meshes mesh ID params
func (o *PutWifiNetworkIDMeshesMeshIDParams) SetMeshID(meshID string) {
	o.MeshID = meshID
}

// WithNetworkID adds the networkID to the put wifi network ID meshes mesh ID params
func (o *PutWifiNetworkIDMeshesMeshIDParams) WithNetworkID(networkID string) *PutWifiNetworkIDMeshesMeshIDParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the put wifi network ID meshes mesh ID params
func (o *PutWifiNetworkIDMeshesMeshIDParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithWifiMesh adds the wifiMesh to the put wifi network ID meshes mesh ID params
func (o *PutWifiNetworkIDMeshesMeshIDParams) WithWifiMesh(wifiMesh *models.WifiMesh) *PutWifiNetworkIDMeshesMeshIDParams {
	o.SetWifiMesh(wifiMesh)
	return o
}

// SetWifiMesh adds the wifiMesh to the put wifi network ID meshes mesh ID params
func (o *PutWifiNetworkIDMeshesMeshIDParams) SetWifiMesh(wifiMesh *models.WifiMesh) {
	o.WifiMesh = wifiMesh
}

// WriteToRequest writes these params to a swagger request
func (o *PutWifiNetworkIDMeshesMeshIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param mesh_id
	if err := r.SetPathParam("mesh_id", o.MeshID); err != nil {
		return err
	}

	// path param network_id
	if err := r.SetPathParam("network_id", o.NetworkID); err != nil {
		return err
	}

	if o.WifiMesh != nil {
		if err := r.SetBodyParam(o.WifiMesh); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}