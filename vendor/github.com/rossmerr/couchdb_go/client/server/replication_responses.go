// Code generated by go-swagger; DO NOT EDIT.

package server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/rossmerr/couchdb_go/models"
)

// ReplicationReader is a Reader for the Replication structure.
type ReplicationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplicationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplicationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewReplicationAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewReplicationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewReplicationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewReplicationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewReplicationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReplicationOK creates a ReplicationOK with default headers values
func NewReplicationOK() *ReplicationOK {
	return &ReplicationOK{}
}

/*ReplicationOK handles this case with default header values.

Replication request successfully completed
*/
type ReplicationOK struct {
	Payload *models.Replication
}

func (o *ReplicationOK) Error() string {
	return fmt.Sprintf("[POST /_replicate][%d] replicationOK  %+v", 200, o.Payload)
}

func (o *ReplicationOK) GetPayload() *models.Replication {
	return o.Payload
}

func (o *ReplicationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Replication)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplicationAccepted creates a ReplicationAccepted with default headers values
func NewReplicationAccepted() *ReplicationAccepted {
	return &ReplicationAccepted{}
}

/*ReplicationAccepted handles this case with default header values.

Continuous replication request has been accepted
*/
type ReplicationAccepted struct {
	Payload *models.Replication
}

func (o *ReplicationAccepted) Error() string {
	return fmt.Sprintf("[POST /_replicate][%d] replicationAccepted  %+v", 202, o.Payload)
}

func (o *ReplicationAccepted) GetPayload() *models.Replication {
	return o.Payload
}

func (o *ReplicationAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Replication)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplicationBadRequest creates a ReplicationBadRequest with default headers values
func NewReplicationBadRequest() *ReplicationBadRequest {
	return &ReplicationBadRequest{}
}

/*ReplicationBadRequest handles this case with default header values.

Invalid JSON data
*/
type ReplicationBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *ReplicationBadRequest) Error() string {
	return fmt.Sprintf("[POST /_replicate][%d] replicationBadRequest  %+v", 400, o.Payload)
}

func (o *ReplicationBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ReplicationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplicationUnauthorized creates a ReplicationUnauthorized with default headers values
func NewReplicationUnauthorized() *ReplicationUnauthorized {
	return &ReplicationUnauthorized{}
}

/*ReplicationUnauthorized handles this case with default header values.

CouchDB Server Administrator privileges required
*/
type ReplicationUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *ReplicationUnauthorized) Error() string {
	return fmt.Sprintf("[POST /_replicate][%d] replicationUnauthorized  %+v", 401, o.Payload)
}

func (o *ReplicationUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ReplicationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplicationNotFound creates a ReplicationNotFound with default headers values
func NewReplicationNotFound() *ReplicationNotFound {
	return &ReplicationNotFound{}
}

/*ReplicationNotFound handles this case with default header values.

Either the source or target DB is not found or attempt to cancel unknown replication task
*/
type ReplicationNotFound struct {
	Payload *models.ErrorResponse
}

func (o *ReplicationNotFound) Error() string {
	return fmt.Sprintf("[POST /_replicate][%d] replicationNotFound  %+v", 404, o.Payload)
}

func (o *ReplicationNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ReplicationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplicationInternalServerError creates a ReplicationInternalServerError with default headers values
func NewReplicationInternalServerError() *ReplicationInternalServerError {
	return &ReplicationInternalServerError{}
}

/*ReplicationInternalServerError handles this case with default header values.

JSON specification was invalid
*/
type ReplicationInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *ReplicationInternalServerError) Error() string {
	return fmt.Sprintf("[POST /_replicate][%d] replicationInternalServerError  %+v", 500, o.Payload)
}

func (o *ReplicationInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ReplicationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}