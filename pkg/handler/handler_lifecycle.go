package handler

import (
	"landmarks/pkg/openapi"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func (h *Handler) Initialize() (err error) {

	if h.db, err = gorm.Open(sqlite.Open("landmarks.sqlite"), &gorm.Config{}); err != nil {
		return err
	}
	if err = h.db.AutoMigrate(&openapi.Landmark{}); err != nil {
		return err
	}
	return nil
}

func (h *Handler) Finalize() {

}
func (h *Handler) ErrorHandler(ctx echo.Context, err *echo.HTTPError) error {
	requestID := ctx.Response().Header().Get(echo.HeaderXRequestID)
	requestURL := ctx.Request().URL.String()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, openapi.BadRequest{
			RequestID: requestID,
			MessageID: "ValidateRequest",
			Arguments: map[string]interface{}{"requestURL": requestURL, "code": err.Code, "message": err.Message.(string), "internal": err.Internal.Error()},
		})
	}
	return err
}
