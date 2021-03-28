package main

import (
	"fmt"
	hay "github.com/keigodasu/haystack-learning-basic-concept"
)

func main() {
	entity := hay.NewEntity("@site")
	entity.SetDis("the office building")

	tags := map[string]hay.Value{
		"unit01_name":   hay.NewStr("xxxxx"),
		"unit02_name":   hay.NewStr("yyyyy"),
		"primary_flag":  hay.NewBool(true),
		"secondly_flag": hay.NewBool(false),
	}

	entity.Tags = tags

	jsonEncoded, _ := entity.MarshallJSON()

	fmt.Println(string(jsonEncoded))

	grid := hay.Grid{
		Meta:     map[string]string{"ver": "3.0", "projName": "test"},
		Entities: []*hay.Entity{entity},
	}
	jsonEncoded, _ = grid.MarshallJSON()

	fmt.Println(string(jsonEncoded))
}
