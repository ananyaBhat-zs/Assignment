package main

import (
	"fmt"
	"sorting.com/Customer"
)

func main() {
	fmt.Println("Enter the no of customers")
	n := 0
	fmt.Scan(&n)
	cust := make([]Customer.Customer, n)
	for i := 0; i < n; i++ {
		var name string
		var age, salary int
		fmt.Println("Enter the name , age and salary of customer", i+1)
		fmt.Scan(&name, &age, &salary)
		cust[i] = *Customer.NewCustomer(name, age, salary)
	}
	namesBasedOnNameAsc := Customer.SortTester(cust, "name", "Asc")
	fmt.Println("Based on names,Asc", namesBasedOnNameAsc)
	namesBasedOnNameDesc := Customer.SortTester(cust, "name", "Desc")
	fmt.Println("Based on names,Desc", namesBasedOnNameDesc)
	namesBasedOnSalaryAsc := Customer.SortTester(cust, "salary", "Asc")
	fmt.Println("Based on Salary Asc", namesBasedOnSalaryAsc)
	nameesBasedOnSalaryDesc := Customer.SortTester(cust, "salary", "Desc")
	fmt.Println("Based on Salary Desc", nameesBasedOnSalaryDesc)
	namesBasedOnAgeAsc := Customer.SortTester(cust, "age", "Asc")
	fmt.Println("Based on age,asc", namesBasedOnAgeAsc)
	namesBasedOnAgeDesc := Customer.SortTester(cust, "age", "Desc")
	fmt.Println("Based on age , desc", namesBasedOnAgeDesc)

}
