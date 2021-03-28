package main

import (
	"fmt"
	hay "github.com/keigodasu/haystack-learning-basic-concept"
)

func main() {
	entity := hay.NewEntity("@site")
	entity.SetDis("the office building")

	tags := map[hay.Label]hay.Value{
		hay.Label("unit01_name"):   hay.NewStr("xxxxx"),
		hay.Label("unit02_name"):   hay.NewStr("yyyyy"),
		hay.Label("primary_flag"):  hay.NewBool(true),
		hay.Label("secondly_flag"): hay.NewBool(false),
	}

	entity.Tags = tags

	jsonEncoded, _ := entity.MarshallJSON()

	fmt.Println(string(jsonEncoded))
}
