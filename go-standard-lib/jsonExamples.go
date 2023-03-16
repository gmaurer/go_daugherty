package main

import (
	"encoding/json"
	"log"
)

type ExampleStruct struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	CanCook bool   `json:"canCook"`
}

type ExampleNestedStruct struct {
	JobName    string          `json:"jobName"`
	Candidates []ExampleStruct `json:"candidates"`
}

func JsonMarshal(v any) []byte {
	msg, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}

	return msg
}

func JsonUnmarshal[T any](v []byte) *T {
	example := new(T)
	err := json.Unmarshal(v, example)
	if err != nil {
		panic(err)
	}

	return example
}

func TestMarshal() {
	c1 := ExampleStruct{Name: "Frank", Age: 50, CanCook: true}
	c2 := ExampleStruct{Name: "Marge", Age: 39, CanCook: false}
	candidates := []ExampleStruct{c1, c2}
	j := ExampleNestedStruct{JobName: "Chef", Candidates: candidates}

	jsonBytes := JsonMarshal(j)
	log.Println(string(jsonBytes))
}

func TestUnmarshal() {
	jsonText :=
		`{
	"jobName": "Accountant",
	"candidates": [
		{
			"name": "George",
			"age": 32,
			"canCook": false
		},
		{
			"name": "April",
			"age": 24,
			"canCook": true
		}
	]
}`

	result := JsonUnmarshal[ExampleNestedStruct]([]byte(jsonText))

	log.Printf("JobName: %s", result.JobName)
	for i, v := range result.Candidates {
		log.Printf("Candidate %v:\n", i+1)
		log.Println("Name:", v.Name)
		log.Println("Age:", v.Age)
		log.Println("Can Cook:", v.CanCook)
	}
}
