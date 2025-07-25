/*
Infoblox DNS API

OpenAPI specification for Infoblox NIOS WAPI DNS objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dns

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/infobloxopen/infoblox-nios-go-client/internal"
)

type RecordAaaaAPI interface {
	/*
		Create Create a record:aaaa object

		Creates a new record:aaaa object

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return RecordAaaaAPICreateRequest
	*/
	Create(ctx context.Context) RecordAaaaAPICreateRequest

	// CreateExecute executes the request
	//  @return CreateRecordAaaaResponse
	CreateExecute(r RecordAaaaAPICreateRequest) (*CreateRecordAaaaResponse, *http.Response, error)
	/*
		Delete Delete a record:aaaa object

		Deletes a specific record:aaaa object by reference

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param reference Reference of the record:aaaa object
		@return RecordAaaaAPIDeleteRequest
	*/
	Delete(ctx context.Context, reference string) RecordAaaaAPIDeleteRequest

	// DeleteExecute executes the request
	DeleteExecute(r RecordAaaaAPIDeleteRequest) (*http.Response, error)
	/*
		List Retrieve record:aaaa objects

		Returns a list of record:aaaa objects matching the search criteria

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return RecordAaaaAPIListRequest
	*/
	List(ctx context.Context) RecordAaaaAPIListRequest

	// ListExecute executes the request
	//  @return ListRecordAaaaResponse
	ListExecute(r RecordAaaaAPIListRequest) (*ListRecordAaaaResponse, *http.Response, error)
	/*
		Read Get a specific record:aaaa object

		Returns a specific record:aaaa object by reference

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param reference Reference of the record:aaaa object
		@return RecordAaaaAPIReadRequest
	*/
	Read(ctx context.Context, reference string) RecordAaaaAPIReadRequest

	// ReadExecute executes the request
	//  @return GetRecordAaaaResponse
	ReadExecute(r RecordAaaaAPIReadRequest) (*GetRecordAaaaResponse, *http.Response, error)
	/*
		Update Update a record:aaaa object

		Updates a specific record:aaaa object by reference

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param reference Reference of the record:aaaa object
		@return RecordAaaaAPIUpdateRequest
	*/
	Update(ctx context.Context, reference string) RecordAaaaAPIUpdateRequest

	// UpdateExecute executes the request
	//  @return UpdateRecordAaaaResponse
	UpdateExecute(r RecordAaaaAPIUpdateRequest) (*UpdateRecordAaaaResponse, *http.Response, error)
}

// RecordAaaaAPIService RecordAaaaAPI service
type RecordAaaaAPIService internal.Service

type RecordAaaaAPICreateRequest struct {
	ctx              context.Context
	ApiService       RecordAaaaAPI
	recordAaaa       *RecordAaaa
	returnFields     *string
	returnFieldsPlus *string
	returnAsObject   *int32
}

// Object data to create
func (r RecordAaaaAPICreateRequest) RecordAaaa(recordAaaa RecordAaaa) RecordAaaaAPICreateRequest {
	r.recordAaaa = &recordAaaa
	return r
}

// Enter the field names followed by comma
func (r RecordAaaaAPICreateRequest) ReturnFields(returnFields string) RecordAaaaAPICreateRequest {
	r.returnFields = &returnFields
	return r
}

// Enter the field names followed by comma, this returns the required fields along with the default fields
func (r RecordAaaaAPICreateRequest) ReturnFieldsPlus(returnFieldsPlus string) RecordAaaaAPICreateRequest {
	r.returnFieldsPlus = &returnFieldsPlus
	return r
}

// Select 1 if result is required as an object
func (r RecordAaaaAPICreateRequest) ReturnAsObject(returnAsObject int32) RecordAaaaAPICreateRequest {
	r.returnAsObject = &returnAsObject
	return r
}

func (r RecordAaaaAPICreateRequest) Execute() (*CreateRecordAaaaResponse, *http.Response, error) {
	return r.ApiService.CreateExecute(r)
}

