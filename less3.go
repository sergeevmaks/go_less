package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	task_1_2()
	task_3()
	task_4()
}

type descr_car struct {
	colour, brand, number              string
	year                               int
	motion, open_windows, engine_start bool
}
type driver struct {
	name           string
	number_license string
}
type location struct {
	N, W float64
}
type crash struct {
	car1, car2        descr_car
	driver1, driver2  driver
	place_of_accident location
}

type vcard struct {
	WorkPhone, MobPhone, HomePhone []int
	Fullname                       string `json:"Full_Name"`
	Email                          string
}

var fifo []string

func task_4() {
	addressBook := make(map[string]vcard)
	addressBook["max"] = vcard{
		Fullname:  "Sergeev Max",
		Email:     "mail@mail.ru",
		WorkPhone: []int{7471111, 3635500, 3635502},
		MobPhone:  []int{89055004444, 89052133322},
		HomePhone: []int{84953430112},
	}

	addressBook["ivan"] = vcard{
		Fullname:  "Ivanov Ivan",
		Email:     "mail@mail.ru",
		WorkPhone: []int{123123123, 3635500, 3635502},
		MobPhone:  []int{11122233344, 12333445566},
		HomePhone: []int{6667781},
	}
	addressBook["petr"] = vcard{
		Fullname:  "Petrov Petr",
		Email:     "mail@mail.ru",
		WorkPhone: []int{1111111, 2222222, 3333333},
		MobPhone:  []int{2223344566},
		HomePhone: []int{7777717},
	}

	json, err := json.Marshal(addressBook)
	if err != nil {
		fmt.Println(err)
		return
	}
	content, errread := ioutil.ReadFile("addressBook.json")
	if errread != nil {
		log.Fatal(errread)
	}

	//b1 := []byte(content)
	//m := make(map[string]vcard)
	//err = json.Unmarshal(b1, &m)
	//if err != nil {
	//		log.Println(err)
	//		return
	//	}

	if string(json) != string(content) { //наверное правильней сравнивать распарсенные хеши addressBook ... но для этой конкретной задачи так короче
		message := []byte(json)
		err = ioutil.WriteFile("addressBook.json", message, 0644)

		if err != nil {
			log.Fatal(err)
		}

	}

}

func task_3() {

	Push("первый")
	Push("второй")
	Push("третий")
	Push("четвертый")
	fmt.Println(fifo)
	fmt.Println(Pop())
	fmt.Println(Pop())
	fmt.Println(Pop())
	fmt.Println(Pop())
}

func Push(str string) {
	fifo = append(fifo, str)
}

func Pop() string {
	num := len(fifo)
	if num == 0 {
		return ""
	}
	popElem := fifo[0]
	fifo = fifo[1:]
	return popElem
}

func task_1_2() {
	driver1 := driver{name: "Max", number_license: "123DD1233"}
	driver2 := driver{name: "Ivan", number_license: "333AAA1111"}
	driver3 := driver{name: "Ivan", number_license: "333AAA1111"}
	crash1 := crash{driver1: driver1, driver2: driver2, place_of_accident: location{55.742793, 37.61540100000002}}
	if driver2 == driver3 {
		fmt.Println("дубликат")
	}
	fmt.Printf("координаты аварии = %v W , %v N \n", crash1.place_of_accident.N, crash1.place_of_accident.W)
}
