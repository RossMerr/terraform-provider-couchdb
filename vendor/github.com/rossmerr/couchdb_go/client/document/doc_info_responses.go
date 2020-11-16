// Code generated by go-swagger; DO NOT EDIT.

package document

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/rossmerr/couchdb_go/models"
)

// DocInfoReader is a Reader for the DocInfo structure.
type DocInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DocInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDocInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewDocInfoNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDocInfoUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDocInfoNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDocInfoOK creates a DocInfoOK with default headers values
func NewDocInfoOK() *DocInfoOK {
	return &DocInfoOK{}
}

/*DocInfoOK handles this case with default header values.

Document exists
*/
type DocInfoOK struct {
	/*Document size
	 */
	ContentLength int64
	/*Double quoted document’s revision token
	 */
	ETag string
}

func (o *DocInfoOK) Error() string {
	return fmt.Sprintf("[HEAD /{db}/{docid}][%d] docInfoOK ", 200)
}

func (o *DocInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Content-Length
	contentLength, err := swag.ConvertInt64(response.GetHeader("Content-Length"))
	if err != nil {
		return errors.InvalidType("Content-Length", "header", "int64", response.GetHeader("Content-Length"))
	}
	o.ContentLength = contentLength

	// response header ETag
	o.ETag = response.GetHeader("ETag")

	return nil
}

// NewDocInfoNotModified creates a DocInfoNotModified with default headers values
func NewDocInfoNotModified() *DocInfoNotModified {
	return &DocInfoNotModified{}
}

/*DocInfoNotModified handles this case with default header values.

Document wasn’t modified since specified revision
*/
type DocInfoNotModified struct {
}

func (o *DocInfoNotModified) Error() string {
	return fmt.Sprintf("[HEAD /{db}/{docid}][%d] docInfoNotModified ", 304)
}

func (o *DocInfoNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDocInfoUnauthorized creates a DocInfoUnauthorized with default headers values
func NewDocInfoUnauthorized() *DocInfoUnauthorized {
	return &DocInfoUnauthorized{}
}

/*DocInfoUnauthorized handles this case with default header values.

Read privilege required
*/
type DocInfoUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *DocInfoUnauthorized) Error() string {
	return fmt.Sprintf("[HEAD /{db}/{docid}][%d] docInfoUnauthorized  %+v", 401, o.Payload)
}

func (o *DocInfoUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DocInfoUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDocInfoNotFound creates a DocInfoNotFound with default headers values
func NewDocInfoNotFound() *DocInfoNotFound {
	return &DocInfoNotFound{}
}

/*DocInfoNotFound handles this case with default header values.

Document not found
*/
type DocInfoNotFound struct {
	Payload *models.ErrorResponse
}

func (o *DocInfoNotFound) Error() string {
	return fmt.Sprintf("[HEAD /{db}/{docid}][%d] docInfoNotFound  %+v", 404, o.Payload)
}

func (o *DocInfoNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DocInfoNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}