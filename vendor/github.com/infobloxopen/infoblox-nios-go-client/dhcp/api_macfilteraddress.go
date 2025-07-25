/*
Infoblox DHCP API

OpenAPI specification for Infoblox NIOS WAPI DHCP objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dhcp

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/infobloxopen/infoblox-nios-go-client/internal"
)

type MacfilteraddressAPI interface {
	/*
		Create Create a macfilteraddress object

		Creates a new macfilteraddress object

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return MacfilteraddressAPICreateRequest
	*/
	Create(ctx context.Context) MacfilteraddressAPICreateRequest

	// CreateExecute executes the request
	//  @return CreateMacfilteraddressResponse
	CreateExecute(r MacfilteraddressAPICreateRequest) (*CreateMacfilteraddressResponse, *http.Response, error)
	/*
		Delete Delete a macfilteraddress object

		Deletes a specific macfilteraddress object by reference

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param reference Reference of the macfilteraddress object
		@return MacfilteraddressAPIDeleteRequest
	*/
	Delete(ctx context.Context, reference string) MacfilteraddressAPIDeleteRequest

	// DeleteExecute executes the request
	DeleteExecute(r MacfilteraddressAPIDeleteRequest) (*http.Response, error)
	/*
		List Retrieve macfilteraddress objects

		Returns a list of macfilteraddress objects matching the search criteria

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return MacfilteraddressAPIListRequest
	*/
	List(ctx context.Context) MacfilteraddressAPIListRequest

	// ListExecute executes the request
	//  @return ListMacfilteraddressResponse
	ListExecute(r MacfilteraddressAPIListRequest) (*ListMacfilteraddressResponse, *http.Response, error)
	/*
		Read Get a specific macfilteraddress object

		Returns a specific macfilteraddress object by reference

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param reference Reference of the macfilteraddress object
		@return MacfilteraddressAPIReadRequest
	*/
	Read(ctx context.Context, reference string) MacfilteraddressAPIReadRequest

	// ReadExecute executes the request
	//  @return GetMacfilteraddressResponse
	ReadExecute(r MacfilteraddressAPIReadRequest) (*GetMacfilteraddressResponse, *http.Response, error)
	/*
		Update Update a macfilteraddress object

		Updates a specific macfilteraddress object by reference

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param reference Reference of the macfilteraddress object
		@return MacfilteraddressAPIUpdateRequest
	*/
	Update(ctx context.Context, reference string) MacfilteraddressAPIUpdateRequest

	// UpdateExecute executes the request
	//  @return UpdateMacfilteraddressResponse
	UpdateExecute(r MacfilteraddressAPIUpdateRequest) (*UpdateMacfilteraddressResponse, *http.Response, error)
}

// MacfilteraddressAPIService MacfilteraddressAPI service
type MacfilteraddressAPIService internal.Service

type MacfilteraddressAPICreateRequest struct {
	ctx              context.Context
	ApiService       MacfilteraddressAPI
	macfilteraddress *Macfilteraddress
	returnFields     *string
	returnFieldsPlus *string
	returnAsObject   *int32
}

// Object data to create
func (r MacfilteraddressAPICreateRequest) Macfilteraddress(macfilteraddress Macfilteraddress) MacfilteraddressAPICreateRequest {
	r.macfilteraddress = &macfilteraddress
	return r
}

// Enter the field names followed by comma
func (r MacfilteraddressAPICreateRequest) ReturnFields(returnFields string) MacfilteraddressAPICreateRequest {
	r.returnFields = &returnFields
	return r
}

// Enter the field names followed by comma, this returns the required fields along with the default fields
func (r MacfilteraddressAPICreateRequest) ReturnFieldsPlus(returnFieldsPlus string) MacfilteraddressAPICreateRequest {
	r.returnFieldsPlus = &returnFieldsPlus
	return r
}

// Select 1 if result is required as an object
func (r MacfilteraddressAPICreateRequest) ReturnAsObject(returnAsObject int32) MacfilteraddressAPICreateRequest {
	r.returnAsObject = &returnAsObject
	return r
}

func (r MacfilteraddressAPICreateRequest) Execute() (*CreateMacfilteraddressResponse, *http.Response, error) {
	return r.ApiService.CreateExecute(r)
}

