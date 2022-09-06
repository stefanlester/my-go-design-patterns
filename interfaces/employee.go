package main

import "fmt"

type Employee interface {
	GetName() string
}

type Engineer struct {
	Name string
}

type Manager struct {
	Name string
}

func (e *Engineer) GetName() string {
	return "Employee Name: " + e.Name
}

func (m *Manager) GetName() string {
	return "Manager Name: " + m.Name
}

func PrintDetails(e Employee) {
	fmt.Println(e.GetName())
}

func main() {
	engineer := &Engineer{Name: "Lester"} 
	manager := &Manager{Name: "Amoah"} 
	PrintDetails(engineer)
	PrintDetails(manager)
}
