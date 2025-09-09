package sql

import "github.com/llyb120/gox"

func TestGox() gox.Query {
	var id = 1
	_ = id
	return gox.Sql( /*
		select * from test where id = #{id}
	*/)
}
