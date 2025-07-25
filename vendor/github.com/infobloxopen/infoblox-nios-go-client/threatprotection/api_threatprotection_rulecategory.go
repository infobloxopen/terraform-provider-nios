/*
Infoblox THREATPROTECTION API

OpenAPI specification for Infoblox NIOS WAPI THREATPROTECTION objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package threatprotection

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/infobloxopen/infoblox-nios-go-client/internal"
)

type ThreatprotectionRulecategoryAPI interface {
	/*
		List Retrieve threatprotection:rulecategory objects

		Returns a list of threatprotection:rulecategory objects matching the search criteria

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ThreatprotectionRulecategoryAPIListRequest
	*/
	List(ctx context.Context) ThreatprotectionRulecategoryAPIListRequest

	// ListExecute executes the request
	//  @return ListThreatprotectionRulecategoryResponse
	ListExecute(r ThreatprotectionRulecategoryAPIListRequest) (*ListThreatprotectionRulecategoryResponse, *http.Response, error)
	/*
		Read Get a specific threatprotection:rulecategory object

		Returns a specific threatprotection:rulecategory object by reference

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param reference Reference of the threatprotection:rulecategory object
		@return ThreatprotectionRulecategoryAPIReadRequest
	*/
	Read(ctx context.Context, reference string) ThreatprotectionRulecategoryAPIReadRequest

	// ReadExecute executes the request
	//  @return GetThreatprotectionRulecategoryResponse
	ReadExecute(r ThreatprotectionRulecategoryAPIReadRequest) (*GetThreatprotectionRulecategoryResponse, *http.Response, error)
}

// ThreatprotectionRulecategoryAPIService ThreatprotectionRulecategoryAPI service
type ThreatprotectionRulecategoryAPIService internal.Service

type ThreatprotectionRulecategoryAPIListRequest struct {
	ctx              context.Context
	ApiService       ThreatprotectionRulecategoryAPI
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
func (r ThreatprotectionRulecategoryAPIListRequest) ReturnFields(returnFields string) ThreatprotectionRulecategoryAPIListRequest {
	r.returnFields = &returnFields
	return r
}

// Enter the field names followed by comma, this returns the required fields along with the default fields
func (r ThreatprotectionRulecategoryAPIListRequest) ReturnFieldsPlus(returnFieldsPlus string) ThreatprotectionRulecategoryAPIListRequest {
	r.returnFieldsPlus = &returnFieldsPlus
	return r
}

// Enter the number of results to be fetched
func (r ThreatprotectionRulecategoryAPIListRequest) MaxResults(maxResults int32) ThreatprotectionRulecategoryAPIListRequest {
	r.maxResults = &maxResults
	return r
}

// Select 1 if result is required as an object
func (r ThreatprotectionRulecategoryAPIListRequest) ReturnAsObject(returnAsObject int32) ThreatprotectionRulecategoryAPIListRequest {
	r.returnAsObject = &returnAsObject
	return r
}

// Control paging of results
func (r ThreatprotectionRulecategoryAPIListRequest) Paging(paging int32) ThreatprotectionRulecategoryAPIListRequest {
	r.paging = &paging
	return r
}

// Page id for retrieving next page of results
func (r ThreatprotectionRulecategoryAPIListRequest) PageId(pageId string) ThreatprotectionRulecategoryAPIListRequest {
	r.pageId = &pageId
	return r
}

func (r ThreatprotectionRulecategoryAPIListRequest) Filters(filters map[string]interface{}) ThreatprotectionRulecategoryAPIListRequest {
	r.filters = &filters
	return r
}

func (r ThreatprotectionRulecategoryAPIListRequest) Extattrfilter(extattrfilter map[string]interface{}) ThreatprotectionRulecategoryAPIListRequest {
	r.extattrfilter = &extattrfilter
	return r
}

func (r ThreatprotectionRulecategoryAPIListRequest) Execute() (*ListThreatprotectionRulecategoryResponse, *http.Response, error) {
	return r.ApiService.ListExecute(r)
}

/*
List Retrieve threatprotection:rulecategory objects

Returns a list of threatprotection:rulecategory objects matching the search criteria

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ThreatprotectionRulecategoryAPIListRequest
*/
func (a *ThreatprotectionRulecategoryAPIService) List(ctx context.Context) ThreatprotectionRulecategoryAPIListRequest {
	return ThreatprotectionRulecategoryAPIListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ListThreatprotectionRulecategoryResponse
func (a *ThreatprotectionRulecategoryAPIService) ListExecute(r ThreatprotectionRulecategoryAPIListRequest) (*ListThreatprotectionRulecategoryResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *ListThreatprotectionRulecategoryResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "ThreatprotectionRulecategoryAPIService.List")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/threatprotection:rulecategory"

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

type ThreatprotectionRulecategoryAPIReadRequest struct {
	ctx              context.Context
	ApiService       ThreatprotectionRulecategoryAPI
	reference        string
	returnFields     *string
	returnFieldsPlus *string
	returnAsObject   *int32
}

// Enter the field names followed by comma
func (r ThreatprotectionRulecategoryAPIReadRequest) ReturnFields(returnFields string) ThreatprotectionRulecategoryAPIReadRequest {
	r.returnFields = &returnFields
	return r
}

// Enter the field names followed by comma, this returns the required fields along with the default fields
func (r ThreatprotectionRulecategoryAPIReadRequest) ReturnFieldsPlus(returnFieldsPlus string) ThreatprotectionRulecategoryAPIReadRequest {
	r.returnFieldsPlus = &returnFieldsPlus
	return r
}

// Select 1 if result is required as an object
func (r ThreatprotectionRulecategoryAPIReadRequest) ReturnAsObject(returnAsObject int32) ThreatprotectionRulecategoryAPIReadRequest {
	r.returnAsObject = &returnAsObject
	return r
}

func (r ThreatprotectionRulecategoryAPIReadRequest) Execute() (*GetThreatprotectionRulecategoryResponse, *http.Response, error) {
	return r.ApiService.ReadExecute(r)
}

/*
Read Get a specific threatprotection:rulecategory object

Returns a specific threatprotection:rulecategory object by reference

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param reference Reference of the threatprotection:rulecategory object
	@return ThreatprotectionRulecategoryAPIReadRequest
*/
func (a *ThreatprotectionRulecategoryAPIService) Read(ctx context.Context, reference string) ThreatprotectionRulecategoryAPIReadRequest {
	return ThreatprotectionRulecategoryAPIReadRequest{
		ApiService: a,
		ctx:        ctx,
		reference:  reference,
	}
}

// Execute executes the request
//
//	@return GetThreatprotectionRulecategoryResponse
func (a *ThreatprotectionRulecategoryAPIService) ReadExecute(r ThreatprotectionRulecategoryAPIReadRequest) (*GetThreatprotectionRulecategoryResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *GetThreatprotectionRulecategoryResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "ThreatprotectionRulecategoryAPIService.Read")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/threatprotection:rulecategory/{reference}"
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
