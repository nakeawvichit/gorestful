package main

var id int

var emps Employees

func init() {
	createEmployee(Employee{Name: "Duckna", Age: 25})
}

func createEmployee(emp Employee) Employees {
	id++
	emp.Id = id
	emps = append(emps, emp)
	return emps
}
