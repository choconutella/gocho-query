package query

import (
	"database/sql"
	"fmt"
)

type Recordset struct {
	Rows *sql.Rows
}

func NewRecordset(rows *sql.Rows) *Recordset {
	return &Recordset{
		Rows: rows,
	}
}

func (r *Recordset) Map() ([][]interface{}, error) {
	columns, err := r.Rows.Columns()
	if err != nil {
		return nil, fmt.Errorf("failed to get column names: %v", err)
	}

	rowData := make([]interface{}, len(columns))
	rowPointers := make([]interface{}, len(columns))
	for i := range rowData {
		rowPointers[i] = &rowData[i]
	}

	var dataFrame [][]interface{}
	rowsCount := 0
	for r.Rows.Next() {
		if err := r.Rows.Scan(rowPointers...); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}

		dataFrame = append(dataFrame, rowData)
		for i, data := range rowData {
			if data != nil {
				dataFrame[rowsCount][i] = data
			}
		}
		rowsCount++
	}

	return dataFrame, nil
}
