package main

import (
	//"encoding/json"
	"fmt"
	"sort"
)

func main() {
	task_1()
	task_2()
	task_3()
}

type Card1 struct {
	WorkPhone, MobPhone, HomePhone []int
	Fullname                       string
	Email                          string
}
type Card2 struct {
	Name, Adress, Surname string
}

func (v Card1) IfemptyName() bool {
	return (v.Fullname == "")

}
func (c Card2) IfemptyName() bool {
	return (c.Name == "")
}

type Ifemptyer interface {
	IfemptyName() bool
}

type addressBook []Card1

func (m addressBook) Less(i, j int) bool {
	return m[i].Fullname < m[j].Fullname
}
func (m addressBook) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
func (m addressBook) Len() int {
	return len(m)
}

func task_1() {

	var s Ifemptyer
	s = Card2{Name: "ivanov"}
	fmt.Println("поле имя заполнено =", s.IfemptyName())
	s = Card1{Email: "mail@mail.ru"}
	fmt.Println("поле имя заполнено =", s.IfemptyName())

}

func task_2() {

	arrAddressBook := []Card1{
		{
			Fullname:  "Sergeev Max",
			Email:     "mail@mail.ru",
			WorkPhone: []int{7471111, 3635500, 3635502},
			MobPhone:  []int{89055004444, 89052133322},
			HomePhone: []int{84953430112},
		},
		{
			Fullname:  "Ivanov Ivan",
			Email:     "mail@mail.ru",
			WorkPhone: []int{123123123, 3635500, 3635502},
			MobPhone:  []int{11122233344, 12333445566},
			HomePhone: []int{6667781},
		},
		{
			Fullname:  "Petrov Petr",
			Email:     "mail@mail.ru",
			WorkPhone: []int{1111111, 2222222, 3333333},
			MobPhone:  []int{2223344566},
			HomePhone: []int{7777717},
		},
	}

	fmt.Println(arrAddressBook)
	sort.Sort(addressBook(arrAddressBook))
	fmt.Println(arrAddressBook)

}

//  или все слишком просто или не понял я сути задание ...
func task_3() {

	input := ""
	for {
		fmt.Print("> ")
		if _, err := fmt.Scanln(&input); err != nil {
			fmt.Println(err)
			continue
		}

		if input == "exit" {
			break
		}
		if input == "help" {
			fmt.Println("spravka .... blablabla")
		}

	}

}