/*
Create Create a record:aaaa object

Creates a new record:aaaa object

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return RecordAaaaAPICreateRequest
*/
func (a *RecordAaaaAPIService) Create(ctx context.Context) RecordAaaaAPICreateRequest {
	return RecordAaaaAPICreateRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CreateRecordAaaaResponse
func (a *RecordAaaaAPIService) CreateExecute(r RecordAaaaAPICreateRequest) (*CreateRecordAaaaResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *CreateRecordAaaaResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "RecordAaaaAPIService.Create")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/record:aaaa"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.recordAaaa == nil {
		return localVarReturnValue, nil, internal.ReportError("recordAaaa is required and must be specified")
	}

	if r.returnFields != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_return_fields", r.returnFields, "form", "")
	}
	if r.returnFieldsPlus != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_return_fields+", r.returnFieldsPlus, "form", "")
	}
	if r.returnAsObject != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_return_as_object", r.returnAsObject, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := internal.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := internal.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if len(a.Client.Cfg.DefaultExtAttrs) > 0 && r.recordAaaa != nil {
		if r.recordAaaa.ExtAttrs == nil {
			r.recordAaaa.ExtAttrs = &map[string]ExtAttrs{}
		}
		for k, v := range a.Client.Cfg.DefaultExtAttrs {
			if _, ok := (*r.recordAaaa.ExtAttrs)[k]; !ok {
				(*r.recordAaaa.ExtAttrs)[k] = ExtAttrs{
					Value: v.Value,
				}
			}
		}
	}
	if r.recordAaaa.FuncCall != nil {
		bodyForFuncCall := r.recordAaaa
		if bodyForFuncCall.FuncCall.AttributeName == "" {
			return localVarReturnValue, nil, internal.ReportError("FuncCall.AttributeName is required and must be specified")
		}
		var funcStr string = bodyForFuncCall.FuncCall.AttributeName
		if funcStr == "Ipv6addr" {
			if bodyForFuncCall.Ipv6addr.String != nil {
				return localVarReturnValue, nil, internal.ReportError("Ipv6addr cannot be provided when function call is used")
			} else {

				var l RecordAaaaIpv6addr
				var m RecordAaaaIpv6addrOneOf
				m.ObjectFunction = bodyForFuncCall.FuncCall.ObjectFunction
				m.Parameters = bodyForFuncCall.FuncCall.Parameters
				m.ResultField = bodyForFuncCall.FuncCall.ResultField
				m.Object = bodyForFuncCall.FuncCall.Object
				m.ObjectParameters = bodyForFuncCall.FuncCall.ObjectParameters

				l.RecordAaaaIpv6addrOneOf = &m
				l.String = nil
				bodyForFuncCall.Ipv6addr = &l
				bodyForFuncCall.FuncCall = nil
			}
		}
		r.recordAaaa = bodyForFuncCall
	}
	// body params
	localVarPostBody = r.recordAaaa
	req, err := a.Client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := internal.NewGenericOpenAPIErrorWithBody(localVarHTTPResponse.Status, localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := internal.NewGenericOpenAPIErrorWithBody(err.Error(), localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}
	return localVarReturnValue, localVarHTTPResponse, nil
}

type RecordAaaaAPIDeleteRequest struct {
	ctx                 context.Context
	ApiService          RecordAaaaAPI
	reference           string
	removeAssociatedPtr *bool
}

// Delete option that indicates whether the associated PTR records should be removed while deleting the specified A record.
func (r RecordAaaaAPIDeleteRequest) RemoveAssociatedPtr(removeAssociatedPtr bool) RecordAaaaAPIDeleteRequest {
	r.removeAssociatedPtr = &removeAssociatedPtr
	return r
}

func (r RecordAaaaAPIDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteExecute(r)
}

/*
Delete Delete a record:aaaa object

Deletes a specific record:aaaa object by reference

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param reference Reference of the record:aaaa object
	@return RecordAaaaAPIDeleteRequest
*/
func (a *RecordAaaaAPIService) Delete(ctx context.Context, reference string) RecordAaaaAPIDeleteRequest {
	return RecordAaaaAPIDeleteRequest{
		ApiService: a,
		ctx:        ctx,
		reference:  reference,
	}
}

