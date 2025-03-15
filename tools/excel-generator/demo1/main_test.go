package main

import (
	"fmt"
	"testing"

	"github.com/xuri/excelize/v2"
)

func TestExcelize(t *testing.T) {
	t.Run("CreateSpreadsheetain", func(t *testing.T) {
		CreateSpreadsheetain()
	})

}

func CreateSpreadsheetain() {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// Create a new sheet.
	index, err := f.NewSheet("Sheet2")
	if err != nil {
		fmt.Println(err)
		return
	}
	// Set value of a cell.
	f.SetCellValue("Sheet2", "A2", "Hello world.")
	f.SetCellValue("Sheet1", "B2", 100)
	// Set active sheet of the workbook.
	f.SetActiveSheet(index)
	// Save spreadsheet by the given path.
	if err := f.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}

	v, err := f.GetCellValue("Sheet1", "B2")
	fmt.Println(v)
}
