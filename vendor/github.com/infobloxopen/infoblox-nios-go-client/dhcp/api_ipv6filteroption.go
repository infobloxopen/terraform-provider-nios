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

type Ipv6filteroptionAPI interface {
	/*
		Create Create a ipv6filteroption object

		Creates a new ipv6filteroption object

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return Ipv6filteroptionAPICreateRequest
	*/
	Create(ctx context.Context) Ipv6filteroptionAPICreateRequest

	// CreateExecute executes the request
	//  @return CreateIpv6filteroptionResponse
	CreateExecute(r Ipv6filteroptionAPICreateRequest) (*CreateIpv6filteroptionResponse, *http.Response, error)
	/*
		Delete Delete a ipv6filteroption object

		Deletes a specific ipv6filteroption object by reference

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param reference Reference of the ipv6filteroption object
		@return Ipv6filteroptionAPIDeleteRequest
	*/
	Delete(ctx context.Context, reference string) Ipv6filteroptionAPIDeleteRequest

	// DeleteExecute executes the request
	DeleteExecute(r Ipv6filteroptionAPIDeleteRequest) (*http.Response, error)
	/*
		List Retrieve ipv6filteroption objects

		Returns a list of ipv6filteroption objects matching the search criteria

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return Ipv6filteroptionAPIListRequest
	*/
	List(ctx context.Context) Ipv6filteroptionAPIListRequest

	// ListExecute executes the request
	//  @return ListIpv6filteroptionResponse
	ListExecute(r Ipv6filteroptionAPIListRequest) (*ListIpv6filteroptionResponse, *http.Response, error)
	/*
		Read Get a specific ipv6filteroption object

		Returns a specific ipv6filteroption object by reference

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param reference Reference of the ipv6filteroption object
		@return Ipv6filteroptionAPIReadRequest
	*/
	Read(ctx context.Context, reference string) Ipv6filteroptionAPIReadRequest

	// ReadExecute executes the request
	//  @return GetIpv6filteroptionResponse
	ReadExecute(r Ipv6filteroptionAPIReadRequest) (*GetIpv6filteroptionResponse, *http.Response, error)
	/*
		Update Update a ipv6filteroption object

		Updates a specific ipv6filteroption object by reference

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param reference Reference of the ipv6filteroption object
		@return Ipv6filteroptionAPIUpdateRequest
	*/
	Update(ctx context.Context, reference string) Ipv6filteroptionAPIUpdateRequest

	// UpdateExecute executes the request
	//  @return UpdateIpv6filteroptionResponse
	UpdateExecute(r Ipv6filteroptionAPIUpdateRequest) (*UpdateIpv6filteroptionResponse, *http.Response, error)
}

// Ipv6filteroptionAPIService Ipv6filteroptionAPI service
type Ipv6filteroptionAPIService internal.Service

type Ipv6filteroptionAPICreateRequest struct {
	ctx              context.Context
	ApiService       Ipv6filteroptionAPI
	ipv6filteroption *Ipv6filteroption
	returnFields     *string
	returnFieldsPlus *string
	returnAsObject   *int32
}

// Object data to create
func (r Ipv6filteroptionAPICreateRequest) Ipv6filteroption(ipv6filteroption Ipv6filteroption) Ipv6filteroptionAPICreateRequest {
	r.ipv6filteroption = &ipv6filteroption
	return r
}

// Enter the field names followed by comma
func (r Ipv6filteroptionAPICreateRequest) ReturnFields(returnFields string) Ipv6filteroptionAPICreateRequest {
	r.returnFields = &returnFields
	return r
}

// Enter the field names followed by comma, this returns the required fields along with the default fields
func (r Ipv6filteroptionAPICreateRequest) ReturnFieldsPlus(returnFieldsPlus string) Ipv6filteroptionAPICreateRequest {
	r.returnFieldsPlus = &returnFieldsPlus
	return r
}

// Select 1 if result is required as an object
func (r Ipv6filteroptionAPICreateRequest) ReturnAsObject(returnAsObject int32) Ipv6filteroptionAPICreateRequest {
	r.returnAsObject = &returnAsObject
	return r
}

func (r Ipv6filteroptionAPICreateRequest) Execute() (*CreateIpv6filteroptionResponse, *http.Response, error) {
	return r.ApiService.CreateExecute(r)
}

