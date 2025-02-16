package controller_test

import (
	"app/core/controller"
	"app/core/use-case/dto"
	"app/mocks/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io"
	"os"
	"testing"
	"time"
)

func TestGenerateReport(t *testing.T) {
	mockUsecase := new(mocks.GenerateReportUsecase)

	startDate := time.Date(2025, 02, 15, 9, 0, 0, 0, time.UTC)
	endDate := time.Date(2025, 02, 15, 12, 0, 0, 0, time.UTC)
	intervalMinutes := 60
	initialBalance := 100000.0

	mockUsecase.EXPECT().Execute(mock.AnythingOfType("dto.GenerateReportDto")).Run(func(dto dto.GenerateReportDto) {
		assert.Equal(t, startDate, dto.GetStartDate())
		assert.Equal(t, endDate, dto.GetEndDate())
		assert.Equal(t, intervalMinutes, dto.GetMinutesInterval())
		assert.Equal(t, initialBalance, dto.GetInitialBalance())
	})

	controller := controller.New(mockUsecase)

	tradesFile, err := os.Open("../../arquivos-exemplo/march_2021_trades.csv")
	if err != nil {
		t.Fatalf("Error loading trades: %v", err)
	}
	defer tradesFile.Close()

	assetsFiles := make(map[string]io.Reader)

	assetAFile, err := os.Open("../../arquivos-exemplo/march_2021_pricesA.csv")
	if err != nil {
		t.Fatalf("Error loading asset A file: %v", err)
	}
	defer assetAFile.Close()
	assetsFiles["A"] = assetAFile

	assetBFile, err := os.Open("../../arquivos-exemplo/march_2021_pricesB.csv")
	if err != nil {
		t.Fatalf("Error loading asset B file: %v", err)
	}
	defer assetBFile.Close()
	assetsFiles["B"] = assetBFile

	controller.GenerateReport(startDate, endDate, intervalMinutes, initialBalance, tradesFile, assetsFiles)

	mockUsecase.AssertExpectations(t)
}
