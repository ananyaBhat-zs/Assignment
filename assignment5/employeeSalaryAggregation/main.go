package main

import (
	"employee.com/employee"
	"fmt"
)

func main() {
	n := 0
	fmt.Println("Enter the no employees")
	fmt.Scan(&n)
	emp := make([]employee.Employee, n)
	for i := 0; i < n; i++ {
		var name, dep string
		var salary float64
		fmt.Println("Enter name ,department and salary of employee ", i)
		fmt.Scan(&name, &dep, &salary)
		emp[i] = *employee.CreateEmployee(name, dep, salary)
	}
	avgSalary := employee.SalaryAggregator(employee.FilterBySalary)
	fmt.Println("The avg salary for employees with salary greater than $5000 is", avgSalary(emp))
	avgSalaryD := employee.SalaryAggregator(employee.FilterByDepartment)
	fmt.Println("The avg salary for employees from Sales department is", avgSalaryD(emp))

}
