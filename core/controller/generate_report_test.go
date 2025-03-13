package controller_test

import (
	"app/core/controller"
	"app/core/use-case/dto"
	mocks_core "app/mocks/core"
	mocks_infra "app/mocks/infra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestGenerateReport(t *testing.T) {
	mockUsecase := new(mocks_core.GenerateReportUsecase)
	mockFileHandler := new(mocks_infra.FilesHandler)

	startDate := time.Date(2025, 02, 15, 9, 0, 0, 0, time.UTC)
	endDate := time.Date(2025, 02, 15, 12, 0, 0, 0, time.UTC)
	intervalMinutes := 60
	initialBalance := 100000.0

	mockTrades := dto.TradeDtos{{}}
	mockPrices := dto.PricesDto{
		time.Date(2025, 02, 15, 10, 0, 0, 0, time.UTC): {
			"A": 100.5,
			"B": 200.7,
		},
	}

	mockFileHandler.On("LoadValuesInInterval", startDate, endDate).
		Return(mockTrades, mockPrices, nil)

	mockUsecase.EXPECT().Execute(mock.AnythingOfType("dto.GenerateReportDto")).
		Run(func(dto dto.GenerateReportDto) {
			assert.Equal(t, startDate, dto.GetStartDate())
			assert.Equal(t, endDate, dto.GetEndDate())
			assert.Equal(t, intervalMinutes, dto.GetMinutesInterval())
			assert.Equal(t, initialBalance, dto.GetInitialBalance())
		}).
		Return(dto.ReportDto{})

	mockFileHandler.On("BuildByteArray", mock.Anything).Return([]byte("mocked report"), nil)

	rController := controller.New(mockUsecase)
	rController.SetFilesHandler(mockFileHandler)

	reportBytes, err := rController.GenerateReport(startDate, endDate, intervalMinutes, initialBalance)

	assert.NoError(t, err)
	assert.NotEmpty(t, reportBytes)
	assert.Equal(t, "mocked report", string(reportBytes))

	mockUsecase.AssertExpectations(t)
	mockFileHandler.AssertExpectations(t)
}
