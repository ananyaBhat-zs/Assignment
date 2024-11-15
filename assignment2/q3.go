package main

import "fmt"

type car struct {
	name       string
	noOfWheels int
	brand      string
}

type bike struct {
	name       string
	noOfWheels int
	brand      string
}
type truck struct {
	name       string
	noOfWheels int
	brand      string
}

func q3() {
	c := car{"Swift", 4, "Maruti Suzuki"}
	b := bike{"KTM", 2, "Honda"}
	t := truck{"ahuja", 6, "tata"}
	findType(c)
	findType(b)
	findType(t)
}
func findType(v interface{}) {
	switch v.(type) {
	case car:
		{
			fmt.Println("car")
		}
	case bike:
		{
			fmt.Println("bike")
		}
	case truck:
		{
			fmt.Println("truck")
		}
	}
}
