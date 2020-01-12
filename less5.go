package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	task_2()
	task_3()

}

func task_2() {
	file, err := os.Open("less5.go")
	if err != nil {
		return
	}
	defer file.Close()
	// выделяем память не сразу под весь файл а считываем кусками по 64байт .... так оптимальней при работе с большими файлами
	data := make([]byte, 64)

	for {
		n, err := file.Read(data)
		if err == io.EOF {
			break
		}
		fmt.Print(string(data[:n]))
	}
}

func task_3() {
	type vcard struct {
		Name, Adress, Surname string
	}
	card := vcard{Name: "Max", Adress: "Moscow", Surname: "Sergeev"}
	enc := csv.NewWriter(os.Stdout)
	enc.Write([]string{"Name", "Adress", "Surname"})
	enc.Write([]string{
		card.Name,
		card.Adress,
		card.Surname,
	})
	enc.Flush()
}
