package cmd

import (
	"context"
	"landmarks/pkg/handler"
	"landmarks/pkg/openapi"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	oapiMiddleware "github.com/oapi-codegen/echo-middleware"
)

func APIMain(address string) {
	h := handler.Handler{}

	e := echo.New()
	e.Use(middleware.RequestID())
	if err := h.Initialize(); err != nil {
		slog.Error("初期化に失敗しました", slog.String("error", err.Error()))
	}
	defer h.Finalize()
	openapi.RegisterHandlers(e, &h)

	if swagger, err := openapi.GetSwagger(); err == nil {
		swagger.Servers = nil // <-これがないと、OpenAPI定義のServersに含まれた形のlisten設定じゃないとエラーになる
		e.Use(oapiMiddleware.OapiRequestValidatorWithOptions(swagger, &oapiMiddleware.Options{ErrorHandler: h.ErrorHandler}))
	} else {
		slog.Warn("OpenAPI仕様が見つからなかったので、リクエストの妥当性確認を行いません.", slog.String("error", err.Error()))
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go func() {
		if err := e.Start(address); err != nil && err != http.ErrServerClosed {
			slog.Error("APIサービスを終了します.")
		}
	}()

	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		slog.Error("APIサービスの終了時にエラーが発生しました.", slog.String("error", err.Error()))
	}
}
