package utils

func BuildEmployeeMap(employees []Employee) map[int][]Employee {
	employeeMap := make(map[int][]Employee)
	for _, emp := range employees {
		employeeMap[emp.Manager] = append(employeeMap[emp.Manager], emp)
	}
	return employeeMap
}

func GetReporteesBFS(employeeMap map[int][]Employee, managerID int) []Employee {
	var reportees []Employee
	queue := []Employee{}

	if employees, found := employeeMap[managerID]; found {
		queue = append(queue, employees...)
	}

	for len(queue) > 0 {
		emp := queue[0]
		queue = queue[1:]
		reportees = append(reportees, emp)

		if employees, found := employeeMap[emp.ID]; found {
			queue = append(queue, employees...)
		}
	}
	return reportees
}
