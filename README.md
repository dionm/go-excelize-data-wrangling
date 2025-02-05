# Excel Staff Hierarchy Tool

A Go CLI tool to process an Excel file and retrieve employees reporting to a specific manager.

## Usage

### List column names:
```sh
go run main.go list-columns data/company.xlsx

### Get reportees under a given manager
```sh
go run main.go reportees data/company.xlsx 1

## Project Structure:

excel_filter/
│── main.go            # CLI Entry Point
│── utils/
│   ├── fileutils.go   # Excel file handling
│   ├── filter.go      # Hierarchical filtering (BFS)
│── data/
│   ├── company.xlsx   # Sample input
│   ├── output.xlsx    # Filtered output
