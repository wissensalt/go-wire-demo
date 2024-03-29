package main

import "fmt"

type (
	Employee struct {
		Id   int
		Name string
	}

	EmployeeRepository interface {
		Save(employee Employee)
	}

	EmployeeRepositoryImpl struct {
	}
)

func (e EmployeeRepositoryImpl) Save(employee Employee) {
	fmt.Printf("Saving Employee %v\n", employee)
}

func ProvideEmployeeRepository() EmployeeRepository {
	return EmployeeRepositoryImpl{}
}
