package nested_test

import (
	"encoding/json"
	//"fmt"
	"testing"

	"github.com/crufter/nested"
)

var j = `{
	"hello": {
		"this": {
			"is": {
				"an": {
					"example": "hi"
				}
			}
		},
		"that": [{"try":"this"}]
	}
}`

func TestBasic(t *testing.T) {
	m := map[string]interface{}{}
	err := json.Unmarshal([]byte(j), &m)
	if err != nil {
		t.Fatal(err)
	}
	magic, ok := nested.GetStr(m, "hello.this.is.an.example")
	if !ok {
		t.Fatal("Can't find.")
	}
	if magic != "hi" {
		t.Fatal(magic)
	}
	v, ok := nested.GetStr(m, "hello.that[0].try")
	if !ok || v != "this" {
		t.Fatal(v)
	}
	// Coming soon.
}
