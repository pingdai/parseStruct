package parseStruct

import (
	"testing"
)

type Example1 struct {
	// Foo Comments
	Foo string `json:"foo"`
}

type Example2 struct {
	// Aoo Comments
	Aoo int `json:"aoo"`
}

func TestParseStructFromSrc(t *testing.T) {
	src := `package main
	        type Example struct {
		// Foo comments
	    Foo string ` + "`json:\"foo\"`\n" +
		"AA int " + "`json:\"aa,string\"`}"

	fieldsMap, err := ParseStruct("./test.go", []byte(src), "json")
	if err != nil {
		t.Fatalf("err:%v", err)
		return
	}

	for structName, fields := range fieldsMap {
		t.Logf("structName:%s", structName)
		for _, field := range fields {
			t.Logf("	FieldName:%s", field.Names[0].Name)
			t.Logf("	FieldType:%s", field.Type)
			t.Logf("	FieldTag:%s", field.Tag.Value)
		}
	}
	return
}

func TestParseStructFromFile(t *testing.T) {

	fieldsMap, err := ParseStruct("parse_struct_test.go", nil, "json")
	if err != nil {
		t.Fatalf("err:%v", err)
		return
	}

	for structName, fields := range fieldsMap {
		t.Logf("structName:%s", structName)
		for _, field := range fields {
			t.Logf("	FieldName:%s", field.Names[0].Name)
			t.Logf("	FieldType:%s", field.Type)
			t.Logf("	FieldTag:%s", field.Tag.Value)
		}
	}
	return
}
