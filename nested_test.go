package nested_test

import (
	"encoding/json"
	//"fmt"
	"github.com/opesun/nested"
	"testing"
)

var j = `{
	"hello": {
		"this": {
			"is": {
				"an": {
					"example": "hi"
				}
			}
		}
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
	// Coming soon.
}
