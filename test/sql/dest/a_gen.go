package sql

import "github.com/llyb120/gox"

func TestGox() gox.Query {
	var id = 1
	_ = id
	return func() (__result gox.Query) {
		__gox_sql_0_builder := gox.NewQueryBuilder()
		__gox_sql_0_builder.AddText("\r")
		__gox_sql_0_builder.AddText("\n")
		__gox_sql_0_builder.AddText("\t\tselect * from test where id = ")
		__gox_sql_0_builder.AddParam(id)
		__gox_sql_0 := __gox_sql_0_builder.Build()
		return __gox_sql_0
	}()
}