// Execute executes the request
func (a *RecordAaaaAPIService) DeleteExecute(r RecordAaaaAPIDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []internal.FormFile
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "RecordAaaaAPIService.Delete")
	if err != nil {
		return nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/record:aaaa/{reference}"
	localVarPath = strings.Replace(localVarPath, "{"+"reference"+"}", url.PathEscape(internal.ParameterValueToString(r.reference, "reference")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.removeAssociatedPtr != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "remove_associated_ptr", r.removeAssociatedPtr, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := internal.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := internal.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.Client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := internal.NewGenericOpenAPIErrorWithBody(localVarHTTPResponse.Status, localVarBody)
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type RecordAaaaAPIListRequest struct {
	ctx              context.Context
	ApiService       RecordAaaaAPI
	returnFields     *string
	returnFieldsPlus *string
	maxResults       *int32
	returnAsObject   *int32
	paging           *int32
	pageId           *string
	filters          *map[string]interface{}
	extattrfilter    *map[string]interface{}
}

// Enter the field names followed by comma
func (r RecordAaaaAPIListRequest) ReturnFields(returnFields string) RecordAaaaAPIListRequest {
	r.returnFields = &returnFields
	return r
}

// Enter the field names followed by comma, this returns the required fields along with the default fields
func (r RecordAaaaAPIListRequest) ReturnFieldsPlus(returnFieldsPlus string) RecordAaaaAPIListRequest {
	r.returnFieldsPlus = &returnFieldsPlus
	return r
}

// Enter the number of results to be fetched
func (r RecordAaaaAPIListRequest) MaxResults(maxResults int32) RecordAaaaAPIListRequest {
	r.maxResults = &maxResults
	return r
}

// Select 1 if result is required as an object
func (r RecordAaaaAPIListRequest) ReturnAsObject(returnAsObject int32) RecordAaaaAPIListRequest {
	r.returnAsObject = &returnAsObject
	return r
}

// Control paging of results
func (r RecordAaaaAPIListRequest) Paging(paging int32) RecordAaaaAPIListRequest {
	r.paging = &paging
	return r
}

// Page id for retrieving next page of results
func (r RecordAaaaAPIListRequest) PageId(pageId string) RecordAaaaAPIListRequest {
	r.pageId = &pageId
	return r
}

func (r RecordAaaaAPIListRequest) Filters(filters map[string]interface{}) RecordAaaaAPIListRequest {
	r.filters = &filters
	return r
}

func (r RecordAaaaAPIListRequest) Extattrfilter(extattrfilter map[string]interface{}) RecordAaaaAPIListRequest {
	r.extattrfilter = &extattrfilter
	return r
}

func (r RecordAaaaAPIListRequest) Execute() (*ListRecordAaaaResponse, *http.Response, error) {
	return r.ApiService.ListExecute(r)
}

/*
List Retrieve record:aaaa objects

Returns a list of record:aaaa objects matching the search criteria

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return RecordAaaaAPIListRequest
*/
func (a *RecordAaaaAPIService) List(ctx context.Context) RecordAaaaAPIListRequest {
	return RecordAaaaAPIListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ListRecordAaaaResponse
func (a *RecordAaaaAPIService) ListExecute(r RecordAaaaAPIListRequest) (*ListRecordAaaaResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *ListRecordAaaaResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "RecordAaaaAPIService.List")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/record:aaaa"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.returnFields != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_return_fields", r.returnFields, "form", "")
	}
	if r.returnFieldsPlus != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_return_fields+", r.returnFieldsPlus, "form", "")
	}
	if r.maxResults != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_max_results", r.maxResults, "form", "")
	}
	if r.returnAsObject != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_return_as_object", r.returnAsObject, "form", "")
	}
	if r.paging != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_paging", r.paging, "form", "")
	}
	if r.pageId != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_page_id", r.pageId, "form", "")
	}
	if r.filters != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "filters", r.filters, "form", "")
	}
	if r.extattrfilter != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "extattrfilter", r.extattrfilter, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := internal.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := internal.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.Client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := internal.NewGenericOpenAPIErrorWithBody(localVarHTTPResponse.Status, localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := internal.NewGenericOpenAPIErrorWithBody(err.Error(), localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}
	return localVarReturnValue, localVarHTTPResponse, nil
}

type RecordAaaaAPIReadRequest struct {
	ctx              context.Context
	ApiService       RecordAaaaAPI
	reference        string
	returnFields     *string
	returnFieldsPlus *string
	returnAsObject   *int32
}

// Enter the field names followed by comma
func (r RecordAaaaAPIReadRequest) ReturnFields(returnFields string) RecordAaaaAPIReadRequest {
	r.returnFields = &returnFields
	return r
}

// Enter the field names followed by comma, this returns the required fields along with the default fields
func (r RecordAaaaAPIReadRequest) ReturnFieldsPlus(returnFieldsPlus string) RecordAaaaAPIReadRequest {
	r.returnFieldsPlus = &returnFieldsPlus
	return r
}

// Select 1 if result is required as an object
func (r RecordAaaaAPIReadRequest) ReturnAsObject(returnAsObject int32) RecordAaaaAPIReadRequest {
	r.returnAsObject = &returnAsObject
	return r
}

func (r RecordAaaaAPIReadRequest) Execute() (*GetRecordAaaaResponse, *http.Response, error) {
	return r.ApiService.ReadExecute(r)
}

/*
Read Get a specific record:aaaa object

Returns a specific record:aaaa object by reference

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param reference Reference of the record:aaaa object
	@return RecordAaaaAPIReadRequest
*/
func (a *RecordAaaaAPIService) Read(ctx context.Context, reference string) RecordAaaaAPIReadRequest {
	return RecordAaaaAPIReadRequest{
		ApiService: a,
		ctx:        ctx,
		reference:  reference,
	}
}