/*
Create Create a macfilteraddress object

Creates a new macfilteraddress object

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return MacfilteraddressAPICreateRequest
*/
func (a *MacfilteraddressAPIService) Create(ctx context.Context) MacfilteraddressAPICreateRequest {
	return MacfilteraddressAPICreateRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CreateMacfilteraddressResponse
func (a *MacfilteraddressAPIService) CreateExecute(r MacfilteraddressAPICreateRequest) (*CreateMacfilteraddressResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *CreateMacfilteraddressResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "MacfilteraddressAPIService.Create")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/macfilteraddress"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.macfilteraddress == nil {
		return localVarReturnValue, nil, internal.ReportError("macfilteraddress is required and must be specified")
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
	if len(a.Client.Cfg.DefaultExtAttrs) > 0 && r.macfilteraddress != nil {
		if r.macfilteraddress.ExtAttrs == nil {
			r.macfilteraddress.ExtAttrs = &map[string]ExtAttrs{}
		}
		for k, v := range a.Client.Cfg.DefaultExtAttrs {
			if _, ok := (*r.macfilteraddress.ExtAttrs)[k]; !ok {
				(*r.macfilteraddress.ExtAttrs)[k] = ExtAttrs{
					Value: v.Value,
				}
			}
		}
	}
	// body params
	localVarPostBody = r.macfilteraddress
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

type MacfilteraddressAPIDeleteRequest struct {
	ctx        context.Context
	ApiService MacfilteraddressAPI
	reference  string
}

func (r MacfilteraddressAPIDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteExecute(r)
}

/*
Delete Delete a macfilteraddress object

Deletes a specific macfilteraddress object by reference

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param reference Reference of the macfilteraddress object
	@return MacfilteraddressAPIDeleteRequest
*/
func (a *MacfilteraddressAPIService) Delete(ctx context.Context, reference string) MacfilteraddressAPIDeleteRequest {
	return MacfilteraddressAPIDeleteRequest{
		ApiService: a,
		ctx:        ctx,
		reference:  reference,
	}
}

// Execute executes the request
func (a *MacfilteraddressAPIService) DeleteExecute(r MacfilteraddressAPIDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []internal.FormFile
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "MacfilteraddressAPIService.Delete")
	if err != nil {
		return nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/macfilteraddress/{reference}"
	localVarPath = strings.Replace(localVarPath, "{"+"reference"+"}", url.PathEscape(internal.ParameterValueToString(r.reference, "reference")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type MacfilteraddressAPIListRequest struct {
	ctx              context.Context
	ApiService       MacfilteraddressAPI
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
func (r MacfilteraddressAPIListRequest) ReturnFields(returnFields string) MacfilteraddressAPIListRequest {
	r.returnFields = &returnFields
	return r
}

// Enter the field names followed by comma, this returns the required fields along with the default fields
func (r MacfilteraddressAPIListRequest) ReturnFieldsPlus(returnFieldsPlus string) MacfilteraddressAPIListRequest {
	r.returnFieldsPlus = &returnFieldsPlus
	return r
}

// Enter the number of results to be fetched
func (r MacfilteraddressAPIListRequest) MaxResults(maxResults int32) MacfilteraddressAPIListRequest {
	r.maxResults = &maxResults
	return r
}

// Select 1 if result is required as an object
func (r MacfilteraddressAPIListRequest) ReturnAsObject(returnAsObject int32) MacfilteraddressAPIListRequest {
	r.returnAsObject = &returnAsObject
	return r
}

// Control paging of results
func (r MacfilteraddressAPIListRequest) Paging(paging int32) MacfilteraddressAPIListRequest {
	r.paging = &paging
	return r
}

// Page id for retrieving next page of results
func (r MacfilteraddressAPIListRequest) PageId(pageId string) MacfilteraddressAPIListRequest {
	r.pageId = &pageId
	return r
}

func (r MacfilteraddressAPIListRequest) Filters(filters map[string]interface{}) MacfilteraddressAPIListRequest {
	r.filters = &filters
	return r
}

func (r MacfilteraddressAPIListRequest) Extattrfilter(extattrfilter map[string]interface{}) MacfilteraddressAPIListRequest {
	r.extattrfilter = &extattrfilter
	return r
}

func (r MacfilteraddressAPIListRequest) Execute() (*ListMacfilteraddressResponse, *http.Response, error) {
	return r.ApiService.ListExecute(r)
}

/*
List Retrieve macfilteraddress objects

Returns a list of macfilteraddress objects matching the search criteria

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return MacfilteraddressAPIListRequest
*/
func (a *MacfilteraddressAPIService) List(ctx context.Context) MacfilteraddressAPIListRequest {
	return MacfilteraddressAPIListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ListMacfilteraddressResponse
func (a *MacfilteraddressAPIService) ListExecute(r MacfilteraddressAPIListRequest) (*ListMacfilteraddressResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *ListMacfilteraddressResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "MacfilteraddressAPIService.List")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/macfilteraddress"

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

type MacfilteraddressAPIReadRequest struct {
	ctx              context.Context
	ApiService       MacfilteraddressAPI
	reference        string
	returnFields     *string
	returnFieldsPlus *string
	returnAsObject   *int32
}

// Enter the field names followed by comma
func (r MacfilteraddressAPIReadRequest) ReturnFields(returnFields string) MacfilteraddressAPIReadRequest {
	r.returnFields = &returnFields
	return r
}

// Enter the field names followed by comma, this returns the required fields along with the default fields
func (r MacfilteraddressAPIReadRequest) ReturnFieldsPlus(returnFieldsPlus string) MacfilteraddressAPIReadRequest {
	r.returnFieldsPlus = &returnFieldsPlus
	return r
}

// Select 1 if result is required as an object
func (r MacfilteraddressAPIReadRequest) ReturnAsObject(returnAsObject int32) MacfilteraddressAPIReadRequest {
	r.returnAsObject = &returnAsObject
	return r
}

func (r MacfilteraddressAPIReadRequest) Execute() (*GetMacfilteraddressResponse, *http.Response, error) {
	return r.ApiService.ReadExecute(r)
}

/*
Read Get a specific macfilteraddress object

Returns a specific macfilteraddress object by reference

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param reference Reference of the macfilteraddress object
	@return MacfilteraddressAPIReadRequest
*/
func (a *MacfilteraddressAPIService) Read(ctx context.Context, reference string) MacfilteraddressAPIReadRequest {
	return MacfilteraddressAPIReadRequest{
		ApiService: a,
		ctx:        ctx,
		reference:  reference,
	}
}

// Execute executes the request
//
//	@return GetMacfilteraddressResponse
func (a *MacfilteraddressAPIService) ReadExecute(r MacfilteraddressAPIReadRequest) (*GetMacfilteraddressResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *GetMacfilteraddressResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "MacfilteraddressAPIService.Read")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/macfilteraddress/{reference}"
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

type MacfilteraddressAPIUpdateRequest struct {
	ctx              context.Context
	ApiService       MacfilteraddressAPI
	reference        string
	macfilteraddress *Macfilteraddress
	returnFields     *string
	returnFieldsPlus *string
	returnAsObject   *int32
}

// Object data to update
func (r MacfilteraddressAPIUpdateRequest) Macfilteraddress(macfilteraddress Macfilteraddress) MacfilteraddressAPIUpdateRequest {
	r.macfilteraddress = &macfilteraddress
	return r
}

// Enter the field names followed by comma
func (r MacfilteraddressAPIUpdateRequest) ReturnFields(returnFields string) MacfilteraddressAPIUpdateRequest {
	r.returnFields = &returnFields
	return r
}

// Enter the field names followed by comma, this returns the required fields along with the default fields
func (r MacfilteraddressAPIUpdateRequest) ReturnFieldsPlus(returnFieldsPlus string) MacfilteraddressAPIUpdateRequest {
	r.returnFieldsPlus = &returnFieldsPlus
	return r
}

// Select 1 if result is required as an object
func (r MacfilteraddressAPIUpdateRequest) ReturnAsObject(returnAsObject int32) MacfilteraddressAPIUpdateRequest {
	r.returnAsObject = &returnAsObject
	return r
}

func (r MacfilteraddressAPIUpdateRequest) Execute() (*UpdateMacfilteraddressResponse, *http.Response, error) {
	return r.ApiService.UpdateExecute(r)
}

/*
Update Update a macfilteraddress object

Updates a specific macfilteraddress object by reference

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param reference Reference of the macfilteraddress object
	@return MacfilteraddressAPIUpdateRequest
*/
func (a *MacfilteraddressAPIService) Update(ctx context.Context, reference string) MacfilteraddressAPIUpdateRequest {
	return MacfilteraddressAPIUpdateRequest{
		ApiService: a,
		ctx:        ctx,
		reference:  reference,
	}
}

// Execute executes the request
//
//	@return UpdateMacfilteraddressResponse
func (a *MacfilteraddressAPIService) UpdateExecute(r MacfilteraddressAPIUpdateRequest) (*UpdateMacfilteraddressResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *UpdateMacfilteraddressResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "MacfilteraddressAPIService.Update")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/macfilteraddress/{reference}"
	localVarPath = strings.Replace(localVarPath, "{"+"reference"+"}", url.PathEscape(internal.ParameterValueToString(r.reference, "reference")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.macfilteraddress == nil {
		return localVarReturnValue, nil, internal.ReportError("macfilteraddress is required and must be specified")
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
	if len(a.Client.Cfg.DefaultExtAttrs) > 0 && r.macfilteraddress != nil {
		if r.macfilteraddress.ExtAttrs == nil {
			r.macfilteraddress.ExtAttrs = &map[string]ExtAttrs{}
		}
		for k, v := range a.Client.Cfg.DefaultExtAttrs {
			if _, ok := (*r.macfilteraddress.ExtAttrs)[k]; !ok {
				(*r.macfilteraddress.ExtAttrs)[k] = ExtAttrs{
					Value: v.Value,
				}
			}
		}
	}
	// body params
	localVarPostBody = r.macfilteraddress
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
