// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package openapi

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// GetLandmarks 登録されている登録地点のインデックス一覧を返します.
	// クエリパラメータnameを使って地点名称によるフィルタをかけることができます。(部分一致)
	// (GET /landmarks)
	GetLandmarks(ctx echo.Context, params GetLandmarksParams) error
	// PostLandmarks 新しい地点を登録します.
	// (POST /landmarks)
	PostLandmarks(ctx echo.Context) error
	// PostLandmarksSearch 検索条件を指定して検索します.
	// (POST /landmarks/_search)
	PostLandmarksSearch(ctx echo.Context) error
	// DeleteLandmarksID 指定された地点データを削除します.
	// (DELETE /landmarks/{id})
	DeleteLandmarksID(ctx echo.Context, id int) error
	// GetLandmarksID 指定された地点データを取得します.
	// (GET /landmarks/{id})
	GetLandmarksID(ctx echo.Context, id int) error
	// PatchLandmarksID 指定地点部分更新 指定された地点データを更新します.
	// (PATCH /landmarks/{id})
	PatchLandmarksID(ctx echo.Context, id int) error
	// PutLandmarksID 指定地点更新 指定された地点データを更新します.
	// (PUT /landmarks/{id})
	PutLandmarksID(ctx echo.Context, id int) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetLandmarks converts echo context to params.
func (w *ServerInterfaceWrapper) GetLandmarks(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetLandmarksParams
	// ------------- Optional query parameter "name" -------------

	err = runtime.BindQueryParameter("form", true, false, "name", ctx.QueryParams(), &params.Name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter name: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetLandmarks(ctx, params)
	return err
}

// PostLandmarks converts echo context to params.
func (w *ServerInterfaceWrapper) PostLandmarks(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostLandmarks(ctx)
	return err
}

// PostLandmarksSearch converts echo context to params.
func (w *ServerInterfaceWrapper) PostLandmarksSearch(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostLandmarksSearch(ctx)
	return err
}

// DeleteLandmarksID converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteLandmarksID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithOptions("simple", "id", ctx.Param("id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DeleteLandmarksID(ctx, id)
	return err
}

// GetLandmarksID converts echo context to params.
func (w *ServerInterfaceWrapper) GetLandmarksID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithOptions("simple", "id", ctx.Param("id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetLandmarksID(ctx, id)
	return err
}

// PatchLandmarksID converts echo context to params.
func (w *ServerInterfaceWrapper) PatchLandmarksID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithOptions("simple", "id", ctx.Param("id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PatchLandmarksID(ctx, id)
	return err
}

// PutLandmarksID converts echo context to params.
func (w *ServerInterfaceWrapper) PutLandmarksID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithOptions("simple", "id", ctx.Param("id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PutLandmarksID(ctx, id)
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

	router.GET(baseURL+"/landmarks", wrapper.GetLandmarks)
	router.POST(baseURL+"/landmarks", wrapper.PostLandmarks)
	router.POST(baseURL+"/landmarks/_search", wrapper.PostLandmarksSearch)
	router.DELETE(baseURL+"/landmarks/:id", wrapper.DeleteLandmarksID)
	router.GET(baseURL+"/landmarks/:id", wrapper.GetLandmarksID)
	router.PATCH(baseURL+"/landmarks/:id", wrapper.PatchLandmarksID)
	router.PUT(baseURL+"/landmarks/:id", wrapper.PutLandmarksID)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
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
