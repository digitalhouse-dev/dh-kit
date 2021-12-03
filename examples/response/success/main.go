package main

import (
	"encoding/json"
	"fmt"

	"github.com/digitalhouse-dev/dh-kit/response"
)

func main() {
	s := struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}{"123", "Example name"}
	var v []byte

	ok := response.OK("", &s, nil, nil)
	v, _ = json.Marshal(ok)
	fmt.Printf("%s\n\n", string(v))

	created := response.Created("", &s, nil, nil)
	v, _ = json.Marshal(created)
	fmt.Printf("%s\n\n", string(v))

	accepted := response.Accepted("", &s, nil, nil)
	v, _ = json.Marshal(accepted)
	fmt.Printf("%s\n\n", string(v))

	nonAuthoritativeInfo := response.NonAuthoritativeInfo("", &s, nil, nil)
	v, _ = json.Marshal(nonAuthoritativeInfo)
	fmt.Printf("%s\n\n", string(v))

	noContent := response.NoContent("", &s, nil, nil)
	v, _ = json.Marshal(noContent)
	fmt.Printf("%s\n\n", string(v))

	resetContent := response.ResetContent("", &s, nil, nil)
	v, _ = json.Marshal(resetContent)
	fmt.Printf("%s\n\n", string(v))

	partialContent := response.PartialContent("", &s, nil, nil)
	v, _ = json.Marshal(partialContent)
	fmt.Printf("%s\n\n", string(v))
}
