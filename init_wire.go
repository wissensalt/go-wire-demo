//go:build wireinject
// +build wireinject

package main

import "github.com/google/wire"

func InitializeRepository() EmployeeRepository {
	wire.Build(ProvideEmployeeRepository)

	return nil
}

func InitializeService(repository EmployeeRepository) EmployeeService {
	wire.Build(ProvideEmployeeService)

	return nil
}
