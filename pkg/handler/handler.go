package handler

import (
	"errors"
	"fmt"
	"landmarks/pkg/database"
	"landmarks/pkg/openapi"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Handler struct {
	db *gorm.DB
}

// GetLandmarks 登録地点の一覧を返します.
// (GET /landmarks)
func (h *Handler) GetLandmarks(ctx echo.Context, params openapi.GetLandmarksParams) error {
	requestID := ctx.Response().Header().Get(echo.HeaderXRequestID)
	h.Delay()
	if response, err := database.GetLandmarks(h.db, params.Page, params.PerPage); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, openapi.InternalServerError{
			RequestID: requestID,
			MessageID: "GetLandmarks.500",
			Arguments: map[string]interface{}{"error": err.Error(), "perPage": params.PerPage, "page": params.Page},
		})
	} else {
		return ctx.JSON(http.StatusOK, response)
	}
}

// PostLandmarks 新しい地点を登録します.
// (POST /landmarks)
func (h *Handler) PostLandmarks(ctx echo.Context) error {
	requestID := ctx.Response().Header().Get(echo.HeaderXRequestID)
	h.Delay()
	var requestBody openapi.PostLandmarksJSONRequestBody
	if err := ctx.Bind(&requestBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, openapi.BadRequest{
			RequestID: requestID,
			MessageID: "GetLandmarks.400",
			Arguments: map[string]interface{}{"error": err.Error()},
		})
	}

	if id, err := database.CreateLandmark(h.db, requestBody); err != nil {
		fmt.Print(err.Error())
		slog.Warn("新しい地点情報の生成に失敗しました.", slog.String("error", err.Error()))
		return echo.NewHTTPError(http.StatusInternalServerError, openapi.InternalServerError{
			RequestID: requestID,
			MessageID: "PostLandmarks.500",
			Arguments: map[string]interface{}{"error": err.Error()},
		})
	} else {
		return ctx.JSON(http.StatusCreated, openapi.Created{
			RequestID: requestID,
			MessageID: "PostLandmarks.500",
			CreatedID: strconv.Itoa(id),
		})
	}
}

// PostLandmarksSearch 検索条件を指定して検索します.
// (POST /landmarks/_search)
func (h *Handler) PostLandmarksSearch(ctx echo.Context) error {
	requestID := ctx.Response().Header().Get(echo.HeaderXRequestID)
	h.Delay()
	var requestBody openapi.PostLandmarksSearchJSONRequestBody
	if err := ctx.Bind(&requestBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, openapi.BadRequest{
			RequestID: requestID,
			MessageID: "PostLandmarksSearch.400",
			Arguments: map[string]interface{}{"error": err.Error()},
		})
	}

	if landmarks, err := database.SearchLandmarks(h.db, requestBody); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, openapi.InternalServerError{
			RequestID: requestID,
			MessageID: "PostLandmarks.500",
			Arguments: map[string]interface{}{"error": err.Error()},
		})
	} else {
		return ctx.JSON(http.StatusOK, landmarks)
	}
}

// DeleteLandmarksID 指定された地点データを削除します.
// (DELETE /landmarks/{id})
func (h *Handler) DeleteLandmarksID(ctx echo.Context, id int) error {
	requestID := ctx.Response().Header().Get(echo.HeaderXRequestID)
	h.Delay()
	if err := database.DeleteLandmark(h.db, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, openapi.NotFound{
				RequestID:  requestID,
				MessageID:  "DeleteLandmarksId.404",
				Arguments:  map[string]interface{}{"error": err.Error()},
				NotFoundID: strconv.Itoa(id),
			})
		} else {
			return echo.NewHTTPError(http.StatusInternalServerError, openapi.InternalServerError{
				RequestID: requestID,
				MessageID: "DeleteLandmarksId.500",
				Arguments: map[string]interface{}{"error": err.Error(), "requiredId": id},
			})
		}
	} else {
		return ctx.JSON(http.StatusOK, openapi.Deleted{
			RequestID: requestID,
			MessageID: "DeleteLandmarksId.200",
			DeletedID: strconv.Itoa(id),
		})
	}
}

// PutLandmarksID 指定地点更新 指定された地点データを更新します.
// (PUT /landmarks/{id})
func (h *Handler) GetLandmarksID(ctx echo.Context, id int) error {
	requestID := ctx.Response().Header().Get(echo.HeaderXRequestID)
	h.Delay()
	if response, err := database.ReadLandmark(h.db, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, openapi.NotFound{
				RequestID:  requestID,
				MessageID:  "GetLandmarksID.404",
				Arguments:  map[string]interface{}{"error": err.Error()},
				NotFoundID: strconv.Itoa(id),
			})
		} else {
			return echo.NewHTTPError(http.StatusInternalServerError, openapi.InternalServerError{
				RequestID: requestID,
				MessageID: "GetLandmarksID.500",
				Arguments: map[string]interface{}{"error": err.Error(), "requiredId": id},
			})
		}
	} else {
		return ctx.JSON(http.StatusOK, response)
	}
}

// PutLandmarksID 指定地点更新
// (PUT /landmarks/{id})
func (h *Handler) PutLandmarksID(ctx echo.Context, id int) error {
	requestID := ctx.Response().Header().Get(echo.HeaderXRequestID)
	h.Delay()
	var requestBody openapi.PutLandmarksIDJSONRequestBody
	if err := ctx.Bind(&requestBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, openapi.BadRequest{
			RequestID: requestID,
			MessageID: "PutLandmarksID.400",
			Arguments: map[string]interface{}{"error": err.Error()},
		})
	}

	if err := database.UpdateLandmark(h.db, id, requestBody); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, openapi.NotFound{
				RequestID:  requestID,
				MessageID:  "PutLandmarksID.404",
				Arguments:  map[string]interface{}{"error": err.Error()},
				NotFoundID: strconv.Itoa(id),
			})
		} else {
			return echo.NewHTTPError(http.StatusInternalServerError, openapi.InternalServerError{
				RequestID: requestID,
				MessageID: "PutLandmarksID.500",
				Arguments: map[string]interface{}{"error": err.Error(), "requiredId": id},
			})
		}
	}
	return echo.NewHTTPError(http.StatusInternalServerError, openapi.InternalServerError{
		RequestID: requestID,
		MessageID: "GetLandmarksID.500",
		Arguments: map[string]interface{}{"error": nil, "requiredId": id},
	})
}

// PatchLandmarksID 指定地点更新 指定された地点データを更新します.
// (PATCH /landmarks/{id})
func (h *Handler) PatchLandmarksID(ctx echo.Context, id int) error {

	return ctx.JSON(http.StatusNotImplemented, nil)
}
