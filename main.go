package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"excel_filter/utils"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage:")
		fmt.Println("  List Columns: go run main.go list-columns <input_file.xlsx>")
		fmt.Println("  Get Reportees: go run main.go reportees <input_file.xlsx> <manager_id>")
		return
	}

	command := os.Args[1]
	inputFile := os.Args[2]

	switch command {
	case "list-columns":
		err := utils.PrintColumnNames(inputFile)
		if err != nil {
			log.Fatalf("Error listing columns: %v", err)
		}
	case "reportees":
		if len(os.Args) < 4 {
			fmt.Println("Usage: go run main.go reportees <input_file.xlsx> <manager_id>")
			return
		}

		managerID, err := strconv.Atoi(os.Args[3])
		if err != nil {
			log.Fatalf("Invalid Manager ID: %v", err)
		}

		employees, err := utils.ReadEmployeesFromExcel(inputFile)
		if err != nil {
			log.Fatalf("Failed to read input file: %v", err)
		}

		employeeMap := utils.BuildEmployeeMap(employees)
		reportees := utils.GetReporteesBFS(employeeMap, managerID)

		fmt.Printf("Reportees under Manager ID %d:\n", managerID)
		for _, emp := range reportees {
			fmt.Printf("- %s (ID: %d)\n", emp.Name, emp.ID)
		}

		outputFile := "data/output.xlsx"
		err = utils.ExportReporteesToExcel(reportees, outputFile)
		if err != nil {
			log.Fatalf("Failed to write output file: %v", err)
		}
		fmt.Println("Filtered report saved to:", outputFile)

	default:
		fmt.Println("Unknown command. Use 'list-columns' or 'reportees'.")
	}
