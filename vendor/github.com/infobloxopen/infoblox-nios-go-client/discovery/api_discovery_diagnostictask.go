/*
Infoblox DISCOVERY API

OpenAPI specification for Infoblox NIOS WAPI DISCOVERY objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package discovery

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/infobloxopen/infoblox-nios-go-client/internal"
)

type DiscoveryDiagnostictaskAPI interface {
	/*
		List Retrieve discovery:diagnostictask objects

		Returns a list of discovery:diagnostictask objects matching the search criteria

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return DiscoveryDiagnostictaskAPIListRequest
	*/
	List(ctx context.Context) DiscoveryDiagnostictaskAPIListRequest

	// ListExecute executes the request
	//  @return ListDiscoveryDiagnostictaskResponse
	ListExecute(r DiscoveryDiagnostictaskAPIListRequest) (*ListDiscoveryDiagnostictaskResponse, *http.Response, error)
	/*
		Read Get a specific discovery:diagnostictask object

		Returns a specific discovery:diagnostictask object by reference

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param reference Reference of the discovery:diagnostictask object
		@return DiscoveryDiagnostictaskAPIReadRequest
	*/
	Read(ctx context.Context, reference string) DiscoveryDiagnostictaskAPIReadRequest

	// ReadExecute executes the request
	//  @return GetDiscoveryDiagnostictaskResponse
	ReadExecute(r DiscoveryDiagnostictaskAPIReadRequest) (*GetDiscoveryDiagnostictaskResponse, *http.Response, error)
	/*
		Update Update a discovery:diagnostictask object

		Updates a specific discovery:diagnostictask object by reference

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param reference Reference of the discovery:diagnostictask object
		@return DiscoveryDiagnostictaskAPIUpdateRequest
	*/
	Update(ctx context.Context, reference string) DiscoveryDiagnostictaskAPIUpdateRequest

	// UpdateExecute executes the request
	//  @return UpdateDiscoveryDiagnostictaskResponse
	UpdateExecute(r DiscoveryDiagnostictaskAPIUpdateRequest) (*UpdateDiscoveryDiagnostictaskResponse, *http.Response, error)
}

// DiscoveryDiagnostictaskAPIService DiscoveryDiagnostictaskAPI service
type DiscoveryDiagnostictaskAPIService internal.Service

