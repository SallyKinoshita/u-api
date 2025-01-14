package main

import (
	"fmt"
	"log"
	"os"

	"github.com/SallyKinoshita/u-api/internal/application/usecase"
	"github.com/SallyKinoshita/u-api/internal/gen/openapi"
	"github.com/SallyKinoshita/u-api/internal/infrastructure/db"
	persistencerepository "github.com/SallyKinoshita/u-api/internal/infrastructure/persistence/repository"
	"github.com/SallyKinoshita/u-api/internal/interface/controller"
	"github.com/labstack/echo/v4"
)

func main() {
	// DB接続
	dbConn, err := db.NewBunDB()
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	// Echoインスタンス作成
	e := echo.New()

	// TODO: middlewareの設定
	// TODO: middlewareに認証周りを追加する

	// DI設定
	invoiceRepo := persistencerepository.NewInvoice()
	invoiceUsecase := usecase.NewInvoice(dbConn, invoiceRepo)
	invoiceController := &controller.Invoice{
		InvoiceUseCase: invoiceUsecase,
	}

	// OpenAPIのハンドラー登録
	openapi.RegisterHandlers(e, invoiceController)

	// サーバー起動
	port := os.Getenv("API_PORT")
	log.Printf("Starting server on %v", port)
	if err := e.StartTLS(fmt.Sprintf(":%s", port), "/app/server.crt", "/app/server.key"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
