package generate_report

import (
	"app/core/use-case/dto"
)

type GenerateReportUsecase interface {
	Execute(input dto.GenerateReportDto) (dto.ReportDto, error)
}

type generateReportUsecase struct {
}

func New() *generateReportUsecase {
	return new(generateReportUsecase)
}