/*
Create Create a ipv6filteroption object

Creates a new ipv6filteroption object

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return Ipv6filteroptionAPICreateRequest
*/
func (a *Ipv6filteroptionAPIService) Create(ctx context.Context) Ipv6filteroptionAPICreateRequest {
	return Ipv6filteroptionAPICreateRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CreateIpv6filteroptionResponse
func (a *Ipv6filteroptionAPIService) CreateExecute(r Ipv6filteroptionAPICreateRequest) (*CreateIpv6filteroptionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *CreateIpv6filteroptionResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "Ipv6filteroptionAPIService.Create")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/ipv6filteroption"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.ipv6filteroption == nil {
		return localVarReturnValue, nil, internal.ReportError("ipv6filteroption is required and must be specified")
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
	if len(a.Client.Cfg.DefaultExtAttrs) > 0 && r.ipv6filteroption != nil {
		if r.ipv6filteroption.ExtAttrs == nil {
			r.ipv6filteroption.ExtAttrs = &map[string]ExtAttrs{}
		}
		for k, v := range a.Client.Cfg.DefaultExtAttrs {
			if _, ok := (*r.ipv6filteroption.ExtAttrs)[k]; !ok {
				(*r.ipv6filteroption.ExtAttrs)[k] = ExtAttrs{
					Value: v.Value,
				}
			}
		}
	}
	// body params
	localVarPostBody = r.ipv6filteroption
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

type Ipv6filteroptionAPIDeleteRequest struct {
	ctx        context.Context
	ApiService Ipv6filteroptionAPI
	reference  string
}

func (r Ipv6filteroptionAPIDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteExecute(r)
}

/*
Delete Delete a ipv6filteroption object

Deletes a specific ipv6filteroption object by reference

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param reference Reference of the ipv6filteroption object
	@return Ipv6filteroptionAPIDeleteRequest
*/
func (a *Ipv6filteroptionAPIService) Delete(ctx context.Context, reference string) Ipv6filteroptionAPIDeleteRequest {
	return Ipv6filteroptionAPIDeleteRequest{
		ApiService: a,
		ctx:        ctx,
		reference:  reference,
	}
}

// Execute executes the request
func (a *Ipv6filteroptionAPIService) DeleteExecute(r Ipv6filteroptionAPIDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []internal.FormFile
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "Ipv6filteroptionAPIService.Delete")
	if err != nil {
		return nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/ipv6filteroption/{reference}"
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

type Ipv6filteroptionAPIListRequest struct {
	ctx              context.Context
	ApiService       Ipv6filteroptionAPI
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
func (r Ipv6filteroptionAPIListRequest) ReturnFields(returnFields string) Ipv6filteroptionAPIListRequest {
	r.returnFields = &returnFields
	return r
}

// Enter the field names followed by comma, this returns the required fields along with the default fields
func (r Ipv6filteroptionAPIListRequest) ReturnFieldsPlus(returnFieldsPlus string) Ipv6filteroptionAPIListRequest {
	r.returnFieldsPlus = &returnFieldsPlus
	return r
}

// Enter the number of results to be fetched
func (r Ipv6filteroptionAPIListRequest) MaxResults(maxResults int32) Ipv6filteroptionAPIListRequest {
	r.maxResults = &maxResults
	return r
}

// Select 1 if result is required as an object
func (r Ipv6filteroptionAPIListRequest) ReturnAsObject(returnAsObject int32) Ipv6filteroptionAPIListRequest {
	r.returnAsObject = &returnAsObject
	return r
}

// Control paging of results
func (r Ipv6filteroptionAPIListRequest) Paging(paging int32) Ipv6filteroptionAPIListRequest {
	r.paging = &paging
	return r
}

// Page id for retrieving next page of results
func (r Ipv6filteroptionAPIListRequest) PageId(pageId string) Ipv6filteroptionAPIListRequest {
	r.pageId = &pageId
	return r
}

func (r Ipv6filteroptionAPIListRequest) Filters(filters map[string]interface{}) Ipv6filteroptionAPIListRequest {
	r.filters = &filters
	return r
}

func (r Ipv6filteroptionAPIListRequest) Extattrfilter(extattrfilter map[string]interface{}) Ipv6filteroptionAPIListRequest {
	r.extattrfilter = &extattrfilter
	return r
}

func (r Ipv6filteroptionAPIListRequest) Execute() (*ListIpv6filteroptionResponse, *http.Response, error) {
	return r.ApiService.ListExecute(r)
}

/*
List Retrieve ipv6filteroption objects

Returns a list of ipv6filteroption objects matching the search criteria

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return Ipv6filteroptionAPIListRequest
*/
func (a *Ipv6filteroptionAPIService) List(ctx context.Context) Ipv6filteroptionAPIListRequest {
	return Ipv6filteroptionAPIListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ListIpv6filteroptionResponse
func (a *Ipv6filteroptionAPIService) ListExecute(r Ipv6filteroptionAPIListRequest) (*ListIpv6filteroptionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *ListIpv6filteroptionResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "Ipv6filteroptionAPIService.List")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/ipv6filteroption"

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

type Ipv6filteroptionAPIReadRequest struct {
	ctx              context.Context
	ApiService       Ipv6filteroptionAPI
	reference        string
	returnFields     *string
	returnFieldsPlus *string
	returnAsObject   *int32
}

// Enter the field names followed by comma
func (r Ipv6filteroptionAPIReadRequest) ReturnFields(returnFields string) Ipv6filteroptionAPIReadRequest {
	r.returnFields = &returnFields
	return r
}

// Enter the field names followed by comma, this returns the required fields along with the default fields
func (r Ipv6filteroptionAPIReadRequest) ReturnFieldsPlus(returnFieldsPlus string) Ipv6filteroptionAPIReadRequest {
	r.returnFieldsPlus = &returnFieldsPlus
	return r
}

// Select 1 if result is required as an object
func (r Ipv6filteroptionAPIReadRequest) ReturnAsObject(returnAsObject int32) Ipv6filteroptionAPIReadRequest {
	r.returnAsObject = &returnAsObject
	return r
}

func (r Ipv6filteroptionAPIReadRequest) Execute() (*GetIpv6filteroptionResponse, *http.Response, error) {
	return r.ApiService.ReadExecute(r)
}

/*
Read Get a specific ipv6filteroption object

Returns a specific ipv6filteroption object by reference

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param reference Reference of the ipv6filteroption object
	@return Ipv6filteroptionAPIReadRequest
*/
func (a *Ipv6filteroptionAPIService) Read(ctx context.Context, reference string) Ipv6filteroptionAPIReadRequest {
	return Ipv6filteroptionAPIReadRequest{
		ApiService: a,
		ctx:        ctx,
		reference:  reference,
	}
}

// Execute executes the request
//
//	@return GetIpv6filteroptionResponse
func (a *Ipv6filteroptionAPIService) ReadExecute(r Ipv6filteroptionAPIReadRequest) (*GetIpv6filteroptionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *GetIpv6filteroptionResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "Ipv6filteroptionAPIService.Read")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/ipv6filteroption/{reference}"
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

type Ipv6filteroptionAPIUpdateRequest struct {
	ctx              context.Context
	ApiService       Ipv6filteroptionAPI
	reference        string
	ipv6filteroption *Ipv6filteroption
	returnFields     *string
	returnFieldsPlus *string
	returnAsObject   *int32
}

// Object data to update
func (r Ipv6filteroptionAPIUpdateRequest) Ipv6filteroption(ipv6filteroption Ipv6filteroption) Ipv6filteroptionAPIUpdateRequest {
	r.ipv6filteroption = &ipv6filteroption
	return r
}

// Enter the field names followed by comma
func (r Ipv6filteroptionAPIUpdateRequest) ReturnFields(returnFields string) Ipv6filteroptionAPIUpdateRequest {
	r.returnFields = &returnFields
	return r
}

// Enter the field names followed by comma, this returns the required fields along with the default fields
func (r Ipv6filteroptionAPIUpdateRequest) ReturnFieldsPlus(returnFieldsPlus string) Ipv6filteroptionAPIUpdateRequest {
	r.returnFieldsPlus = &returnFieldsPlus
	return r
}

// Select 1 if result is required as an object
func (r Ipv6filteroptionAPIUpdateRequest) ReturnAsObject(returnAsObject int32) Ipv6filteroptionAPIUpdateRequest {
	r.returnAsObject = &returnAsObject
	return r
}

func (r Ipv6filteroptionAPIUpdateRequest) Execute() (*UpdateIpv6filteroptionResponse, *http.Response, error) {
	return r.ApiService.UpdateExecute(r)
}

/*
Update Update a ipv6filteroption object

Updates a specific ipv6filteroption object by reference

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param reference Reference of the ipv6filteroption object
	@return Ipv6filteroptionAPIUpdateRequest
*/
func (a *Ipv6filteroptionAPIService) Update(ctx context.Context, reference string) Ipv6filteroptionAPIUpdateRequest {
	return Ipv6filteroptionAPIUpdateRequest{
		ApiService: a,
		ctx:        ctx,
		reference:  reference,
	}
}

// Execute executes the request
//
//	@return UpdateIpv6filteroptionResponse
func (a *Ipv6filteroptionAPIService) UpdateExecute(r Ipv6filteroptionAPIUpdateRequest) (*UpdateIpv6filteroptionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *UpdateIpv6filteroptionResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "Ipv6filteroptionAPIService.Update")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/ipv6filteroption/{reference}"
	localVarPath = strings.Replace(localVarPath, "{"+"reference"+"}", url.PathEscape(internal.ParameterValueToString(r.reference, "reference")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.ipv6filteroption == nil {
		return localVarReturnValue, nil, internal.ReportError("ipv6filteroption is required and must be specified")
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
	if len(a.Client.Cfg.DefaultExtAttrs) > 0 && r.ipv6filteroption != nil {
		if r.ipv6filteroption.ExtAttrs == nil {
			r.ipv6filteroption.ExtAttrs = &map[string]ExtAttrs{}
		}
		for k, v := range a.Client.Cfg.DefaultExtAttrs {
			if _, ok := (*r.ipv6filteroption.ExtAttrs)[k]; !ok {
				(*r.ipv6filteroption.ExtAttrs)[k] = ExtAttrs{
					Value: v.Value,
				}
			}
		}
	}
	// body params
	localVarPostBody = r.ipv6filteroption
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
