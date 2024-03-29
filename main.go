package main

func main() {
	employeeService := InitializeService(InitializeRepository())
	employeeService.CreateEmployee(Employee{
		Id:   1,
		Name: "John Doe",
	})
}