// Execute executes the request
//
//	@return GetRecordAaaaResponse
func (a *RecordAaaaAPIService) ReadExecute(r RecordAaaaAPIReadRequest) (*GetRecordAaaaResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *GetRecordAaaaResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "RecordAaaaAPIService.Read")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/record:aaaa/{reference}"
	localVarPath = strings.Replace(localVarPath, "{"+"reference"+"}", url.PathEscape(internal.ParameterValueToString(r.reference, "reference")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.returnFields != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_return_fields", r.returnFields, "form", "")
	}
	if r.returnFieldsPlus != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_return_fields+", r.returnFieldsPlus, "form", "")
	}
	if r.returnAsObject != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_return_as_object", r.returnAsObject, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := internal.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := internal.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.Client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := internal.NewGenericOpenAPIErrorWithBody(localVarHTTPResponse.Status, localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := internal.NewGenericOpenAPIErrorWithBody(err.Error(), localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}
	return localVarReturnValue, localVarHTTPResponse, nil
}

type RecordAaaaAPIUpdateRequest struct {
	ctx              context.Context
	ApiService       RecordAaaaAPI
	reference        string
	recordAaaa       *RecordAaaa
	returnFields     *string
	returnFieldsPlus *string
	returnAsObject   *int32
}

// Object data to update
func (r RecordAaaaAPIUpdateRequest) RecordAaaa(recordAaaa RecordAaaa) RecordAaaaAPIUpdateRequest {
	r.recordAaaa = &recordAaaa
	return r
}

// Enter the field names followed by comma
func (r RecordAaaaAPIUpdateRequest) ReturnFields(returnFields string) RecordAaaaAPIUpdateRequest {
	r.returnFields = &returnFields
	return r
}

// Enter the field names followed by comma, this returns the required fields along with the default fields
func (r RecordAaaaAPIUpdateRequest) ReturnFieldsPlus(returnFieldsPlus string) RecordAaaaAPIUpdateRequest {
	r.returnFieldsPlus = &returnFieldsPlus
	return r
}

// Select 1 if result is required as an object
func (r RecordAaaaAPIUpdateRequest) ReturnAsObject(returnAsObject int32) RecordAaaaAPIUpdateRequest {
	r.returnAsObject = &returnAsObject
	return r
}

func (r RecordAaaaAPIUpdateRequest) Execute() (*UpdateRecordAaaaResponse, *http.Response, error) {
	return r.ApiService.UpdateExecute(r)
}

/*
Update Update a record:aaaa object

Updates a specific record:aaaa object by reference

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param reference Reference of the record:aaaa object
	@return RecordAaaaAPIUpdateRequest
*/
func (a *RecordAaaaAPIService) Update(ctx context.Context, reference string) RecordAaaaAPIUpdateRequest {
	return RecordAaaaAPIUpdateRequest{
		ApiService: a,
		ctx:        ctx,
		reference:  reference,
	}
}

// Execute executes the request
//
//	@return UpdateRecordAaaaResponse
func (a *RecordAaaaAPIService) UpdateExecute(r RecordAaaaAPIUpdateRequest) (*UpdateRecordAaaaResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *UpdateRecordAaaaResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "RecordAaaaAPIService.Update")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/record:aaaa/{reference}"
	localVarPath = strings.Replace(localVarPath, "{"+"reference"+"}", url.PathEscape(internal.ParameterValueToString(r.reference, "reference")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.recordAaaa == nil {
		return localVarReturnValue, nil, internal.ReportError("recordAaaa is required and must be specified")
	}

	if r.returnFields != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_return_fields", r.returnFields, "form", "")
	}
	if r.returnFieldsPlus != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_return_fields+", r.returnFieldsPlus, "form", "")
	}
	if r.returnAsObject != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_return_as_object", r.returnAsObject, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := internal.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := internal.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if len(a.Client.Cfg.DefaultExtAttrs) > 0 && r.recordAaaa != nil {
		if r.recordAaaa.ExtAttrs == nil {
			r.recordAaaa.ExtAttrs = &map[string]ExtAttrs{}
		}
		for k, v := range a.Client.Cfg.DefaultExtAttrs {
			if _, ok := (*r.recordAaaa.ExtAttrs)[k]; !ok {
				(*r.recordAaaa.ExtAttrs)[k] = ExtAttrs{
					Value: v.Value,
				}
			}
		}
	}
	if r.recordAaaa.FuncCall != nil {
		bodyForFuncCall := r.recordAaaa
		bodyForFuncCall.FuncCall = nil
		bodyForFuncCall.Ipv6addr = nil
		r.recordAaaa = bodyForFuncCall
	}
	// body params
	localVarPostBody = r.recordAaaa
	req, err := a.Client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := internal.NewGenericOpenAPIErrorWithBody(localVarHTTPResponse.Status, localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := internal.NewGenericOpenAPIErrorWithBody(err.Error(), localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}
	return localVarReturnValue, localVarHTTPResponse, nil
}
