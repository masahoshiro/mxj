package mxj

import (
	"testing"
)

func TestRenameKey(t *testing.T) {
	m := map[string]interface{}{
		"Div": map[string]interface{}{
			"Colour": "blue",
		},
	}
	mv := Map(m)
	err := mv.RenameKey("Div.Colour", "Color")
	if err != nil {
		t.Fatal("err: ", err)
	}
	values, err := mv.ValuesForPath("Div.Color")
	if len(values) != 1 {
		t.Fatal("err: didn't add the new key")
	}
	if values[0] != "blue" {
		t.Fatal("err: value is changed")
	}
	values, err = mv.ValuesForPath("Div.Colour")
	if len(values) > 0 {
		t.Fatal("err: didn't removed the old key")
	}
}