type DiscoveryDiagnostictaskAPIListRequest struct {
	ctx              context.Context
	ApiService       DiscoveryDiagnostictaskAPI
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
func (r DiscoveryDiagnostictaskAPIListRequest) ReturnFields(returnFields string) DiscoveryDiagnostictaskAPIListRequest {
	r.returnFields = &returnFields
	return r
}

// Enter the field names followed by comma, this returns the required fields along with the default fields
func (r DiscoveryDiagnostictaskAPIListRequest) ReturnFieldsPlus(returnFieldsPlus string) DiscoveryDiagnostictaskAPIListRequest {
	r.returnFieldsPlus = &returnFieldsPlus
	return r
}

// Enter the number of results to be fetched
func (r DiscoveryDiagnostictaskAPIListRequest) MaxResults(maxResults int32) DiscoveryDiagnostictaskAPIListRequest {
	r.maxResults = &maxResults
	return r
}

// Select 1 if result is required as an object
func (r DiscoveryDiagnostictaskAPIListRequest) ReturnAsObject(returnAsObject int32) DiscoveryDiagnostictaskAPIListRequest {
	r.returnAsObject = &returnAsObject
	return r
}

// Control paging of results
func (r DiscoveryDiagnostictaskAPIListRequest) Paging(paging int32) DiscoveryDiagnostictaskAPIListRequest {
	r.paging = &paging
	return r
}

// Page id for retrieving next page of results
func (r DiscoveryDiagnostictaskAPIListRequest) PageId(pageId string) DiscoveryDiagnostictaskAPIListRequest {
	r.pageId = &pageId
	return r
}

func (r DiscoveryDiagnostictaskAPIListRequest) Filters(filters map[string]interface{}) DiscoveryDiagnostictaskAPIListRequest {
	r.filters = &filters
	return r
}

func (r DiscoveryDiagnostictaskAPIListRequest) Extattrfilter(extattrfilter map[string]interface{}) DiscoveryDiagnostictaskAPIListRequest {
	r.extattrfilter = &extattrfilter
	return r
}

func (r DiscoveryDiagnostictaskAPIListRequest) Execute() (*ListDiscoveryDiagnostictaskResponse, *http.Response, error) {
	return r.ApiService.ListExecute(r)
}

/*
List Retrieve discovery:diagnostictask objects

Returns a list of discovery:diagnostictask objects matching the search criteria

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return DiscoveryDiagnostictaskAPIListRequest
*/
func (a *DiscoveryDiagnostictaskAPIService) List(ctx context.Context) DiscoveryDiagnostictaskAPIListRequest {
	return DiscoveryDiagnostictaskAPIListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ListDiscoveryDiagnostictaskResponse
func (a *DiscoveryDiagnostictaskAPIService) ListExecute(r DiscoveryDiagnostictaskAPIListRequest) (*ListDiscoveryDiagnostictaskResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *ListDiscoveryDiagnostictaskResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "DiscoveryDiagnostictaskAPIService.List")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/discovery:diagnostictask"

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

type DiscoveryDiagnostictaskAPIReadRequest struct {
	ctx              context.Context
	ApiService       DiscoveryDiagnostictaskAPI
	reference        string
	returnFields     *string
	returnFieldsPlus *string
	returnAsObject   *int32
}

// Enter the field names followed by comma
func (r DiscoveryDiagnostictaskAPIReadRequest) ReturnFields(returnFields string) DiscoveryDiagnostictaskAPIReadRequest {
	r.returnFields = &returnFields
	return r
}

// Enter the field names followed by comma, this returns the required fields along with the default fields
func (r DiscoveryDiagnostictaskAPIReadRequest) ReturnFieldsPlus(returnFieldsPlus string) DiscoveryDiagnostictaskAPIReadRequest {
	r.returnFieldsPlus = &returnFieldsPlus
	return r
}

// Select 1 if result is required as an object
func (r DiscoveryDiagnostictaskAPIReadRequest) ReturnAsObject(returnAsObject int32) DiscoveryDiagnostictaskAPIReadRequest {
	r.returnAsObject = &returnAsObject
	return r
}

func (r DiscoveryDiagnostictaskAPIReadRequest) Execute() (*GetDiscoveryDiagnostictaskResponse, *http.Response, error) {
	return r.ApiService.ReadExecute(r)
}

/*
Read Get a specific discovery:diagnostictask object

Returns a specific discovery:diagnostictask object by reference

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param reference Reference of the discovery:diagnostictask object
	@return DiscoveryDiagnostictaskAPIReadRequest
*/
func (a *DiscoveryDiagnostictaskAPIService) Read(ctx context.Context, reference string) DiscoveryDiagnostictaskAPIReadRequest {
	return DiscoveryDiagnostictaskAPIReadRequest{
		ApiService: a,
		ctx:        ctx,
		reference:  reference,
	}
}

// Execute executes the request
//
//	@return GetDiscoveryDiagnostictaskResponse
func (a *DiscoveryDiagnostictaskAPIService) ReadExecute(r DiscoveryDiagnostictaskAPIReadRequest) (*GetDiscoveryDiagnostictaskResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *GetDiscoveryDiagnostictaskResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "DiscoveryDiagnostictaskAPIService.Read")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/discovery:diagnostictask/{reference}"
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

type DiscoveryDiagnostictaskAPIUpdateRequest struct {
	ctx                     context.Context
	ApiService              DiscoveryDiagnostictaskAPI
	reference               string
	discoveryDiagnostictask *DiscoveryDiagnostictask
	returnFields            *string
	returnFieldsPlus        *string
	returnAsObject          *int32
}

// Object data to update
func (r DiscoveryDiagnostictaskAPIUpdateRequest) DiscoveryDiagnostictask(discoveryDiagnostictask DiscoveryDiagnostictask) DiscoveryDiagnostictaskAPIUpdateRequest {
	r.discoveryDiagnostictask = &discoveryDiagnostictask
	return r
}

// Enter the field names followed by comma
func (r DiscoveryDiagnostictaskAPIUpdateRequest) ReturnFields(returnFields string) DiscoveryDiagnostictaskAPIUpdateRequest {
	r.returnFields = &returnFields
	return r
}

// Enter the field names followed by comma, this returns the required fields along with the default fields
func (r DiscoveryDiagnostictaskAPIUpdateRequest) ReturnFieldsPlus(returnFieldsPlus string) DiscoveryDiagnostictaskAPIUpdateRequest {
	r.returnFieldsPlus = &returnFieldsPlus
	return r
}

// Select 1 if result is required as an object
func (r DiscoveryDiagnostictaskAPIUpdateRequest) ReturnAsObject(returnAsObject int32) DiscoveryDiagnostictaskAPIUpdateRequest {
	r.returnAsObject = &returnAsObject
	return r
}

func (r DiscoveryDiagnostictaskAPIUpdateRequest) Execute() (*UpdateDiscoveryDiagnostictaskResponse, *http.Response, error) {
	return r.ApiService.UpdateExecute(r)
}

/*
Update Update a discovery:diagnostictask object

Updates a specific discovery:diagnostictask object by reference

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param reference Reference of the discovery:diagnostictask object
	@return DiscoveryDiagnostictaskAPIUpdateRequest
*/
func (a *DiscoveryDiagnostictaskAPIService) Update(ctx context.Context, reference string) DiscoveryDiagnostictaskAPIUpdateRequest {
	return DiscoveryDiagnostictaskAPIUpdateRequest{
		ApiService: a,
		ctx:        ctx,
		reference:  reference,
	}
}

// Execute executes the request
//
//	@return UpdateDiscoveryDiagnostictaskResponse
func (a *DiscoveryDiagnostictaskAPIService) UpdateExecute(r DiscoveryDiagnostictaskAPIUpdateRequest) (*UpdateDiscoveryDiagnostictaskResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *UpdateDiscoveryDiagnostictaskResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "DiscoveryDiagnostictaskAPIService.Update")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/discovery:diagnostictask/{reference}"
	localVarPath = strings.Replace(localVarPath, "{"+"reference"+"}", url.PathEscape(internal.ParameterValueToString(r.reference, "reference")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.discoveryDiagnostictask == nil {
		return localVarReturnValue, nil, internal.ReportError("discoveryDiagnostictask is required and must be specified")
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
	// body params
	localVarPostBody = r.discoveryDiagnostictask
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
