package csv

import (
	"app/core/use-case/dto"
	"bytes"
	"encoding/csv"
	"fmt"
)

func (adp cvsHandler) BuildByteArray(report dto.ReportDto) ([]byte, error) {
	var buf bytes.Buffer
	writer := csv.NewWriter(&buf)

	err := writer.Write([]string{"timestamp", "Patrimônio Total", "Rentabilidade Acumulada"})
	if err != nil {
		return nil, err
	}

	for _, reportLine := range report {
		err := writer.Write([]string{
			reportLine.GetDate().Format("2006-01-02 15:04:05"),
			fmt.Sprintf("%.4f", reportLine.GetTotalBalance()),
			fmt.Sprintf("%.5f", reportLine.GetAccumulatedProfit()),
		})
		if err != nil {
			return nil, err
		}
	}
	writer.Flush()

	return buf.Bytes(), nil
}
