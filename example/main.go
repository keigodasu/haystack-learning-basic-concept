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

	grid := hay.Grid{
		Meta:     map[string]string{"ver": "3.0", "projName": "test"},
		Entities: []*hay.Entity{entity},
	}
	jsonEncoded, _ := grid.MarshallJSON()

	fmt.Println(string(jsonEncoded))
	/*
	{
	  "meta": {
	    "projName": "test",
	    "ver": "3.0"
	  },
	  "cols": [
	    {
	      "name": "secondly_flag"
	    },
	    {
	      "name": "unit01_name"
	    },
	    {
	      "name": "unit02_name"
	    },
	    {
	      "name": "primary_flag"
	    }
	  ],
	  "rows": [
	    {
	      "primary_flag": "b:true",
	      "secondly_flag": "b:false",
	      "unit01_name": "s:xxxxx",
	      "unit02_name": "s:yyyyy"
	    }
	  ]
	}
	 */
}
