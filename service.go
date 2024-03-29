package main

type (
	EmployeeService interface {
		CreateEmployee(employee Employee)
	}

	EmployeeServiceImpl struct {
		EmployeeRepository
	}
)

func (e EmployeeServiceImpl) CreateEmployee(employee Employee) {
	e.EmployeeRepository.Save(employee)
}

func ProvideEmployeeService(repository EmployeeRepository) EmployeeService {
	return EmployeeServiceImpl{repository}
}
