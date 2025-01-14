package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/SallyKinoshita/u-api/internal/application/usecase"
	"github.com/SallyKinoshita/u-api/internal/gen/openapi"
	"github.com/SallyKinoshita/u-api/internal/interface/presenter"
	"github.com/labstack/echo/v4"
)

type Invoice struct {
	InvoiceUseCase usecase.Invoice
}

func (i *Invoice) GetApiInvoices(ctx echo.Context, params openapi.GetApiInvoicesParams) error {
	page := 1
	if params.Page != nil {
		page = *params.Page
	}

	perPage := 10
	if params.PerPage != nil {
		perPage = *params.PerPage
	}

	startDate := time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC)
	if !params.StartDate.IsZero() {
		startDate = params.StartDate.Time
	}

	endDate := time.Date(9999, 12, 31, 23, 59, 59, 0, time.UTC)
	if !params.EndDate.IsZero() {
		endDate = params.EndDate.Time
	}

	invoices, total, err := i.InvoiceUseCase.List(ctx.Request().Context(), startDate, endDate, page, perPage)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	response := presenter.ToInvoiceListResponse(invoices, total, page, perPage)

	return ctx.JSON(http.StatusOK, response)
}

func (i *Invoice) PostApiInvoices(ctx echo.Context) error {
	var body openapi.PostApiInvoicesJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body").SetInternal(err)
	}

	_, err := i.InvoiceUseCase.Create(ctx.Request().Context(), strToInt(body.PartnerId), intToF64(body.PaymentAmount), body.IssueDate.Time, body.PaymentDueDate.Time)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create invoice").SetInternal(err)
	}

	// 成功メッセージを返す
	response := map[string]string{
		"message": "請求書が正常に作成されました。",
	}

	return ctx.JSON(http.StatusCreated, response)
}

func strToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return i
}

func intToF64(i int) float64 {
	return float64(i)
}
