package main

import (
	"errors"
	"fmt"
	"strings"
)

const (
	HR      = "HR"
	IT      = "IT"
	Finance = "Finance"
)

type Employee struct {
	ID         int
	Name       string
	Age        int
	Department string
}

var employees []Employee

func AddEmployee(id int, name string, age int, department string) error {
	for _, emp := range employees {
		if emp.ID == id {
			return errors.New("ID must be unique")
		}
	}

	if age <= 18 {
		return errors.New("Age must be greater than 18")
	}

	employees = append(employees, Employee{
		ID:         id,
		Name:       name,
		Age:        age,
		Department: department,
	})
	return nil
}

func SearchEmployee(query string) (*Employee, error) {
	for _, emp := range employees {
		if strings.EqualFold(fmt.Sprint(emp.ID), query) || strings.EqualFold(emp.Name, query) {
			return &emp, nil
		}
	}
	return nil, errors.New("Employee not found")
}

func ListEmployeesByDepartment(department string) []Employee {
	var filteredEmployees []Employee
	for _, emp := range employees {
		if strings.EqualFold(emp.Department, department) {
			filteredEmployees = append(filteredEmployees, emp)
		}
	}
	return filteredEmployees
}

func CountEmployeesByDepartment(department string) int {
	count := 0
	for _, emp := range employees {
		if strings.EqualFold(emp.Department, department) {
			count++
		}
	}
	return count
}

func main() {
	if err := AddEmployee(1, "Vishal", 25, IT); err != nil {
		fmt.Println("Error:", err)
	}
	if err := AddEmployee(2, "Nishant", 30, HR); err != nil {
		fmt.Println("Error:", err)
	}
	if err := AddEmployee(3, "Aryan", 22, IT); err != nil {
		fmt.Println("Error:", err)
	}

	emp, err := SearchEmployee("Vishal")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Employee found:", *emp)
	}

	fmt.Println("Employees in IT department:")
	for _, emp := range ListEmployeesByDepartment(IT) {
		fmt.Println(emp)
	}

	fmt.Printf("Number of employees in HR department: %d\n", CountEmployeesByDepartment(HR))
}
