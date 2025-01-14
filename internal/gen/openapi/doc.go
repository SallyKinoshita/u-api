// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package openapi

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

const (
	BearerAuthScopes = "BearerAuth.Scopes"
)

// Defines values for InvoiceResponseStatus.
const (
	Completed  InvoiceResponseStatus = "completed"
	Error      InvoiceResponseStatus = "error"
	Pending    InvoiceResponseStatus = "pending"
	Processing InvoiceResponseStatus = "processing"
)

// ErrorResponse defines model for ErrorResponse.
type ErrorResponse struct {
	// Code エラーコード
	Code *string `json:"code,omitempty"`

	// Message エラーメッセージ
	Message *string `json:"message,omitempty"`
}

// InvoiceListResponse defines model for InvoiceListResponse.
type InvoiceListResponse struct {
	Invoices *[]InvoiceResponse `json:"invoices,omitempty"`

	// Page 現在のページ番号
	Page *int `json:"page,omitempty"`

	// PerPage 1ページあたりの表示件数
	PerPage *int `json:"per_page,omitempty"`

	// Total 総件数
	Total *int `json:"total,omitempty"`
}

// InvoiceRequest defines model for InvoiceRequest.
type InvoiceRequest struct {
	// IssueDate 発行日（指定がない場合は現在日付）
	IssueDate *openapi_types.Date `json:"issue_date,omitempty"`

	// PartnerId 取引先ID
	PartnerId string `json:"partner_id"`

	// PaymentAmount 支払金額
	PaymentAmount int `json:"payment_amount"`

	// PaymentDueDate 支払期日
	PaymentDueDate openapi_types.Date `json:"payment_due_date"`
}

