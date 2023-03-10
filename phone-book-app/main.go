package main

import (
	"fmt"
	"os"
	"path"
)

var data = []Record{}

type Record struct {
	Name	string
	Surname	string
	Tel 	string
}

func search(key string) *Record {
	for i, v := range data {
		if v.Surname == key {
			fmt.Println(data[i])
			return &data[i]
		}
	}
	return nil
}

func test() Record {
	return data[0]
}

func list() {
	for _, record := range data {
		fmt.Println(record)
	}
}
func init() {
	fmt.Println("in init")
	data = append(data, Record{"Idris", "Yakup", "1234134"})
	data = append(data, Record{"Zekeriyya", "Yahya", "1234134"})
}

func main() {
	fmt.Println("in main")
	r := data[1]
	fmt.Println("---")	
	fmt.Println(&r)
	fmt.Println("---")	
	arguments := os.Args
	if len(arguments) == 1 {
		exe := path.Base(arguments[0])
		fmt.Printf("Usage: %s seach|list <arguments>\n", exe)
		return
	}

	switch arguments[1] {
	case "search":
		if len(arguments) !=3 {
			fmt.Println("Usage: search Surname")
			return
		}
		result := search(arguments[2])
		fmt.Println(result)
		fmt.Println(*result)
	case "list":
		list()
	default:
		fmt.Println("Not a valid option")
	}

}