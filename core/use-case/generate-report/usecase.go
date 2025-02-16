package generate_report

import (
	"app/core/use-case/dto"
)

type GenerateReportUsecase interface {
	Execute(dto dto.GenerateReportDto) (string, error)
}

type generateReportUsecase struct {
}

func New() *generateReportUsecase {
	return new(generateReportUsecase)
}
