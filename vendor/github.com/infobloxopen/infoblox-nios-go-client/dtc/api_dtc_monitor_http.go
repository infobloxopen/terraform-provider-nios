/*
Infoblox DTC API

OpenAPI specification for Infoblox NIOS WAPI DTC objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dtc

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/infobloxopen/infoblox-nios-go-client/internal"
)

type DtcMonitorHttpAPI interface {
	/*
		Create Create a dtc:monitor:http object

		Creates a new dtc:monitor:http object

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return DtcMonitorHttpAPICreateRequest
	*/
	Create(ctx context.Context) DtcMonitorHttpAPICreateRequest

	// CreateExecute executes the request
	//  @return CreateDtcMonitorHttpResponse
	CreateExecute(r DtcMonitorHttpAPICreateRequest) (*CreateDtcMonitorHttpResponse, *http.Response, error)
	/*
		Delete Delete a dtc:monitor:http object

		Deletes a specific dtc:monitor:http object by reference

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param reference Reference of the dtc:monitor:http object
		@return DtcMonitorHttpAPIDeleteRequest
	*/
	Delete(ctx context.Context, reference string) DtcMonitorHttpAPIDeleteRequest

	// DeleteExecute executes the request
	DeleteExecute(r DtcMonitorHttpAPIDeleteRequest) (*http.Response, error)
	/*
		List Retrieve dtc:monitor:http objects

		Returns a list of dtc:monitor:http objects matching the search criteria

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return DtcMonitorHttpAPIListRequest
	*/
	List(ctx context.Context) DtcMonitorHttpAPIListRequest

	// ListExecute executes the request
	//  @return ListDtcMonitorHttpResponse
	ListExecute(r DtcMonitorHttpAPIListRequest) (*ListDtcMonitorHttpResponse, *http.Response, error)
	/*
		Read Get a specific dtc:monitor:http object

		Returns a specific dtc:monitor:http object by reference

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param reference Reference of the dtc:monitor:http object
		@return DtcMonitorHttpAPIReadRequest
	*/
	Read(ctx context.Context, reference string) DtcMonitorHttpAPIReadRequest

	// ReadExecute executes the request
	//  @return GetDtcMonitorHttpResponse
	ReadExecute(r DtcMonitorHttpAPIReadRequest) (*GetDtcMonitorHttpResponse, *http.Response, error)
	/*
		Update Update a dtc:monitor:http object

		Updates a specific dtc:monitor:http object by reference

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param reference Reference of the dtc:monitor:http object
		@return DtcMonitorHttpAPIUpdateRequest
	*/
	Update(ctx context.Context, reference string) DtcMonitorHttpAPIUpdateRequest

	// UpdateExecute executes the request
	//  @return UpdateDtcMonitorHttpResponse
	UpdateExecute(r DtcMonitorHttpAPIUpdateRequest) (*UpdateDtcMonitorHttpResponse, *http.Response, error)
}

// DtcMonitorHttpAPIService DtcMonitorHttpAPI service
type DtcMonitorHttpAPIService internal.Service

type DtcMonitorHttpAPICreateRequest struct {
	ctx              context.Context
	ApiService       DtcMonitorHttpAPI
	dtcMonitorHttp   *DtcMonitorHttp
	returnFields     *string
	returnFieldsPlus *string
	returnAsObject   *int32
}

// Object data to create
func (r DtcMonitorHttpAPICreateRequest) DtcMonitorHttp(dtcMonitorHttp DtcMonitorHttp) DtcMonitorHttpAPICreateRequest {
	r.dtcMonitorHttp = &dtcMonitorHttp
	return r
}

// Enter the field names followed by comma
func (r DtcMonitorHttpAPICreateRequest) ReturnFields(returnFields string) DtcMonitorHttpAPICreateRequest {
	r.returnFields = &returnFields
	return r
}

// Enter the field names followed by comma, this returns the required fields along with the default fields
func (r DtcMonitorHttpAPICreateRequest) ReturnFieldsPlus(returnFieldsPlus string) DtcMonitorHttpAPICreateRequest {
	r.returnFieldsPlus = &returnFieldsPlus
	return r
}

// Select 1 if result is required as an object
func (r DtcMonitorHttpAPICreateRequest) ReturnAsObject(returnAsObject int32) DtcMonitorHttpAPICreateRequest {
	r.returnAsObject = &returnAsObject
	return r
}

func (r DtcMonitorHttpAPICreateRequest) Execute() (*CreateDtcMonitorHttpResponse, *http.Response, error) {
	return r.ApiService.CreateExecute(r)
}

