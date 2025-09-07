//go:build ignore

package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

const funcTemplate = `// Func_{{.ParamCount}}_{{.ReturnCount}}: {{.ParamCount}}个参数，{{.ReturnCount}}个返回值
func Func_{{.ParamCount}}_{{.ReturnCount}}[{{.TypeParams}}](fn func({{.FuncParams}}) {{.FuncReturns}}, keyGenerator func({{.FuncParams}}) string, ttlFn func({{.FuncParams}}) time.Duration) func({{.FuncParams}}) {{.FuncReturns}} {
	prefix := cfg.GetString("cache.prefix") + ":" + cfg.GetString("server.environment")
	return func({{.NamedParams}}) {{.NamedReturns}} {
		key := prefix + ":" + keyGenerator({{.ParamNames}})
		if bs, err := Get(redisConn, key); err == nil && len(bs) > 0 {
			var res = []any{}
			err := json.Unmarshal(bs, &res)
			if err != nil {
				return {{.ReturnZeros}}
			}
{{.ResultAssignments}}
			return {{.ReturnNames}}
		}
		{{.ReturnNames}} = fn({{.ParamNames}})
		// 检查是否有error，如果有，则不设置缓存
{{.ErrorChecks}}
		defer Set(redisConn, key, []any{ {{- .ReturnNames -}} }, ttlFn({{.ParamNames}}))
		return {{.ReturnNames}}
	}
}

`

type FuncData struct {
	ParamCount        int
	ReturnCount       int
	TypeParams        string
	FuncParams        string
	FuncReturns       string
	NamedParams       string
	NamedReturns      string
	ParamNames        string
	ReturnNames       string
	ReturnZeros       string
	ResultAssignments string
	ErrorChecks       string
}

func generateTypeParams(paramCount, returnCount int) string {
	var parts []string

	// 添加参数类型
	for i := 0; i < paramCount; i++ {
		parts = append(parts, fmt.Sprintf("P%d", i))
	}

	// 添加返回值类型
	for i := 0; i < returnCount; i++ {
		parts = append(parts, fmt.Sprintf("R%d", i))
	}

	if len(parts) == 0 {
		return ""
	}
	return strings.Join(parts, ", ") + " any"
}

func generateFuncParams(paramCount int) string {
	if paramCount == 0 {
		return ""
	}
	var parts []string
	for i := 0; i < paramCount; i++ {
		parts = append(parts, fmt.Sprintf("P%d", i))
	}
	return strings.Join(parts, ", ")
}

func generateFuncReturns(returnCount int) string {
	if returnCount == 1 {
		return "R0"
	}
	var parts []string
	for i := 0; i < returnCount; i++ {
		parts = append(parts, fmt.Sprintf("R%d", i))
	}
	return "(" + strings.Join(parts, ", ") + ")"
}

func generateNamedParams(paramCount int) string {
	if paramCount == 0 {
		return ""
	}
	var parts []string
	for i := 0; i < paramCount; i++ {
		parts = append(parts, fmt.Sprintf("p%d P%d", i, i))
	}
	return strings.Join(parts, ", ")
}

func generateNamedReturns(returnCount int) string {
	var parts []string
	for i := 0; i < returnCount; i++ {
		parts = append(parts, fmt.Sprintf("r%d R%d", i, i))
	}
	return "(" + strings.Join(parts, ", ") + ")"
}

func generateParamNames(paramCount int) string {
	if paramCount == 0 {
		return ""
	}
	var parts []string
	for i := 0; i < paramCount; i++ {
		parts = append(parts, fmt.Sprintf("p%d", i))
	}
	return strings.Join(parts, ", ")
}

func generateReturnNames(returnCount int) string {
	var parts []string
	for i := 0; i < returnCount; i++ {
		parts = append(parts, fmt.Sprintf("r%d", i))
	}
	return strings.Join(parts, ", ")
}

func generateReturnZeros(returnCount int) string {
	var parts []string
	for i := 0; i < returnCount; i++ {
		parts = append(parts, fmt.Sprintf("r%d", i))
	}
	return strings.Join(parts, ", ")
}

func generateResultAssignments(returnCount int) string {
	var parts []string
	for i := 0; i < returnCount; i++ {
		parts = append(parts, fmt.Sprintf("\t\t\tif len(res) > %d {\n\t\t\t\tr%d, _ = res[%d].(R%d)\n\t\t\t}", i, i, i, i))
	}
	return strings.Join(parts, "\n")
}

func generateErrorChecks(returnCount int) string {
	var parts []string
	for i := 0; i < returnCount; i++ {
		parts = append(parts, fmt.Sprintf("\t\tif _, ok := any(r%d).(error); ok {\n\t\t\treturn %s\n\t\t}", i, generateReturnNames(returnCount)))
	}
	return strings.Join(parts, "\n")
}

func main() {
	tmpl, err := template.New("func").Parse(funcTemplate)
	if err != nil {
		panic(err)
	}

	var output strings.Builder

	// 生成从 Func_0_1 到 Func_6_6 的所有函数
	for paramCount := 0; paramCount <= 6; paramCount++ {
		for returnCount := 1; returnCount <= 6; returnCount++ {
			// 跳过已存在的 Func_2_2
			if paramCount == 2 && returnCount == 2 {
				continue
			}

			data := FuncData{
				ParamCount:        paramCount,
				ReturnCount:       returnCount,
				TypeParams:        generateTypeParams(paramCount, returnCount),
				FuncParams:        generateFuncParams(paramCount),
				FuncReturns:       generateFuncReturns(returnCount),
				NamedParams:       generateNamedParams(paramCount),
				NamedReturns:      generateNamedReturns(returnCount),
				ParamNames:        generateParamNames(paramCount),
				ReturnNames:       generateReturnNames(returnCount),
				ReturnZeros:       generateReturnZeros(returnCount),
				ResultAssignments: generateResultAssignments(returnCount),
				ErrorChecks:       generateErrorChecks(returnCount),
			}

			err := tmpl.Execute(&output, data)
			if err != nil {
				panic(err)
			}
		}
	}

	// 写入到文件
	err = os.WriteFile("cache_funcs_generated.go", []byte("package cache\n\nimport (\n\t\"time\"\n)\n\n"+output.String()), 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("Generated cache functions successfully!")
}
