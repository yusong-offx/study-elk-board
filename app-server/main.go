package main

import (
	"fmt"
	"log"

	"github.com/yusong-offx/study-board/components"
)

type Test struct {
	Id       uint64 `json:'id'`
	Name     string `json:'name'`
	Password string `json:'password'`
}

func main() {
	components.DBConnect()
	defer components.DB.Close()

	test := Test{}
	if err := components.DB.QueryRow("SELECT * FROM TEST WHERE name = 'yusong'").Scan(&test.Id, &test.Name, &test.Password); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", test)
}
