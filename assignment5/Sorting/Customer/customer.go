package Customer

type Customer struct {
	name   string
	age    int
	salary int
}

func NewCustomer(name string, age int, salary int) *Customer {
	return &Customer{
		name:   name,
		age:    age,
		salary: salary,
	}
}
func SortTester(cust []Customer, attribute string, direction string) []string {
	comparator := findComparator(attribute, direction)
	names := selectionSort(cust, comparator)
	return names
}
func selectionSort(cust []Customer, comparator func(Customer, Customer) bool) []string {
	names := make([]string, len(cust))
	for i := 0; i < len(cust); i++ {
		minIndex := i
		for j := i + 1; j < len(cust); j++ {
			if comparator(cust[j], cust[minIndex]) {
				minIndex = j
			}
		}
		cust[minIndex], cust[i] = cust[i], cust[minIndex]
	}
	for i, _ := range cust {
		names[i] = cust[i].name
	}
	return names
}
func findComparator(attribute string, direction string) func(c1 Customer, c2 Customer) bool {
	switch attribute {
	case "name":
		{
			return func(c1 Customer, c2 Customer) bool {
				if direction == "Asc" {
					return c1.name < c2.name
				} else {
					return c1.name > c2.name
				}
			}
		}
	case "age":
		{
			return func(c1 Customer, c2 Customer) bool {
				if direction == "Asc" {
					return c1.age < c2.age
				} else {
					return c1.age > c2.age
				}
			}
		}
	case "salary":
		{
			return func(c1 Customer, c2 Customer) bool {
				if direction == "Asc" {
					return c1.salary < c2.salary
				} else {
					return c1.salary > c2.salary
				}
			}
		}
	}
	return func(c1 Customer, c2 Customer) bool {
		return false
	}
}
