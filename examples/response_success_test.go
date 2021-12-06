package dhkit_test

import (
	"encoding/json"
	"fmt"

	"github.com/digitalhouse-dev/dh-kit/response"
)

type exampleStruct struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func Example_OK_Logger() {
	ok := response.OK("", &exampleStruct{ID: "123", Name: "Example name"}, nil, nil)
	v, _ := json.Marshal(ok)
	fmt.Println(string(v))
	// Output: {"message":"Success request.","status":200,"data":{"id":"123","name":"Example name"}}
}

func Example_Created_Logger() {
	created := response.Created("", &exampleStruct{ID: "123", Name: "Example name"}, nil, nil)
	v, _ := json.Marshal(created)
	fmt.Println(string(v))
	// Output: {"message":"Created request.","status":201,"data":{"id":"123","name":"Example name"}}
}

func Example_Accepted_Logger() {
	accepted := response.Accepted("", &exampleStruct{ID: "123", Name: "Example name"}, nil, nil)
	v, _ := json.Marshal(accepted)
	fmt.Println(string(v))
	// Output: {"message":"Accepted request.","status":202,"data":{"id":"123","name":"Example name"}}
}

func Example_NonAuthoritativeInfo_Logger() {
	nonAuthoritativeInfo := response.NonAuthoritativeInfo("", &exampleStruct{ID: "123", Name: "Example name"}, nil, nil)
	v, _ := json.Marshal(nonAuthoritativeInfo)
	fmt.Println(string(v))
	// Output: {"message":"Non-Authoritative Information request.","status":203,"data":{"id":"123","name":"Example name"}}
}
func Example_NoContent_Logger() {
	noContent := response.NoContent("", &exampleStruct{ID: "123", Name: "Example name"}, nil, nil)
	v, _ := json.Marshal(noContent)
	fmt.Println(string(v))
	// Output: {"message":"No Content request.","status":204,"data":{"id":"123","name":"Example name"}}
}
func Example_ResetContent_Logger() {
	resetContent := response.ResetContent("", &exampleStruct{ID: "123", Name: "Example name"}, nil, nil)
	v, _ := json.Marshal(resetContent)
	fmt.Println(string(v))
	// Output: {"message":"Reset Content request.","status":205,"data":{"id":"123","name":"Example name"}}
}
func Example_PartialContent_Logger() {
	partialContent := response.PartialContent("", &exampleStruct{ID: "123", Name: "Example name"}, nil, nil)
	v, _ := json.Marshal(partialContent)
	fmt.Println(string(v))
	// Output: {"message":"Partial Content request.","status":206,"data":{"id":"123","name":"Example name"}}
}
