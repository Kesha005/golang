package main

import (
	"fmt"
)

func main() {
	var manager User
	var workers []User
	manager.Name = "Kesha"
	manager.Surname = "Saparow"

	workers = append(workers, User{"Somename","SomeSurname"})
	workers = append(workers, User{"Somename2","SomeSurname2"})
	office := OfficceConstructor(manager,workers)
	//Get office manager name
	fmt.Println(office.Manager.Name)
}

type User struct {
	Name    string
	Surname string
}

type Officce struct {
	Manager User
	Workers []User
}

func OfficceConstructor(manager User, workers []User) *Officce {
	return &Officce{
		Manager: manager,
		Workers: workers,
	}
}
