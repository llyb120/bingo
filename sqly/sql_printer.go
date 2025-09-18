package sqly

import (
	"fmt"
	"github.com/llyb120/bingo/core"
	"github.com/llyb120/bingo/log"
	"reflect"
	"regexp"
	"strings"
)

var spaceReg = regexp.MustCompile(`--.*?\n|[\r\n\t]+`)
var spaceReg_ = regexp.MustCompile(`[\r\n\t]+`)
var spaceReg2 = regexp.MustCompile(`\s{2,}`)

func PrintSql(c core.Context, sql string, args ...interface{}) string {
	//now := time.Now()
	defer func() {
		if err := recover(); err != nil {
			log.Error(c, "SQL执行打印错误,err = %v", err)
			//logging.ErrorF(c, "SQL执行打印错误,err = %v", err)
		}
	}()
	if len(args) == 0 || args[0] == nil {
		handleSql(&sql, true)
		return sql
	}
	var finalSql = ""

	nArgs := make([]interface{}, 0)
	for _, arg := range args {
		switch arg.(type) {
		case []string, []int, []int32, []int64, []float32, []float64, []interface{}:
			slice, _ := TakeSliceArg(arg)
			nArgs = append(nArgs, slice...)
		default:
			nArgs = append(nArgs, arg)
		}
	}

	pos := 0
	totalLen := len(sql)
	current := 0
	var builder strings.Builder
	for pos < totalLen {
		if sql[pos] == '\'' || sql[pos] == '"' || sql[pos] == '`' {
			breakWord := sql[pos]
			builder.WriteByte(breakWord)
			// 读到下一个'
			for pos = pos + 1; pos < totalLen && sql[pos] != breakWord; pos++ {
				builder.WriteByte(sql[pos])
			}
			builder.WriteByte(breakWord)
			if pos == totalLen {
				// 没有找到对应的'
				break
			}
		} else if pos < totalLen-1 && sql[pos] == '-' && sql[pos+1] == '-' {
			// 处理SQL注释：从--开始到行尾
			builder.WriteByte(' ')
			for pos = pos + 2; pos < totalLen && sql[pos] != '\n' && sql[pos] != '\r'; pos++ {
				// 跳过注释内容，不写入builder
			}
			if pos < totalLen && (sql[pos] == '\n' || sql[pos] == '\r') {
				builder.WriteByte(' ')
			}
		} else if sql[pos] == '?' {
			if current >= len(nArgs) {
				break
			}
			v := nArgs[current]
			replacedStr := ""
			// reflectValue := reflect.ValueOf(v)
			// reflectKind := reflectValue.Kind()
			switch v.(type) {
			case int, int8, int16, int32, int64:
				replacedStr = fmt.Sprintf("%d", v)
			case float32, float64:
				replacedStr = fmt.Sprintf("%f", v)
			default:
				if v == nil {
					replacedStr = "NULL"
				} else {
					str := fmt.Sprintf("%v", v)
					if str == "<nil>" {
						replacedStr = "NULL"
					} else {
						// 对单引号进行转义
						str = strings.ReplaceAll(str, "'", "\\'")
						replacedStr = "'" + str + "'"
					}
				}
			}
			builder.WriteString(replacedStr)
			current++
		} else {
			builder.WriteByte(sql[pos])
		}
		pos++
	}
	finalSql = builder.String()

	// 处理skyWalking对换行符的支持问题
	//finalSql = spaceReg.ReplaceAllString(finalSql, " ")
	//finalSql = spaceReg2.ReplaceAllString(finalSql, " ")
	handleSql(&finalSql, false)
	return finalSql
}

// 是否去除注释
func handleSql(sql *string, isRemoveComment bool) {
	if isRemoveComment {
		*sql = spaceReg.ReplaceAllString(*sql, " ")
	} else {
		// 保留注释
		*sql = spaceReg_.ReplaceAllString(*sql, " ")
	}
	*sql = spaceReg2.ReplaceAllString(*sql, " ")

}

// 通用函数：参数分片
func TakeSliceArg(arg interface{}) (out []interface{}, ok bool) {
	slice, success := TakeArg(arg, reflect.Slice)
	if !success {
		ok = false
		return
	}
	c := slice.Len()
	out = make([]interface{}, c)
	for i := 0; i < c; i++ {
		out[i] = slice.Index(i).Interface()
	}
	return out, true
}

// 通用函数：参数分片
func TakeArg(arg interface{}, kind reflect.Kind) (val reflect.Value, ok bool) {
	val = reflect.ValueOf(arg)
	if val.Kind() == kind {
		ok = true
	}
	return
}
