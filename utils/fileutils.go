package utils

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"strconv"
	"errors"
)

type Employee struct {
	ID      int
	Name    string
	Manager int
}

func ReadEmployeesFromExcel(filename string) ([]Employee, error) {
	f, err := excelize.OpenFile(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	sheetName := f.GetSheetName(0)
	rows, err := f.GetRows(sheetName)
	if err != nil {
		return nil, err
	}

	if len(rows) < 2 {
		return nil, errors.New("invalid data format")
	}

	var employees []Employee
	for i, row := range rows[1:] {
		if len(row) < 3 {
			continue
		}

		id, err := strconv.Atoi(row[0])
		managerID, err2 := strconv.Atoi(row[2])

		if err != nil || err2 != nil {
			fmt.Printf("Skipping invalid row %d\n", i+2)
			continue
		}

		employees = append(employees, Employee{
			ID:      id,
			Name:    row[1],
			Manager: managerID,
		})
	}

	return employees, nil
}

func ExportReporteesToExcel(reportees []Employee, filename string) error {
	f := excelize.NewFile()
	sheetName := "Reportees"
	f.SetSheetName("Sheet1", sheetName)

	f.SetCellValue(sheetName, "A1", "ID")
	f.SetCellValue(sheetName, "B1", "Name")
	f.SetCellValue(sheetName, "C1", "Manager")

	for i, emp := range reportees {
		f.SetCellValue(sheetName, fmt.Sprintf("A%d", i+2), emp.ID)
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", i+2), emp.Name)
		f.SetCellValue(sheetName, fmt.Sprintf("C%d", i+2), emp.Manager)
	}

	return f.SaveAs(filename)
}

func PrintColumnNames(filename string) error {
	f, err := excelize.OpenFile(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	sheetName := f.GetSheetName(0)
	rows, err := f.GetRows(sheetName)
	if err != nil {
		return err
	}

	if len(rows) == 0 {
		return errors.New("empty sheet")
	}

	headerRow := rows[0]

	fmt.Println("Column Names and Indexes:")
	for i, colName := range headerRow {
		fmt.Printf("Index: %d → Column Name: %s\n", i, colName)
	}

	return nil
}
