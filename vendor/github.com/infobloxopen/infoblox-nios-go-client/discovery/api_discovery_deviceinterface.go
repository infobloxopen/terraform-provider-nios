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

type DiscoveryDeviceinterfaceAPI interface {
	/*
		List Retrieve discovery:deviceinterface objects

		Returns a list of discovery:deviceinterface objects matching the search criteria

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return DiscoveryDeviceinterfaceAPIListRequest
	*/
	List(ctx context.Context) DiscoveryDeviceinterfaceAPIListRequest

	// ListExecute executes the request
	//  @return ListDiscoveryDeviceinterfaceResponse
	ListExecute(r DiscoveryDeviceinterfaceAPIListRequest) (*ListDiscoveryDeviceinterfaceResponse, *http.Response, error)
	/*
		Read Get a specific discovery:deviceinterface object

		Returns a specific discovery:deviceinterface object by reference

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param reference Reference of the discovery:deviceinterface object
		@return DiscoveryDeviceinterfaceAPIReadRequest
	*/
	Read(ctx context.Context, reference string) DiscoveryDeviceinterfaceAPIReadRequest

	// ReadExecute executes the request
	//  @return GetDiscoveryDeviceinterfaceResponse
	ReadExecute(r DiscoveryDeviceinterfaceAPIReadRequest) (*GetDiscoveryDeviceinterfaceResponse, *http.Response, error)
	/*
		Update Update a discovery:deviceinterface object

		Updates a specific discovery:deviceinterface object by reference

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param reference Reference of the discovery:deviceinterface object
		@return DiscoveryDeviceinterfaceAPIUpdateRequest
	*/
	Update(ctx context.Context, reference string) DiscoveryDeviceinterfaceAPIUpdateRequest

	// UpdateExecute executes the request
	//  @return UpdateDiscoveryDeviceinterfaceResponse
	UpdateExecute(r DiscoveryDeviceinterfaceAPIUpdateRequest) (*UpdateDiscoveryDeviceinterfaceResponse, *http.Response, error)
}

// DiscoveryDeviceinterfaceAPIService DiscoveryDeviceinterfaceAPI service
type DiscoveryDeviceinterfaceAPIService internal.Service

