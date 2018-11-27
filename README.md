# parseStruct
parse struct by tag
<br>
从文件中获取所有包含指定tag的结构体信息

## Process
- [x] 获取单文件中满足指定tag条件的结构体
- [ ] 获取嵌入式结构体信息

## Example

```
package main

import "fmt"

func main() {

	fieldsMap, err := ParseStruct("parse_struct_test.go", nil, "json")
	if err != nil {
		panic(err)
	}

	for structName, fields := range fieldsMap {
		fmt.Printf("structName:%s\n", structName)
		for _, field := range fields {
			fmt.Printf("	FieldName:%s\n", field.Names[0].Name)
			fmt.Printf("	FieldType:%s\n", field.Type)
			fmt.Printf("	FieldTag:%s\n", field.Tag.Value)
		}
	}
	return
}

// Output
parse_struct_test.go:50: structName:Example1
parse_struct_test.go:52: 	FieldName:Foo
parse_struct_test.go:53: 	FieldType:string
parse_struct_test.go:54: 	FieldTag:`json:"foo"`
parse_struct_test.go:50: structName:Example2
parse_struct_test.go:52: 	FieldName:Aoo
parse_struct_test.go:53: 	FieldType:int
parse_struct_test.go:54: 	FieldTag:`json:"aoo"`
```

## External references
* https://github.com/fatih/structtag