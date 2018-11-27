package parseStruct

import (
	"github.com/fatih/structtag"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"strings"
)

type StructFields []*ast.Field

func ParseStruct(filename string, src []byte, tagName string) (structMap map[string]StructFields, err error) {
	structMap = make(map[string]StructFields)

	if src == nil {
		src, err = ioutil.ReadFile(filename)
		if err != nil {
			return structMap, err
		}
	}
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, filename, src, parser.ParseComments)
	if err != nil {
		return structMap, err
	}

	collectStructs := func(x ast.Node) bool {
		ts, ok := x.(*ast.TypeSpec)
		if !ok || ts.Type == nil {
			return true
		}

		// 获取结构体名称
		structName := ts.Name.Name

		s, ok := ts.Type.(*ast.StructType)
		if !ok {
			return true
		}

		for _, field := range s.Fields.List {
			tag := field.Tag.Value
			tag = strings.Trim(tag, "`")
			tags, err := structtag.Parse(string(tag))
			if err != nil {
				return true
			}
			_, err = tags.Get(tagName)
			if err == nil {
				structMap[structName] = append(structMap[structName], field)
			}
		}
		return false
	}

	ast.Inspect(file, collectStructs)

	return structMap, nil
}