type DiscoveryDeviceinterfaceAPIListRequest struct {
	ctx              context.Context
	ApiService       DiscoveryDeviceinterfaceAPI
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
func (r DiscoveryDeviceinterfaceAPIListRequest) ReturnFields(returnFields string) DiscoveryDeviceinterfaceAPIListRequest {
	r.returnFields = &returnFields
	return r
}

// Enter the field names followed by comma, this returns the required fields along with the default fields
func (r DiscoveryDeviceinterfaceAPIListRequest) ReturnFieldsPlus(returnFieldsPlus string) DiscoveryDeviceinterfaceAPIListRequest {
	r.returnFieldsPlus = &returnFieldsPlus
	return r
}

// Enter the number of results to be fetched
func (r DiscoveryDeviceinterfaceAPIListRequest) MaxResults(maxResults int32) DiscoveryDeviceinterfaceAPIListRequest {
	r.maxResults = &maxResults
	return r
}

// Select 1 if result is required as an object
func (r DiscoveryDeviceinterfaceAPIListRequest) ReturnAsObject(returnAsObject int32) DiscoveryDeviceinterfaceAPIListRequest {
	r.returnAsObject = &returnAsObject
	return r
}

// Control paging of results
func (r DiscoveryDeviceinterfaceAPIListRequest) Paging(paging int32) DiscoveryDeviceinterfaceAPIListRequest {
	r.paging = &paging
	return r
}

// Page id for retrieving next page of results
func (r DiscoveryDeviceinterfaceAPIListRequest) PageId(pageId string) DiscoveryDeviceinterfaceAPIListRequest {
	r.pageId = &pageId
	return r
}

func (r DiscoveryDeviceinterfaceAPIListRequest) Filters(filters map[string]interface{}) DiscoveryDeviceinterfaceAPIListRequest {
	r.filters = &filters
	return r
}

func (r DiscoveryDeviceinterfaceAPIListRequest) Extattrfilter(extattrfilter map[string]interface{}) DiscoveryDeviceinterfaceAPIListRequest {
	r.extattrfilter = &extattrfilter
	return r
}

func (r DiscoveryDeviceinterfaceAPIListRequest) Execute() (*ListDiscoveryDeviceinterfaceResponse, *http.Response, error) {
	return r.ApiService.ListExecute(r)
}

/*
List Retrieve discovery:deviceinterface objects

Returns a list of discovery:deviceinterface objects matching the search criteria

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return DiscoveryDeviceinterfaceAPIListRequest
*/
func (a *DiscoveryDeviceinterfaceAPIService) List(ctx context.Context) DiscoveryDeviceinterfaceAPIListRequest {
	return DiscoveryDeviceinterfaceAPIListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ListDiscoveryDeviceinterfaceResponse
func (a *DiscoveryDeviceinterfaceAPIService) ListExecute(r DiscoveryDeviceinterfaceAPIListRequest) (*ListDiscoveryDeviceinterfaceResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *ListDiscoveryDeviceinterfaceResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "DiscoveryDeviceinterfaceAPIService.List")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/discovery:deviceinterface"

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

type DiscoveryDeviceinterfaceAPIReadRequest struct {
	ctx              context.Context
	ApiService       DiscoveryDeviceinterfaceAPI
	reference        string
	returnFields     *string
	returnFieldsPlus *string
	returnAsObject   *int32
}

// Enter the field names followed by comma
func (r DiscoveryDeviceinterfaceAPIReadRequest) ReturnFields(returnFields string) DiscoveryDeviceinterfaceAPIReadRequest {
	r.returnFields = &returnFields
	return r
}

// Enter the field names followed by comma, this returns the required fields along with the default fields
func (r DiscoveryDeviceinterfaceAPIReadRequest) ReturnFieldsPlus(returnFieldsPlus string) DiscoveryDeviceinterfaceAPIReadRequest {
	r.returnFieldsPlus = &returnFieldsPlus
	return r
}

// Select 1 if result is required as an object
func (r DiscoveryDeviceinterfaceAPIReadRequest) ReturnAsObject(returnAsObject int32) DiscoveryDeviceinterfaceAPIReadRequest {
	r.returnAsObject = &returnAsObject
	return r
}

func (r DiscoveryDeviceinterfaceAPIReadRequest) Execute() (*GetDiscoveryDeviceinterfaceResponse, *http.Response, error) {
	return r.ApiService.ReadExecute(r)
}

/*
Read Get a specific discovery:deviceinterface object

Returns a specific discovery:deviceinterface object by reference

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param reference Reference of the discovery:deviceinterface object
	@return DiscoveryDeviceinterfaceAPIReadRequest
*/
func (a *DiscoveryDeviceinterfaceAPIService) Read(ctx context.Context, reference string) DiscoveryDeviceinterfaceAPIReadRequest {
	return DiscoveryDeviceinterfaceAPIReadRequest{
		ApiService: a,
		ctx:        ctx,
		reference:  reference,
	}
}

// Execute executes the request
//
//	@return GetDiscoveryDeviceinterfaceResponse
func (a *DiscoveryDeviceinterfaceAPIService) ReadExecute(r DiscoveryDeviceinterfaceAPIReadRequest) (*GetDiscoveryDeviceinterfaceResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *GetDiscoveryDeviceinterfaceResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "DiscoveryDeviceinterfaceAPIService.Read")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/discovery:deviceinterface/{reference}"
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

type DiscoveryDeviceinterfaceAPIUpdateRequest struct {
	ctx                      context.Context
	ApiService               DiscoveryDeviceinterfaceAPI
	reference                string
	discoveryDeviceinterface *DiscoveryDeviceinterface
	returnFields             *string
	returnFieldsPlus         *string
	returnAsObject           *int32
}

// Object data to update
func (r DiscoveryDeviceinterfaceAPIUpdateRequest) DiscoveryDeviceinterface(discoveryDeviceinterface DiscoveryDeviceinterface) DiscoveryDeviceinterfaceAPIUpdateRequest {
	r.discoveryDeviceinterface = &discoveryDeviceinterface
	return r
}

// Enter the field names followed by comma
func (r DiscoveryDeviceinterfaceAPIUpdateRequest) ReturnFields(returnFields string) DiscoveryDeviceinterfaceAPIUpdateRequest {
	r.returnFields = &returnFields
	return r
}

// Enter the field names followed by comma, this returns the required fields along with the default fields
func (r DiscoveryDeviceinterfaceAPIUpdateRequest) ReturnFieldsPlus(returnFieldsPlus string) DiscoveryDeviceinterfaceAPIUpdateRequest {
	r.returnFieldsPlus = &returnFieldsPlus
	return r
}

// Select 1 if result is required as an object
func (r DiscoveryDeviceinterfaceAPIUpdateRequest) ReturnAsObject(returnAsObject int32) DiscoveryDeviceinterfaceAPIUpdateRequest {
	r.returnAsObject = &returnAsObject
	return r
}

func (r DiscoveryDeviceinterfaceAPIUpdateRequest) Execute() (*UpdateDiscoveryDeviceinterfaceResponse, *http.Response, error) {
	return r.ApiService.UpdateExecute(r)
}

/*
Update Update a discovery:deviceinterface object

Updates a specific discovery:deviceinterface object by reference

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param reference Reference of the discovery:deviceinterface object
	@return DiscoveryDeviceinterfaceAPIUpdateRequest
*/
func (a *DiscoveryDeviceinterfaceAPIService) Update(ctx context.Context, reference string) DiscoveryDeviceinterfaceAPIUpdateRequest {
	return DiscoveryDeviceinterfaceAPIUpdateRequest{
		ApiService: a,
		ctx:        ctx,
		reference:  reference,
	}
}

// Execute executes the request
//
//	@return UpdateDiscoveryDeviceinterfaceResponse
func (a *DiscoveryDeviceinterfaceAPIService) UpdateExecute(r DiscoveryDeviceinterfaceAPIUpdateRequest) (*UpdateDiscoveryDeviceinterfaceResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *UpdateDiscoveryDeviceinterfaceResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "DiscoveryDeviceinterfaceAPIService.Update")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/discovery:deviceinterface/{reference}"
	localVarPath = strings.Replace(localVarPath, "{"+"reference"+"}", url.PathEscape(internal.ParameterValueToString(r.reference, "reference")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.discoveryDeviceinterface == nil {
		return localVarReturnValue, nil, internal.ReportError("discoveryDeviceinterface is required and must be specified")
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
	if len(a.Client.Cfg.DefaultExtAttrs) > 0 && r.discoveryDeviceinterface != nil {
		if r.discoveryDeviceinterface.ExtAttrs == nil {
			r.discoveryDeviceinterface.ExtAttrs = &map[string]ExtAttrs{}
		}
		for k, v := range a.Client.Cfg.DefaultExtAttrs {
			if _, ok := (*r.discoveryDeviceinterface.ExtAttrs)[k]; !ok {
				(*r.discoveryDeviceinterface.ExtAttrs)[k] = ExtAttrs{
					Value: v.Value,
				}
			}
		}
	}
	// body params
	localVarPostBody = r.discoveryDeviceinterface
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