/*
Create Create a dtc:monitor:http object

Creates a new dtc:monitor:http object

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return DtcMonitorHttpAPICreateRequest
*/
func (a *DtcMonitorHttpAPIService) Create(ctx context.Context) DtcMonitorHttpAPICreateRequest {
	return DtcMonitorHttpAPICreateRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CreateDtcMonitorHttpResponse
func (a *DtcMonitorHttpAPIService) CreateExecute(r DtcMonitorHttpAPICreateRequest) (*CreateDtcMonitorHttpResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *CreateDtcMonitorHttpResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "DtcMonitorHttpAPIService.Create")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/dtc:monitor:http"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.dtcMonitorHttp == nil {
		return localVarReturnValue, nil, internal.ReportError("dtcMonitorHttp is required and must be specified")
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
	if len(a.Client.Cfg.DefaultExtAttrs) > 0 && r.dtcMonitorHttp != nil {
		if r.dtcMonitorHttp.ExtAttrs == nil {
			r.dtcMonitorHttp.ExtAttrs = &map[string]ExtAttrs{}
		}
		for k, v := range a.Client.Cfg.DefaultExtAttrs {
			if _, ok := (*r.dtcMonitorHttp.ExtAttrs)[k]; !ok {
				(*r.dtcMonitorHttp.ExtAttrs)[k] = ExtAttrs{
					Value: v.Value,
				}
			}
		}
	}
	// body params
	localVarPostBody = r.dtcMonitorHttp
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

type DtcMonitorHttpAPIDeleteRequest struct {
	ctx        context.Context
	ApiService DtcMonitorHttpAPI
	reference  string
}

func (r DtcMonitorHttpAPIDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteExecute(r)
}

/*
Delete Delete a dtc:monitor:http object

Deletes a specific dtc:monitor:http object by reference

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param reference Reference of the dtc:monitor:http object
	@return DtcMonitorHttpAPIDeleteRequest
*/
func (a *DtcMonitorHttpAPIService) Delete(ctx context.Context, reference string) DtcMonitorHttpAPIDeleteRequest {
	return DtcMonitorHttpAPIDeleteRequest{
		ApiService: a,
		ctx:        ctx,
		reference:  reference,
	}
}

// Execute executes the request
func (a *DtcMonitorHttpAPIService) DeleteExecute(r DtcMonitorHttpAPIDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []internal.FormFile
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "DtcMonitorHttpAPIService.Delete")
	if err != nil {
		return nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/dtc:monitor:http/{reference}"
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

type DtcMonitorHttpAPIListRequest struct {
	ctx              context.Context
	ApiService       DtcMonitorHttpAPI
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
func (r DtcMonitorHttpAPIListRequest) ReturnFields(returnFields string) DtcMonitorHttpAPIListRequest {
	r.returnFields = &returnFields
	return r
}

// Enter the field names followed by comma, this returns the required fields along with the default fields
func (r DtcMonitorHttpAPIListRequest) ReturnFieldsPlus(returnFieldsPlus string) DtcMonitorHttpAPIListRequest {
	r.returnFieldsPlus = &returnFieldsPlus
	return r
}

// Enter the number of results to be fetched
func (r DtcMonitorHttpAPIListRequest) MaxResults(maxResults int32) DtcMonitorHttpAPIListRequest {
	r.maxResults = &maxResults
	return r
}

// Select 1 if result is required as an object
func (r DtcMonitorHttpAPIListRequest) ReturnAsObject(returnAsObject int32) DtcMonitorHttpAPIListRequest {
	r.returnAsObject = &returnAsObject
	return r
}

// Control paging of results
func (r DtcMonitorHttpAPIListRequest) Paging(paging int32) DtcMonitorHttpAPIListRequest {
	r.paging = &paging
	return r
}

// Page id for retrieving next page of results
func (r DtcMonitorHttpAPIListRequest) PageId(pageId string) DtcMonitorHttpAPIListRequest {
	r.pageId = &pageId
	return r
}

func (r DtcMonitorHttpAPIListRequest) Filters(filters map[string]interface{}) DtcMonitorHttpAPIListRequest {
	r.filters = &filters
	return r
}

func (r DtcMonitorHttpAPIListRequest) Extattrfilter(extattrfilter map[string]interface{}) DtcMonitorHttpAPIListRequest {
	r.extattrfilter = &extattrfilter
	return r
}

func (r DtcMonitorHttpAPIListRequest) Execute() (*ListDtcMonitorHttpResponse, *http.Response, error) {
	return r.ApiService.ListExecute(r)
}

/*
List Retrieve dtc:monitor:http objects

Returns a list of dtc:monitor:http objects matching the search criteria

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return DtcMonitorHttpAPIListRequest
*/
func (a *DtcMonitorHttpAPIService) List(ctx context.Context) DtcMonitorHttpAPIListRequest {
	return DtcMonitorHttpAPIListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ListDtcMonitorHttpResponse
func (a *DtcMonitorHttpAPIService) ListExecute(r DtcMonitorHttpAPIListRequest) (*ListDtcMonitorHttpResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *ListDtcMonitorHttpResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "DtcMonitorHttpAPIService.List")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/dtc:monitor:http"

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

type DtcMonitorHttpAPIReadRequest struct {
	ctx              context.Context
	ApiService       DtcMonitorHttpAPI
	reference        string
	returnFields     *string
	returnFieldsPlus *string
	returnAsObject   *int32
}

// Enter the field names followed by comma
func (r DtcMonitorHttpAPIReadRequest) ReturnFields(returnFields string) DtcMonitorHttpAPIReadRequest {
	r.returnFields = &returnFields
	return r
}

// Enter the field names followed by comma, this returns the required fields along with the default fields
func (r DtcMonitorHttpAPIReadRequest) ReturnFieldsPlus(returnFieldsPlus string) DtcMonitorHttpAPIReadRequest {
	r.returnFieldsPlus = &returnFieldsPlus
	return r
}

// Select 1 if result is required as an object
func (r DtcMonitorHttpAPIReadRequest) ReturnAsObject(returnAsObject int32) DtcMonitorHttpAPIReadRequest {
	r.returnAsObject = &returnAsObject
	return r
}

func (r DtcMonitorHttpAPIReadRequest) Execute() (*GetDtcMonitorHttpResponse, *http.Response, error) {
	return r.ApiService.ReadExecute(r)
}

/*
Read Get a specific dtc:monitor:http object

Returns a specific dtc:monitor:http object by reference

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param reference Reference of the dtc:monitor:http object
	@return DtcMonitorHttpAPIReadRequest
*/
func (a *DtcMonitorHttpAPIService) Read(ctx context.Context, reference string) DtcMonitorHttpAPIReadRequest {
	return DtcMonitorHttpAPIReadRequest{
		ApiService: a,
		ctx:        ctx,
		reference:  reference,
	}
}

// Execute executes the request
//
//	@return GetDtcMonitorHttpResponse
func (a *DtcMonitorHttpAPIService) ReadExecute(r DtcMonitorHttpAPIReadRequest) (*GetDtcMonitorHttpResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *GetDtcMonitorHttpResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "DtcMonitorHttpAPIService.Read")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/dtc:monitor:http/{reference}"
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

type DtcMonitorHttpAPIUpdateRequest struct {
	ctx              context.Context
	ApiService       DtcMonitorHttpAPI
	reference        string
	dtcMonitorHttp   *DtcMonitorHttp
	returnFields     *string
	returnFieldsPlus *string
	returnAsObject   *int32
}

// Object data to update
func (r DtcMonitorHttpAPIUpdateRequest) DtcMonitorHttp(dtcMonitorHttp DtcMonitorHttp) DtcMonitorHttpAPIUpdateRequest {
	r.dtcMonitorHttp = &dtcMonitorHttp
	return r
}

// Enter the field names followed by comma
func (r DtcMonitorHttpAPIUpdateRequest) ReturnFields(returnFields string) DtcMonitorHttpAPIUpdateRequest {
	r.returnFields = &returnFields
	return r
}

// Enter the field names followed by comma, this returns the required fields along with the default fields
func (r DtcMonitorHttpAPIUpdateRequest) ReturnFieldsPlus(returnFieldsPlus string) DtcMonitorHttpAPIUpdateRequest {
	r.returnFieldsPlus = &returnFieldsPlus
	return r
}

// Select 1 if result is required as an object
func (r DtcMonitorHttpAPIUpdateRequest) ReturnAsObject(returnAsObject int32) DtcMonitorHttpAPIUpdateRequest {
	r.returnAsObject = &returnAsObject
	return r
}

func (r DtcMonitorHttpAPIUpdateRequest) Execute() (*UpdateDtcMonitorHttpResponse, *http.Response, error) {
	return r.ApiService.UpdateExecute(r)
}

/*
Update Update a dtc:monitor:http object

Updates a specific dtc:monitor:http object by reference

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param reference Reference of the dtc:monitor:http object
	@return DtcMonitorHttpAPIUpdateRequest
*/
func (a *DtcMonitorHttpAPIService) Update(ctx context.Context, reference string) DtcMonitorHttpAPIUpdateRequest {
	return DtcMonitorHttpAPIUpdateRequest{
		ApiService: a,
		ctx:        ctx,
		reference:  reference,
	}
}

// Execute executes the request
//
//	@return UpdateDtcMonitorHttpResponse
func (a *DtcMonitorHttpAPIService) UpdateExecute(r DtcMonitorHttpAPIUpdateRequest) (*UpdateDtcMonitorHttpResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *UpdateDtcMonitorHttpResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "DtcMonitorHttpAPIService.Update")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/dtc:monitor:http/{reference}"
	localVarPath = strings.Replace(localVarPath, "{"+"reference"+"}", url.PathEscape(internal.ParameterValueToString(r.reference, "reference")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.dtcMonitorHttp == nil {
		return localVarReturnValue, nil, internal.ReportError("dtcMonitorHttp is required and must be specified")
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
	if len(a.Client.Cfg.DefaultExtAttrs) > 0 && r.dtcMonitorHttp != nil {
		if r.dtcMonitorHttp.ExtAttrs == nil {
			r.dtcMonitorHttp.ExtAttrs = &map[string]ExtAttrs{}
		}
		for k, v := range a.Client.Cfg.DefaultExtAttrs {
			if _, ok := (*r.dtcMonitorHttp.ExtAttrs)[k]; !ok {
				(*r.dtcMonitorHttp.ExtAttrs)[k] = ExtAttrs{
					Value: v.Value,
				}
			}
		}
	}
	// body params
	localVarPostBody = r.dtcMonitorHttp
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
