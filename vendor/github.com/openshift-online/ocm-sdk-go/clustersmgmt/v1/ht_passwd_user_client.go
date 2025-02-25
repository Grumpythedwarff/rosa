/*
Copyright (c) 2020 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1

import (
	"bufio"
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/openshift-online/ocm-sdk-go/errors"
	"github.com/openshift-online/ocm-sdk-go/helpers"
)

// HTPasswdUserClient is the client of the 'HT_passwd_user' resource.
//
// Manages a specific _HTPasswd_ user.
type HTPasswdUserClient struct {
	transport http.RoundTripper
	path      string
}

// NewHTPasswdUserClient creates a new client for the 'HT_passwd_user'
// resource using the given transport to send the requests and receive the
// responses.
func NewHTPasswdUserClient(transport http.RoundTripper, path string) *HTPasswdUserClient {
	return &HTPasswdUserClient{
		transport: transport,
		path:      path,
	}
}

// Delete creates a request for the 'delete' method.
//
// Deletes the user.
func (c *HTPasswdUserClient) Delete() *HTPasswdUserDeleteRequest {
	return &HTPasswdUserDeleteRequest{
		transport: c.transport,
		path:      c.path,
	}
}

// Get creates a request for the 'get' method.
//
// Retrieves the details of the user.
func (c *HTPasswdUserClient) Get() *HTPasswdUserGetRequest {
	return &HTPasswdUserGetRequest{
		transport: c.transport,
		path:      c.path,
	}
}

// Update creates a request for the 'update' method.
//
// Updates the user's password. The username is not editable
func (c *HTPasswdUserClient) Update() *HTPasswdUserUpdateRequest {
	return &HTPasswdUserUpdateRequest{
		transport: c.transport,
		path:      c.path,
	}
}

// HTPasswdUserPollRequest is the request for the Poll method.
type HTPasswdUserPollRequest struct {
	request    *HTPasswdUserGetRequest
	interval   time.Duration
	statuses   []int
	predicates []func(interface{}) bool
}

// Parameter adds a query parameter to all the requests that will be used to retrieve the object.
func (r *HTPasswdUserPollRequest) Parameter(name string, value interface{}) *HTPasswdUserPollRequest {
	r.request.Parameter(name, value)
	return r
}

// Header adds a request header to all the requests that will be used to retrieve the object.
func (r *HTPasswdUserPollRequest) Header(name string, value interface{}) *HTPasswdUserPollRequest {
	r.request.Header(name, value)
	return r
}

// Interval sets the polling interval. This parameter is mandatory and must be greater than zero.
func (r *HTPasswdUserPollRequest) Interval(value time.Duration) *HTPasswdUserPollRequest {
	r.interval = value
	return r
}

// Status set the expected status of the response. Multiple values can be set calling this method
// multiple times. The response will be considered successful if the status is any of those values.
func (r *HTPasswdUserPollRequest) Status(value int) *HTPasswdUserPollRequest {
	r.statuses = append(r.statuses, value)
	return r
}

// Predicate adds a predicate that the response should satisfy be considered successful. Multiple
// predicates can be set calling this method multiple times. The response will be considered successful
// if all the predicates are satisfied.
func (r *HTPasswdUserPollRequest) Predicate(value func(*HTPasswdUserGetResponse) bool) *HTPasswdUserPollRequest {
	r.predicates = append(r.predicates, func(response interface{}) bool {
		return value(response.(*HTPasswdUserGetResponse))
	})
	return r
}

// StartContext starts the polling loop. Responses will be considered successful if the status is one of
// the values specified with the Status method and if all the predicates specified with the Predicate
// method return nil.
//
// The context must have a timeout or deadline, otherwise this method will immediately return an error.
func (r *HTPasswdUserPollRequest) StartContext(ctx context.Context) (response *HTPasswdUserPollResponse, err error) {
	result, err := helpers.PollContext(ctx, r.interval, r.statuses, r.predicates, r.task)
	if result != nil {
		response = &HTPasswdUserPollResponse{
			response: result.(*HTPasswdUserGetResponse),
		}
	}
	return
}

// task adapts the types of the request/response types so that they can be used with the generic
// polling function from the helpers package.
func (r *HTPasswdUserPollRequest) task(ctx context.Context) (status int, result interface{}, err error) {
	response, err := r.request.SendContext(ctx)
	if response != nil {
		status = response.Status()
		result = response
	}
	return
}

// HTPasswdUserPollResponse is the response for the Poll method.
type HTPasswdUserPollResponse struct {
	response *HTPasswdUserGetResponse
}

// Status returns the response status code.
func (r *HTPasswdUserPollResponse) Status() int {
	if r == nil {
		return 0
	}
	return r.response.Status()
}

// Header returns header of the response.
func (r *HTPasswdUserPollResponse) Header() http.Header {
	if r == nil {
		return nil
	}
	return r.response.Header()
}

// Error returns the response error.
func (r *HTPasswdUserPollResponse) Error() *errors.Error {
	if r == nil {
		return nil
	}
	return r.response.Error()
}

// Body returns the value of the 'body' parameter.
//
//
func (r *HTPasswdUserPollResponse) Body() *HTPasswdUser {
	return r.response.Body()
}

// GetBody returns the value of the 'body' parameter and
// a flag indicating if the parameter has a value.
//
//
func (r *HTPasswdUserPollResponse) GetBody() (value *HTPasswdUser, ok bool) {
	return r.response.GetBody()
}

// Poll creates a request to repeatedly retrieve the object till the response has one of a given set
// of states and satisfies a set of predicates.
func (c *HTPasswdUserClient) Poll() *HTPasswdUserPollRequest {
	return &HTPasswdUserPollRequest{
		request: c.Get(),
	}
}

// HTPasswdUserDeleteRequest is the request for the 'delete' method.
type HTPasswdUserDeleteRequest struct {
	transport http.RoundTripper
	path      string
	query     url.Values
	header    http.Header
}

// Parameter adds a query parameter.
func (r *HTPasswdUserDeleteRequest) Parameter(name string, value interface{}) *HTPasswdUserDeleteRequest {
	helpers.AddValue(&r.query, name, value)
	return r
}

// Header adds a request header.
func (r *HTPasswdUserDeleteRequest) Header(name string, value interface{}) *HTPasswdUserDeleteRequest {
	helpers.AddHeader(&r.header, name, value)
	return r
}

// Impersonate wraps requests on behalf of another user.
// Note: Services that do not support this feature may silently ignore this call.
func (r *HTPasswdUserDeleteRequest) Impersonate(user string) *HTPasswdUserDeleteRequest {
	helpers.AddImpersonationHeader(&r.header, user)
	return r
}

// Send sends this request, waits for the response, and returns it.
//
// This is a potentially lengthy operation, as it requires network communication.
// Consider using a context and the SendContext method.
func (r *HTPasswdUserDeleteRequest) Send() (result *HTPasswdUserDeleteResponse, err error) {
	return r.SendContext(context.Background())
}

// SendContext sends this request, waits for the response, and returns it.
func (r *HTPasswdUserDeleteRequest) SendContext(ctx context.Context) (result *HTPasswdUserDeleteResponse, err error) {
	query := helpers.CopyQuery(r.query)
	header := helpers.CopyHeader(r.header)
	uri := &url.URL{
		Path:     r.path,
		RawQuery: query.Encode(),
	}
	request := &http.Request{
		Method: "DELETE",
		URL:    uri,
		Header: header,
	}
	if ctx != nil {
		request = request.WithContext(ctx)
	}
	response, err := r.transport.RoundTrip(request)
	if err != nil {
		return
	}
	defer response.Body.Close()
	result = &HTPasswdUserDeleteResponse{}
	result.status = response.StatusCode
	result.header = response.Header
	reader := bufio.NewReader(response.Body)
	_, err = reader.Peek(1)
	if err == io.EOF {
		err = nil
		return
	}
	if result.status >= 400 {
		result.err, err = errors.UnmarshalErrorStatus(reader, result.status)
		if err != nil {
			return
		}
		err = result.err
		return
	}
	return
}

// HTPasswdUserDeleteResponse is the response for the 'delete' method.
type HTPasswdUserDeleteResponse struct {
	status int
	header http.Header
	err    *errors.Error
}

// Status returns the response status code.
func (r *HTPasswdUserDeleteResponse) Status() int {
	if r == nil {
		return 0
	}
	return r.status
}

// Header returns header of the response.
func (r *HTPasswdUserDeleteResponse) Header() http.Header {
	if r == nil {
		return nil
	}
	return r.header
}

// Error returns the response error.
func (r *HTPasswdUserDeleteResponse) Error() *errors.Error {
	if r == nil {
		return nil
	}
	return r.err
}

// HTPasswdUserGetRequest is the request for the 'get' method.
type HTPasswdUserGetRequest struct {
	transport http.RoundTripper
	path      string
	query     url.Values
	header    http.Header
}

// Parameter adds a query parameter.
func (r *HTPasswdUserGetRequest) Parameter(name string, value interface{}) *HTPasswdUserGetRequest {
	helpers.AddValue(&r.query, name, value)
	return r
}

// Header adds a request header.
func (r *HTPasswdUserGetRequest) Header(name string, value interface{}) *HTPasswdUserGetRequest {
	helpers.AddHeader(&r.header, name, value)
	return r
}

// Impersonate wraps requests on behalf of another user.
// Note: Services that do not support this feature may silently ignore this call.
func (r *HTPasswdUserGetRequest) Impersonate(user string) *HTPasswdUserGetRequest {
	helpers.AddImpersonationHeader(&r.header, user)
	return r
}

// Send sends this request, waits for the response, and returns it.
//
// This is a potentially lengthy operation, as it requires network communication.
// Consider using a context and the SendContext method.
func (r *HTPasswdUserGetRequest) Send() (result *HTPasswdUserGetResponse, err error) {
	return r.SendContext(context.Background())
}

// SendContext sends this request, waits for the response, and returns it.
func (r *HTPasswdUserGetRequest) SendContext(ctx context.Context) (result *HTPasswdUserGetResponse, err error) {
	query := helpers.CopyQuery(r.query)
	header := helpers.CopyHeader(r.header)
	uri := &url.URL{
		Path:     r.path,
		RawQuery: query.Encode(),
	}
	request := &http.Request{
		Method: "GET",
		URL:    uri,
		Header: header,
	}
	if ctx != nil {
		request = request.WithContext(ctx)
	}
	response, err := r.transport.RoundTrip(request)
	if err != nil {
		return
	}
	defer response.Body.Close()
	result = &HTPasswdUserGetResponse{}
	result.status = response.StatusCode
	result.header = response.Header
	reader := bufio.NewReader(response.Body)
	_, err = reader.Peek(1)
	if err == io.EOF {
		err = nil
		return
	}
	if result.status >= 400 {
		result.err, err = errors.UnmarshalErrorStatus(reader, result.status)
		if err != nil {
			return
		}
		err = result.err
		return
	}
	err = readHTPasswdUserGetResponse(result, reader)
	if err != nil {
		return
	}
	return
}

// HTPasswdUserGetResponse is the response for the 'get' method.
type HTPasswdUserGetResponse struct {
	status int
	header http.Header
	err    *errors.Error
	body   *HTPasswdUser
}

// Status returns the response status code.
func (r *HTPasswdUserGetResponse) Status() int {
	if r == nil {
		return 0
	}
	return r.status
}

// Header returns header of the response.
func (r *HTPasswdUserGetResponse) Header() http.Header {
	if r == nil {
		return nil
	}
	return r.header
}

// Error returns the response error.
func (r *HTPasswdUserGetResponse) Error() *errors.Error {
	if r == nil {
		return nil
	}
	return r.err
}

// Body returns the value of the 'body' parameter.
//
//
func (r *HTPasswdUserGetResponse) Body() *HTPasswdUser {
	if r == nil {
		return nil
	}
	return r.body
}

// GetBody returns the value of the 'body' parameter and
// a flag indicating if the parameter has a value.
//
//
func (r *HTPasswdUserGetResponse) GetBody() (value *HTPasswdUser, ok bool) {
	ok = r != nil && r.body != nil
	if ok {
		value = r.body
	}
	return
}

// HTPasswdUserUpdateRequest is the request for the 'update' method.
type HTPasswdUserUpdateRequest struct {
	transport http.RoundTripper
	path      string
	query     url.Values
	header    http.Header
	body      *HTPasswdUser
}

// Parameter adds a query parameter.
func (r *HTPasswdUserUpdateRequest) Parameter(name string, value interface{}) *HTPasswdUserUpdateRequest {
	helpers.AddValue(&r.query, name, value)
	return r
}

// Header adds a request header.
func (r *HTPasswdUserUpdateRequest) Header(name string, value interface{}) *HTPasswdUserUpdateRequest {
	helpers.AddHeader(&r.header, name, value)
	return r
}

// Impersonate wraps requests on behalf of another user.
// Note: Services that do not support this feature may silently ignore this call.
func (r *HTPasswdUserUpdateRequest) Impersonate(user string) *HTPasswdUserUpdateRequest {
	helpers.AddImpersonationHeader(&r.header, user)
	return r
}

// Body sets the value of the 'body' parameter.
//
//
func (r *HTPasswdUserUpdateRequest) Body(value *HTPasswdUser) *HTPasswdUserUpdateRequest {
	r.body = value
	return r
}

// Send sends this request, waits for the response, and returns it.
//
// This is a potentially lengthy operation, as it requires network communication.
// Consider using a context and the SendContext method.
func (r *HTPasswdUserUpdateRequest) Send() (result *HTPasswdUserUpdateResponse, err error) {
	return r.SendContext(context.Background())
}

// SendContext sends this request, waits for the response, and returns it.
func (r *HTPasswdUserUpdateRequest) SendContext(ctx context.Context) (result *HTPasswdUserUpdateResponse, err error) {
	query := helpers.CopyQuery(r.query)
	header := helpers.CopyHeader(r.header)
	buffer := &bytes.Buffer{}
	err = writeHTPasswdUserUpdateRequest(r, buffer)
	if err != nil {
		return
	}
	uri := &url.URL{
		Path:     r.path,
		RawQuery: query.Encode(),
	}
	request := &http.Request{
		Method: "PATCH",
		URL:    uri,
		Header: header,
		Body:   ioutil.NopCloser(buffer),
	}
	if ctx != nil {
		request = request.WithContext(ctx)
	}
	response, err := r.transport.RoundTrip(request)
	if err != nil {
		return
	}
	defer response.Body.Close()
	result = &HTPasswdUserUpdateResponse{}
	result.status = response.StatusCode
	result.header = response.Header
	reader := bufio.NewReader(response.Body)
	_, err = reader.Peek(1)
	if err == io.EOF {
		err = nil
		return
	}
	if result.status >= 400 {
		result.err, err = errors.UnmarshalErrorStatus(reader, result.status)
		if err != nil {
			return
		}
		err = result.err
		return
	}
	err = readHTPasswdUserUpdateResponse(result, reader)
	if err != nil {
		return
	}
	return
}

// HTPasswdUserUpdateResponse is the response for the 'update' method.
type HTPasswdUserUpdateResponse struct {
	status int
	header http.Header
	err    *errors.Error
	body   *HTPasswdUser
}

// Status returns the response status code.
func (r *HTPasswdUserUpdateResponse) Status() int {
	if r == nil {
		return 0
	}
	return r.status
}

// Header returns header of the response.
func (r *HTPasswdUserUpdateResponse) Header() http.Header {
	if r == nil {
		return nil
	}
	return r.header
}

// Error returns the response error.
func (r *HTPasswdUserUpdateResponse) Error() *errors.Error {
	if r == nil {
		return nil
	}
	return r.err
}

// Body returns the value of the 'body' parameter.
//
//
func (r *HTPasswdUserUpdateResponse) Body() *HTPasswdUser {
	if r == nil {
		return nil
	}
	return r.body
}

// GetBody returns the value of the 'body' parameter and
// a flag indicating if the parameter has a value.
//
//
func (r *HTPasswdUserUpdateResponse) GetBody() (value *HTPasswdUser, ok bool) {
	ok = r != nil && r.body != nil
	if ok {
		value = r.body
	}
	return
}
