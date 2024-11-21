package employee

type Employee struct {
	name       string
	department string
	salary     float64
}

func CreateEmployee(name string, department string, salary float64) *Employee {
	return &Employee{
		name:       name,
		department: department,
		salary:     salary,
	}
}
func SalaryAggregator(filter func(Employee) bool) func(employees []Employee) float64 {
	salary := 0.0
	count := 0
	return func(employees []Employee) float64 {
		for _, emp := range employees {
			if filter(emp) {
				salary = salary + emp.salary
				count++
			}
		}
		return salary / float64(count)
	}
}
func FilterByDepartment(employee Employee) bool {
	if employee.department == "Sales" {
		return true
	}
	return false
}
func FilterBySalary(employee Employee) bool {
	if employee.salary > 5000 {
		return true
	}
	return false
}