// InvoiceResponse defines model for InvoiceResponse.
type InvoiceResponse struct {
	// CompanyId 企業ID
	CompanyId *string `json:"company_id,omitempty"`

	// CreatedAt 作成日時
	CreatedAt *time.Time `json:"created_at,omitempty"`

	// Fee 手数料
	Fee *int `json:"fee,omitempty"`

	// FeeRate 手数料率
	FeeRate *float32 `json:"fee_rate,omitempty"`

	// Id 請求書ID
	Id *string `json:"id,omitempty"`

	// IssueDate 発行日
	IssueDate *openapi_types.Date `json:"issue_date,omitempty"`

	// PartnerId 取引先ID
	PartnerId *string `json:"partner_id,omitempty"`

	// PaymentAmount 支払金額
	PaymentAmount *int `json:"payment_amount,omitempty"`

	// PaymentDueDate 支払期日
	PaymentDueDate *openapi_types.Date `json:"payment_due_date,omitempty"`

	// Status ステータス
	Status *InvoiceResponseStatus `json:"status,omitempty"`

	// Tax 消費税
	Tax *int `json:"tax,omitempty"`

	// TaxRate 消費税率
	TaxRate *float32 `json:"tax_rate,omitempty"`

	// TotalAmount 請求金額
	TotalAmount *int `json:"total_amount,omitempty"`

	// UpdatedAt 更新日時
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// InvoiceResponseStatus ステータス
type InvoiceResponseStatus string

// GetApiInvoicesParams defines parameters for GetApiInvoices.
type GetApiInvoicesParams struct {
	// StartDate 検索開始日 (YYYY-MM-DD)
	StartDate openapi_types.Date `form:"start_date" json:"start_date"`

	// EndDate 検索終了日 (YYYY-MM-DD)
	EndDate openapi_types.Date `form:"end_date" json:"end_date"`

	// Page ページ番号
	Page *int `form:"page,omitempty" json:"page,omitempty"`

	// PerPage 1ページあたりの表示件数
	PerPage *int `form:"per_page,omitempty" json:"per_page,omitempty"`
}

// PostApiInvoicesJSONRequestBody defines body for PostApiInvoices for application/json ContentType.
type PostApiInvoicesJSONRequestBody = InvoiceRequest

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// GetApiInvoices request
	GetApiInvoices(ctx context.Context, params *GetApiInvoicesParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiInvoicesWithBody request with any body
	PostApiInvoicesWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiInvoices(ctx context.Context, body PostApiInvoicesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GetApiInvoices(ctx context.Context, params *GetApiInvoicesParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiInvoicesRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostApiInvoicesWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiInvoicesRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostApiInvoices(ctx context.Context, body PostApiInvoicesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiInvoicesRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetApiInvoicesRequest generates requests for GetApiInvoices
func NewGetApiInvoicesRequest(server string, params *GetApiInvoicesParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/invoices")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "start_date", runtime.ParamLocationQuery, params.StartDate); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "end_date", runtime.ParamLocationQuery, params.EndDate); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

		if params.Page != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "page", runtime.ParamLocationQuery, *params.Page); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		if params.PerPage != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "per_page", runtime.ParamLocationQuery, *params.PerPage); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		queryURL.RawQuery = queryValues.Encode()
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewPostApiInvoicesRequest calls the generic PostApiInvoices builder with application/json body
func NewPostApiInvoicesRequest(server string, body PostApiInvoicesJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostApiInvoicesRequestWithBody(server, "application/json", bodyReader)
}

// NewPostApiInvoicesRequestWithBody generates requests for PostApiInvoices with any type of body
func NewPostApiInvoicesRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/invoices")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// GetApiInvoicesWithResponse request
	GetApiInvoicesWithResponse(ctx context.Context, params *GetApiInvoicesParams, reqEditors ...RequestEditorFn) (*GetApiInvoicesResponse, error)

	// PostApiInvoicesWithBodyWithResponse request with any body
	PostApiInvoicesWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiInvoicesResponse, error)

	PostApiInvoicesWithResponse(ctx context.Context, body PostApiInvoicesJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiInvoicesResponse, error)
}

type GetApiInvoicesResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *InvoiceListResponse
	JSON400      *ErrorResponse
}

// Status returns HTTPResponse.Status
func (r GetApiInvoicesResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetApiInvoicesResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type PostApiInvoicesResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON201      *struct {
		// Message 成功メッセージ
		Message *string `json:"message,omitempty"`
	}
	JSON400 *ErrorResponse
}

// Status returns HTTPResponse.Status
func (r PostApiInvoicesResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PostApiInvoicesResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetApiInvoicesWithResponse request returning *GetApiInvoicesResponse
func (c *ClientWithResponses) GetApiInvoicesWithResponse(ctx context.Context, params *GetApiInvoicesParams, reqEditors ...RequestEditorFn) (*GetApiInvoicesResponse, error) {
	rsp, err := c.GetApiInvoices(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiInvoicesResponse(rsp)
}

// PostApiInvoicesWithBodyWithResponse request with arbitrary body returning *PostApiInvoicesResponse
func (c *ClientWithResponses) PostApiInvoicesWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiInvoicesResponse, error) {
	rsp, err := c.PostApiInvoicesWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiInvoicesResponse(rsp)
}

func (c *ClientWithResponses) PostApiInvoicesWithResponse(ctx context.Context, body PostApiInvoicesJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiInvoicesResponse, error) {
	rsp, err := c.PostApiInvoices(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiInvoicesResponse(rsp)
}

// ParseGetApiInvoicesResponse parses an HTTP response from a GetApiInvoicesWithResponse call
func ParseGetApiInvoicesResponse(rsp *http.Response) (*GetApiInvoicesResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiInvoicesResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest InvoiceListResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	}

	return response, nil
}

// ParsePostApiInvoicesResponse parses an HTTP response from a PostApiInvoicesWithResponse call
func ParsePostApiInvoicesResponse(rsp *http.Response) (*PostApiInvoicesResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostApiInvoicesResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 201:
		var dest struct {
			// Message 成功メッセージ
			Message *string `json:"message,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON201 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	}

	return response, nil
}

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// 請求書データ一覧の取得
	// (GET /api/invoices)
	GetApiInvoices(ctx echo.Context, params GetApiInvoicesParams) error
	// 新規請求書データの作成
	// (POST /api/invoices)
	PostApiInvoices(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetApiInvoices converts echo context to params.
func (w *ServerInterfaceWrapper) GetApiInvoices(ctx echo.Context) error {
	var err error

	ctx.Set(BearerAuthScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetApiInvoicesParams
	// ------------- Required query parameter "start_date" -------------

	err = runtime.BindQueryParameter("form", true, true, "start_date", ctx.QueryParams(), &params.StartDate)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter start_date: %s", err))
	}

	// ------------- Required query parameter "end_date" -------------

	err = runtime.BindQueryParameter("form", true, true, "end_date", ctx.QueryParams(), &params.EndDate)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter end_date: %s", err))
	}

	// ------------- Optional query parameter "page" -------------

	err = runtime.BindQueryParameter("form", true, false, "page", ctx.QueryParams(), &params.Page)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter page: %s", err))
	}

	// ------------- Optional query parameter "per_page" -------------

	err = runtime.BindQueryParameter("form", true, false, "per_page", ctx.QueryParams(), &params.PerPage)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter per_page: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetApiInvoices(ctx, params)
	return err
}

// PostApiInvoices converts echo context to params.
func (w *ServerInterfaceWrapper) PostApiInvoices(ctx echo.Context) error {
	var err error

	ctx.Set(BearerAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostApiInvoices(ctx)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/api/invoices", wrapper.GetApiInvoices)
	router.POST(baseURL+"/api/invoices", wrapper.PostApiInvoices)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"EQAA",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
