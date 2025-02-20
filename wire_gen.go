// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

// Injectors from init_wire.go:

func InitializeRepository() EmployeeRepository {
	employeeRepository := ProvideEmployeeRepository()
	return employeeRepository
}

func InitializeService(repository EmployeeRepository) EmployeeService {
	employeeService := ProvideEmployeeService(repository)
	return employeeService
}
